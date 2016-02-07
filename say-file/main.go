package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/unixpickle/gospeech"
	"github.com/unixpickle/wav"
)

func main() {
	var rawPhonetics bool
	if len(os.Args) == 2 {
		if os.Args[1] == "-phonetics" {
			rawPhonetics = true
		} else {
			fmt.Fprintln(os.Stderr, "Usage: say-file [-phonetics]")
		}
	}

	if rawPhonetics {
		fmt.Println("Please enter some IPA text:")
	} else {
		fmt.Println("Please enter some English text:")
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var phonetics string
	if rawPhonetics {
		phonetics = string(input)
	} else {
		dict, err := gospeech.LoadDictionary("../dict/cmudict-IPA.txt")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		phonetics = dict.TranslateToIPA(string(input))
	}

	synthesized := gospeech.DefaultVoice.Synthesize(phonetics, 6)
	wav.WriteFile(synthesized, "output.wav")
	fmt.Println("Saved output.wav")
}
