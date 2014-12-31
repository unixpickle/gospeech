package gospeech

import (
	"github.com/unixpickle/wav"
	"math"
	"strings"
	"time"
)

func SynthesizeWord(word string, d Dictionary, v Voice) wav.Sound {
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
	
	sound := v.Get(phonemes[0]).Clone()
	for _, phoneme := range phonemes[1:] {
		aSound := v.Get(phoneme).Clone()
		
		if phoneme.Type != TypeVowel {
			const offsetTime = time.Millisecond * 20
			wav.Overlay(sound, aSound, sound.Duration() - offsetTime)
			continue
		}
		
		// Fade in the vowel sound
		const fadeTime = time.Millisecond * 20
		const vowelTime = time.Millisecond * 200
		wav.Crop(aSound, aSound.Duration() - vowelTime, aSound.Duration())
		wav.Gradient(aSound, 0, fadeTime)
		wav.Gradient(aSound, vowelTime, vowelTime - fadeTime)
		wav.Overlay(sound, aSound, sound.Duration() - fadeTime)
	}
	
	return sound
}

func SynthesizeSentence(sentence string, d Dictionary, v Voice) wav.Sound {
	res := wav.NewPCM16Sound(2, 44100)
	for i, word := range strings.Split(sentence, " ") {
		if i != 0 {
			wav.AppendSilence(res, time.Second / 3)
		}
		wav.Append(res, SynthesizeWord(word, d, v))
	}
	return res
}
