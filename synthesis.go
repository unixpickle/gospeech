package gospeech

import (
	"github.com/unixpickle/wav"
	"math"
	"strings"
	"time"
)

func SynthesizeWord(word string, d *Dictionary, s *Sounds) wav.Sound {
	phonemes := d.Get(word)
	
	// If the word is not in the dictionary, return a second-long bleep sound.
	if phonemes == nil || len(phonemes) == 0 {
		sampleRate := 44100
		sound := wav.NewPCM8Sound(1, sampleRate)
		const freq = 500
		for i := 0; i < sampleRate*1; i++ {
			time := float64(i) / float64(sampleRate)
			value := wav.Sample(math.Sin(time * math.Pi * 2 * freq))
			sound.SetSamples(append(sound.Samples(), value))
		}
		return sound
	}
	
	// TODO: use some fancy overlay/merging stuff here
	sound := s.Get(phonemes[0]).Clone()
	for _, aSound := range phonemes[1:] {
		wav.Append(sound, s.Get(aSound))
	}
	
	return sound
}

func SynthesizeSentence(sentence string, d *Dictionary, s *Sounds) wav.Sound {
	res := wav.NewPCM16Sound(2, 44100)
	for i, word := range strings.Split(sentence, " ") {
		if i != 0 {
			wav.AppendSilence(res, time.Second / 2)
		}
		wav.Append(res, SynthesizeWord(word, d, s))
	}
	return res
}
