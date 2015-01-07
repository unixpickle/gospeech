package main

import (
	"fmt"
	"github.com/unixpickle/gospeech"
	"github.com/unixpickle/wav"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "Usage: say-phones <output> <voice>",
			"<phones ...>")
		os.Exit(1)
	}
	phones := make([]gospeech.Phone, 0, len(os.Args)-3)
	for _, str := range os.Args[3:] {
		p, err := gospeech.ParsePhone(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid phone: "+str)
			os.Exit(1)
		}
		phones = append(phones, p)
	}
	voice, err := gospeech.LoadVoice(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load voice:", err)
		os.Exit(1)
	}
	sound := gospeech.SynthesizePhones(phones, voice)
	if err := wav.WriteFile(sound, os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write output:", err)
		os.Exit(1)
	}
}
