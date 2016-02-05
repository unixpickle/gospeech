package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/unixpickle/gospeech"
	"github.com/unixpickle/wav"
)

func main() {
	fmt.Println("Please enter some phonetic text:")
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	phonetics := string(input)
	synthesized := gospeech.DefaultVoice.Synthesize(phonetics, 6)
	wav.WriteFile(synthesized, "output.wav")
	fmt.Println("Saved output.wav")
}
