package main

import (
	"errors"
	"fmt"
	"github.com/unixpickle/gospeech"
	"os"
	"sort"
	"strings"
)

func main() {
	if err := ErrMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ErrMain() error {
	if len(os.Args) != 2 {
		return errors.New("Usage: diphones <dictionary.txt>")
	}
	dict, err := gospeech.LoadDictionary(os.Args[1])
	if err != nil {
		return err
	}
	diphones := map[string][]string{}
	for word, raw := range dict {
		phones := dict.Get(word)
		if phones == nil {
			continue
		}
		// Find all the diphones and check them
		for i := 0; i < len(phones)-1; i++ {
			diphone := phones[i].Letters + "-" + phones[i+1].Letters
			if list, ok := diphones[diphone]; ok && len(list) > 10 {
				continue
			} else if !ok {
				diphones[diphone] = []string{}
			}
			example := strings.ToLower(word) + " (" + raw + ")"
			diphones[diphone] = append(diphones[diphone], example)
		}
	}
	keys := make([]string, 0, len(diphones))
	for key, _ := range diphones {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, diphone := range keys {
		fmt.Println(diphone, "::", strings.Join(diphones[diphone], ", "))
		fmt.Println("")
	}
	return nil
}
