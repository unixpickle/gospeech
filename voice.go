package gospeech

import (
	"fmt"
	"time"

	"github.com/unixpickle/wav"
)

type Voice struct {
	Phones map[string]*Phone
}

func (v Voice) Synthesize(ipaString string, phoneRate float64) wav.Sound {
	s := wav.NewPCM8Sound(1, 44100)
	var lastPhone *Phone
	for _, ph := range []rune(ipaString) {
		if ph == ' ' {
			wav.AppendSilence(s, time.Second)
			lastPhone = nil
			continue
		}

		phone := v.Phones[string(ph)]
		if phone != nil {
			fmt.Println("recognized phone:", string(ph))
			phone.Synthesize(lastPhone, s, phoneRate)
			lastPhone = phone
		} else {
			fmt.Println("unknown phone:", string(ph))
		}
	}
	return s
}

var DefaultVoice = Voice{Phones: map[string]*Phone{
	"i": &Phone{
		Duration:       1,
		Formants:       [3]float64{280, 2250, 2890},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"I": &Phone{
		Duration:       1,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"e": &Phone{
		Duration:       1,
		Formants:       [3]float64{400, 2200, 2890},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ə": &Phone{
		Duration:       1,
		Formants:       [3]float64{500, 1500, 2490},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ɛ": &Phone{
		Duration:       1,
		Formants:       [3]float64{550, 1770, 2490},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ʌ": &Phone{
		Duration:       1,
		Formants:       [3]float64{550, 1000, 2490},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"æ": &Phone{
		Duration:       1,
		Formants:       [3]float64{690, 1660, 2490},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"u": &Phone{
		Duration:       1,
		Formants:       [3]float64{310, 870, 2250},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ʊ": &Phone{
		Duration:       1,
		Formants:       [3]float64{450, 1030, 2380},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"o": &Phone{
		Duration:       1,
		Formants:       [3]float64{450, 700, 2380},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ɔ": &Phone{
		Duration:       1,
		Formants:       [3]float64{590, 880, 2540},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"a": &Phone{
		Duration:       1,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"p": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{310, 870, 2250},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         false,
	},
	"b": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{310, 870, 2250},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         true,
	},
	"t": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         false,
	},
	"d": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         true,
	},
	"k": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         false,
	},
	"g": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         true,
	},
	"θ": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{500, 1500, 2490},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         false,
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"ð": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{500, 1500, 2490},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         true,
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"f": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{450, 1030, 2380},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         false,
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"v": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{450, 1030, 2380},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         true,
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"s": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         false,
		NoiseFrequency: 2000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"z": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         true,
		NoiseFrequency: 2000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"ʃ": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{280, 2250, 2890},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         false,
		NoiseFrequency: 1300,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"ʒ": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{280, 2250, 2890},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         true,
		NoiseFrequency: 1300,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"h": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
		Voiced:         false,
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"m": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{450, 1030, 2380},
		FormantVolumes: [3]float64{0.7, 0.7, 0.7},
		Consonant:      true,
		Voiced:         true,
		Nasal:          true,
	},
	"n": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{500, 1500, 2490},
		FormantVolumes: [3]float64{0.7, 0.7, 0.7},
		Consonant:      true,
		Voiced:         true,
		Nasal:          true,
	},
	"ŋ": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0.7, 0.7, 0.7},
		Consonant:      true,
		Voiced:         true,
		Nasal:          true,
	},
	"l": &Phone{
		Duration:       0.7,
		Formants:       [3]float64{550, 1000, 2490},
		FormantVolumes: [3]float64{0.7, 0.7, 0.7},
		Consonant:      true,
		Voiced:         true,
	},
	"ɹ": &Phone{
		Duration:       1,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0.5, 0.5, 0.5},
		Consonant:      true,
		Voiced:         true,
	},
	"j": &Phone{
		Duration:       0.2,
		Formants:       [3]float64{280, 2250, 2890},
		FormantVolumes: [3]float64{1, 1, 1},
		Consonant:      true,
	},
	"w": &Phone{
		Duration:       0.2,
		Formants:       [3]float64{310, 870, 2250},
		FormantVolumes: [3]float64{1, 1, 1},
		Consonant:      true,
	},
	"ʔ": &Phone{
		Duration:       0.2,
		Formants:       [3]float64{100, 100, 100},
		FormantVolumes: [3]float64{0, 0, 0},
		Consonant:      true,
	},
}}
