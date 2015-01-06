package main

import (
	"errors"
	"fmt"
	"github.com/unixpickle/gospeech"
	"io/ioutil"
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
	if len(os.Args) != 3 {
		return errors.New("Usage: startstops <dictionary.txt> <common.txt>")
	}
	
	// Read input files
	fmt.Println("Reading input files...")
	dict, err := gospeech.LoadDictionary(os.Args[1])
	if err != nil {
		return err
	}
	common, err := ReadLines(os.Args[2])
	if err != nil {
		return err
	}
	
	// Sort the dictionary
	fmt.Println("Sorting dictionary...")
	dictKeys := SortDictionary(dict, common)
	
	// Generate keys and values
	fmt.Println("Generating starts and stops...")
	result := map[string][]string{}
	for _, word := range dictKeys {
		phones := dict.Get(word)
		if phones == nil {
			continue
		}
		
		prefixPhone := "-" + phones[0].Letters
		suffixPhone := phones[len(phones)-1].Letters + "-"
		
		for _, key := range []string{prefixPhone, suffixPhone} {
			if list, ok := result[key]; !ok {
				result[key] = []string{word}
			} else if len(list) < 10 {
				result[key] = append(result[prefixPhone], word)
			}
		}
	}
	
	PrintKeyValues(result)
	return nil
}

func PrintKeyValues(phones map[string][]string) {
	keys := make([]string, 0, len(phones))
	for key, _ := range phones {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, phone := range keys {
		fmt.Println(phone, "::", strings.Join(phones[phone], ", "))
		fmt.Println("")
	}
}

func ReadLines(path string) ([]string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	list := strings.Split(string(contents), "\n")
	sort.Strings(list)
	return list, nil
}

func SortDictionary(dict gospeech.Dictionary, common []string) []string {
	inCommon := make([]string, 0)
	uncommon := make([]string, 0)
	for word, _ := range dict {
		w := strings.ToLower(word)
		idx := sort.SearchStrings(common, w)
		if idx == len(common) || common[idx] != w {
			uncommon = append(uncommon, word)
		} else {
			inCommon = append(inCommon, word)
		}
	}
	return append(inCommon, uncommon...)
}
