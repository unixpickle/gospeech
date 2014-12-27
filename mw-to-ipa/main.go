package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	if err := errMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addSuffixes(words map[string]string) {
	suffixes := map[string]string{"ing": "ɪŋ", "s": "s", "er": "ɜr",
		"ers": "ɜrs", "ly": "lē", "ness": "nis", "ment": "mənt",
		"ments": "mənts"}
	// Change words but don't change the original so that shorter versions of
	// words don't change.
	newWords := map[string]string{}
	for word, ipa := range words {
		// Use the longest matching suffix
		suffix := ""
		addition := ""
		for theSuffix, theAddition := range suffixes {
			if !strings.HasSuffix(word, theSuffix) {
				continue
			} else if len(theSuffix) > len(suffix) {
				suffix = theSuffix
				addition = theAddition
			}
		}
		if suffix == "" {
			continue
		}
		orig := word[0:len(word)-len(suffix)]
		if origIPA, ok := words[orig]; ok && origIPA == ipa {
			newWords[word] = ipa + addition
		}
	}
	for word, ipa := range newWords {
		words[word] = ipa
	}
}

func convertAlphabet(words map[string]string) {
	for word, mw := range words {
		// TODO: replace mw strings with ipa strings
	}
}

func errMain() error {
	if len(os.Args) != 3 {
		return errors.New("Usage: mw-to-ipa <input.txt> <output.txt>")
	}
	
	result, err := readInput(os.Args[1])
	if err != nil {
		return err
	}
	
	removeExtraInfo(result)
	convertAlphabet(result)
	addSuffixes(result)
	
	return writeOutput(os.Args[2], result)
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
	return res, nil
}

func removeExtraInfo(words map[string]string) {
	for word, phonetics := range words {
		phonetics = phonetics[1:len(ipa)-1]
		idx := strings.Index(phonetics, " ")
		if idx >= 0 {
			phonetics = phonetics[0:idx]
		}
		if strings.HasSuffix(phonetics, ",") {
			phonetics = phonetics[0:len(phonetics)-1]
		}
		words[word] = phonetics
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

