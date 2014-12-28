package main

import (
	"errors"
	"fmt"
	"github.com/unixpickle/gospeech/dictionary"
	"github.com/unixpickle/gospeech/wordlist"
	"os"
	"sync"
	"sync/atomic"
)

func main() {
	if err := errMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func errMain() error {
	if len(os.Args) != 3 {
		return errors.New("Usage: find-ipa <wordlist.txt> <output.txt>")
	}
	words, err := wordlist.ReadLines(os.Args[1])
	if err != nil {
		return err
	}

	// Run 16 worker threads
	wg := sync.WaitGroup{}
	errs := int32(0)
	lock := sync.Mutex{}
	ch := make(chan string)
	res := wordlist.Mapping{}
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			for {
				word, ok := <-ch
				if !ok {
					break
				}
				ipa, err := cambridge.FindFullIPA(word)
				if err != nil {
					atomic.AddInt32(&errs, 1)
				} else {
					lock.Lock()
					res[word] = ipa
					lock.Unlock()
				}
			}
			wg.Done()
		}()
	}
	// Pass each word to a worker thread
	for i, word := range words {
		ch <- word
		if i%500 == 0 {
			errCount := atomic.LoadInt32(&errs)
			fmt.Println("Did", i, "words,", errCount, "errors")
		}
	}
	close(ch)
	wg.Wait()
	fmt.Println("Done with", atomic.LoadInt32(&errs), "errors")
	return res.WriteFile(os.Args[2])
}
