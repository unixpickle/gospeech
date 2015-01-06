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
		return errors.New("Usage: diphones <dictionary.txt> <common.txt>")
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
	
	// Generate diphones
	fmt.Println("Generating diphones...")
	diphones := map[string][]string{}
	for _, word := range dictKeys {
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
			example := strings.ToLower(word) + " (" + dict[word] + ")"
			diphones[diphone] = append(diphones[diphone], example)
		}
	}
	
	PrintDiphones(diphones)
	return nil
}

func PrintDiphones(diphones map[string][]string) {
	keys := make([]string, 0, len(diphones))
	for key, _ := range diphones {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, diphone := range keys {
		fmt.Println(diphone, "::", strings.Join(diphones[diphone], ", "))
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
