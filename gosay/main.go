package main

import (
	"errors"
	"fmt"
	"github.com/unixpickle/gospeech"
	"github.com/unixpickle/wav"
	"os"
)

func main() {
	if err := ErrMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ErrMain() error {
	if len(os.Args) != 5 {
		return errors.New("Usage: gospeech <dictionary.txt> <voicedir> " +
			"<sentence> <output.wav>")
	}
	fmt.Println("Loading dictionary...")
	dict, err := gospeech.LoadDictionary(os.Args[1])
	if err != nil {
		return err
	}
	fmt.Println("Loading voice...")
	voice, err := gospeech.LoadVoice(os.Args[2])
	if err != nil {
		return err
	}
	fmt.Println("Synthesizing...")
	output := gospeech.SynthesizeSentence(os.Args[3], dict, voice)
	fmt.Println("Saving...")
	return wav.WriteFile(output, os.Args[4])
}