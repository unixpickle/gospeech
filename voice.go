package gospeech

import (
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
			wav.AppendSilence(s, time.Duration(float64(time.Second)/phoneRate))
			lastPhone = nil
			continue
		}
		phone := v.Phones[string(ph)]
		if phone != nil {
			phone.Synthesize(lastPhone, s, phoneRate)
			lastPhone = phone
		}
	}
	return s
}

var DefaultVoice = Voice{Phones: map[string]*Phone{
	"i": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{280, 2250, 2890},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"I": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"e": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{400, 2200, 2890},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ə": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{500, 1500, 2490},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ɛ": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{550, 1770, 2490},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ʌ": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{550, 1000, 2490},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"æ": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{690, 1660, 2490},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"u": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{310, 870, 2250},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ʊ": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{450, 1030, 2380},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"o": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{450, 700, 2380},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ɔ": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{590, 880, 2540},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"a": &Phone{
		Type:           Vowel,
		Duration:       1,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"p": &Phone{
		Type:           Stop,
		Duration:       0.5,
		Formants:       [3]float64{310, 870, 2250},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
	},
	"b": &Phone{
		Type:           Stop,
		Duration:       0.5,
		Formants:       [3]float64{310, 870, 2250},
		FormantVolumes: [3]float64{0, 0, 0},
	},
	"t": &Phone{
		Type:           Stop,
		Duration:       0.5,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
	},
	"d": &Phone{
		Type:           Stop,
		Duration:       0.5,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{0, 0, 0},
	},
	"k": &Phone{
		Type:           Stop,
		Duration:       0.5,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
	},
	"g": &Phone{
		Type:           Stop,
		Duration:       0.5,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0, 0, 0},
	},
	"θ": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{500, 1500, 2490},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"ð": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{500, 1500, 2490},
		FormantVolumes: [3]float64{0, 0, 0},
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"f": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{450, 1030, 2380},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"v": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{450, 1030, 2380},
		FormantVolumes: [3]float64{0, 0, 0},
		NoiseFrequency: 1000,
		NoiseSpread:    200,
		NoiseVolume:    0.3,
	},
	"s": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
		NoiseFrequency: 5000,
		NoiseSpread:    1000,
		NoiseVolume:    1,
	},
	"z": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{400, 1920, 2560},
		FormantVolumes: [3]float64{0, 0, 0},
		NoiseFrequency: 2000,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"ʃ": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{280, 2250, 2890},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
		NoiseFrequency: 1300,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"ʒ": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{280, 2250, 2890},
		FormantVolumes: [3]float64{0, 0, 0},
		NoiseFrequency: 1300,
		NoiseSpread:    200,
		NoiseVolume:    1,
	},
	"h": &Phone{
		Type:           Fricative,
		Duration:       0.7,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
		NoiseFrequency: 600,
		NoiseSpread:    600,
		NoiseVolume:    1,
	},
	"m": &Phone{
		Type:           Nasal,
		Duration:       0.7,
		Formants:       [3]float64{450, 1030, 2380},
		FormantVolumes: [3]float64{0, 0, 0},
	},
	"n": &Phone{
		Type:           Nasal,
		Duration:       0.7,
		Formants:       [3]float64{500, 1500, 2490},
		FormantVolumes: [3]float64{0, 0, 0},
	},
	"ŋ": &Phone{
		Type:           Nasal,
		Duration:       0.7,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0.7, 0.7, 0.7},
	},
	"l": &Phone{
		Type:           Liquid,
		Duration:       0.7,
		Formants:       [3]float64{550, 1000, 2490},
		FormantVolumes: [3]float64{0.7, 0.7, 0.7},
	},
	"ɹ": &Phone{
		Type:           Liquid,
		Duration:       0.7,
		Formants:       [3]float64{710, 1100, 2540},
		FormantVolumes: [3]float64{0.5, 0.5, 0.5},
	},
	"j": &Phone{
		Type:           Glide,
		Duration:       0.2,
		Formants:       [3]float64{280, 2250, 2890},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"w": &Phone{
		Type:           Glide,
		Duration:       0.2,
		Formants:       [3]float64{310, 870, 2250},
		FormantVolumes: [3]float64{1, 1, 1},
	},
	"ʔ": &Phone{
		Type:           Stop,
		Duration:       0.2,
		Formants:       [3]float64{100, 100, 100},
		FormantVolumes: [3]float64{0, 0, 0},
		Unvoiced:       true,
	},
}}
