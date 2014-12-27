package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	if err := errMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func errMain() error {
	if len(os.Args) != 3 {
		return errors.New("Usage: normalize-ipa <input.txt> <output.txt>")
	}
	input, err := readInput(os.Args[1])
	if err != nil {
		return err
	}
	// First, replace each IPA string with the first full word in it
	
	return nil
}

func readInput(path string) (map[string]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	res := map[string]string{}
	for _, line := range bytes.Split(content, []byte("\n")) {
		idx := bytes.Index(line, []byte("|"))
		if idx < 0 {
			continue
		}
		res[string(line[0:idx])] = string(line[idx+1:])
	}
}

func writeOutput(path string, res map[string]string) error {
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()

	// Alphabetize the keys
	keys := make([]string, 0, len(res))
	for word, _ := range res {
		keys = append(keys, word)
	}
	sort.Strings(keys)

	// Write the sorted mapping
	for _, word := range keys {
		output.Write([]byte(word + "|" + res[word] + "\n"))
	}
	return nil
}

