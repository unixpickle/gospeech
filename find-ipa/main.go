package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"sync"
	"sync/atomic"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: find-ipa <wordlist.txt> <output.txt>")
		os.Exit(1)
	}

	// Open the output file to make sure it can be created.
	output, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer output.Close()

	// Read the words
	file := os.Args[1]
	words, err := readWords(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Download the IPA versions
	res := downloadIPA(words)
	keys := make([]string, 0, len(res))
	for word, _ := range res {
		keys = append(keys, word)
	}
	sort.Strings(keys)

	// Write the sorted IPA versions
	for _, word := range keys {
		output.Write([]byte(word + "|" + res[word] + "\n"))
	}
}

func downloadIPA(words []string) map[string]string {
	res := map[string]string{}
	var numErrors int32
	lock := sync.Mutex{}
	ch := make(chan string)
	wg := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				word, ok := <-ch
				if !ok {
					return
				}
				ipa, err := findIPA(word)
				if err == nil {
					lock.Lock()
					res[word] = ipa
					lock.Unlock()
				} else {
					atomic.AddInt32(&numErrors, 1)
				}
			}
		}()
	}
	for i, word := range words {
		if i%1000 == 0 {
			errs := atomic.LoadInt32(&numErrors)
			fmt.Println("Did", i, "words and", errs, "errors")
		}
		ch <- word
	}
	close(ch)
	wg.Wait()
	errs := atomic.LoadInt32(&numErrors)
	fmt.Println("Done with", errs, "errors")
	return res
}

func findIPA(word string) (string, error) {
	url := "http://www.merriam-webster.com/dictionary/" + word
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	e := "<span class=\"pr\">\\\\.*?\\\\<\\/span>"
	match := regexp.MustCompile(e).FindString(string(data))
	if match == "" {
		return "", errors.New("Not found.")
	}
	return regexp.MustCompile("<.*?>").ReplaceAllString(match, ""), nil
}

func readWords(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b := bufio.NewReader(f)
	res := make([]string, 0)
	for {
		line, _, err := b.ReadLine()
		if err != nil {
			break
		}
		res = append(res, string(line))
	}
	return res, nil
}
