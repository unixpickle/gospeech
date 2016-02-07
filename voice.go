package gospeech

import (
	"time"

	"github.com/unixpickle/wav"
)

type Voice struct {
	Phones map[string]Phone
}

func (v Voice) Synthesize(ipaString string) wav.Sound {
	vocalSystem := NewVocalSystem()

	words := [][]Phone{}
	word := []Phone{}

	for _, ph := range []rune(ipaString) {
		if ph == ' ' {
			words = append(words, word)
			word = []Phone{}
		} else {
			if phone := v.Phones[string(ph)]; phone != nil {
				word = append(word, phone)
			}
		}
	}

	if len(word) > 0 {
		words = append(words, word)
	}

	for _, word := range words {
		for i, phone := range word {
			var lastPhone, nextPhone Phone
			if i > 0 {
				lastPhone = word[i-1]
			}
			if i < len(word)-1 {
				nextPhone = word[i+1]
			}
			phone.EncodeBeginning(vocalSystem, lastPhone, nextPhone)
		}
		vocalSystem.AdjustVolume(0, time.Millisecond*50)
		vocalSystem.Continue(time.Millisecond * 300)
	}

	s := wav.NewPCM8Sound(1, 44100)
	s.SetSamples(vocalSystem.Encode(44100))
	return s
}

var DefaultVoice = Voice{
	Phones: map[string]Phone{
		"i": Vowel{
			Formants: NewFormantState(280, 0.3, 2250, 0.3, 2890, 0.3),
			Duration: time.Millisecond * 200,
		},
		"I": Vowel{
			Formants: NewFormantState(400, 0.3, 1920, 0.3, 2560, 0.3),
			Duration: time.Millisecond * 200,
		},
		"e": Vowel{
			Formants: NewFormantState(400, 0.3, 2200, 0.3, 2890, 0.3),
			Duration: time.Millisecond * 200,
		},
		"ə": Vowel{
			Formants: NewFormantState(500, 0.3, 1500, 0.3, 2490, 0.3),
			Duration: time.Millisecond * 200,
		},
		"ɛ": Vowel{
			Formants: NewFormantState(550, 0.3, 1770, 0.3, 2490, 0.3),
			Duration: time.Millisecond * 200,
		},
		"ʌ": Vowel{
			Formants: NewFormantState(550, 0.3, 1000, 0.3, 2490, 0.3),
			Duration: time.Millisecond * 200,
		},
		"æ": Vowel{
			Formants: NewFormantState(690, 0.3, 1660, 0.3, 2490, 0.3),
			Duration: time.Millisecond * 200,
		},
		"u": Vowel{
			Formants: NewFormantState(310, 0.3, 870, 0.3, 2250, 0.3),
			Duration: time.Millisecond * 200,
		},
		"ʊ": Vowel{
			Formants: NewFormantState(450, 0.3, 1030, 0.3, 2380, 0.3),
			Duration: time.Millisecond * 200,
		},
		"o": Vowel{
			Formants: NewFormantState(450, 0.3, 700, 0.3, 2380, 0.3),
			Duration: time.Millisecond * 200,
		},
		"ɔ": Vowel{
			Formants: NewFormantState(590, 0.3, 880, 0.3, 2540, 0.3),
			Duration: time.Millisecond * 200,
		},
		"a": Vowel{
			Formants: NewFormantState(710, 0.3, 1100, 0.3, 2540, 0.3),
			Duration: time.Millisecond * 200,
		},
		"p": BilabialPlosive{Voiced: false},
		"b": BilabialPlosive{Voiced: true},
		"t": AlveolarPlosive{Voiced: false},
		"d": AlveolarPlosive{Voiced: true},
		"k": VelarPlosive{Voiced: false},
		"g": VelarPlosive{Voiced: true},
		"m": Nasal{
			Type:     "m",
			Formants: NewFormantState(250, 0.3, 2500, 0.1, 3250, 0.1),
		},
		"n": Nasal{
			Type:     "n",
			Formants: NewFormantState(250, 0.3, 2500, 0.1, 3250, 0.1),
		},
		"ŋ": Nasal{
			Type:     "ŋ",
			Formants: NewFormantState(250, 0.3, 2500, 0.1, 3250, 0.1),
		},
		"θ": Fricative{
			Type:   "TH",
			Voiced: false,
		},
		"ð": Fricative{
			Type:   "TH",
			Voiced: true,
		},
		"f": Fricative{
			Type:   "F",
			Voiced: false,
		},
		"v": Fricative{
			Type:   "F",
			Voiced: true,
		},
		"s": Fricative{
			Type:   "S",
			Voiced: false,
		},
		"z": Fricative{
			Type:   "S",
			Voiced: true,
		},
		"ʃ": Fricative{
			Type:   "SH",
			Voiced: false,
		},
		"ʒ": Fricative{
			Type:   "SH",
			Voiced: true,
		},
		"h": Fricative{
			Type:   "H",
			Voiced: false,
		},
		"l": LateralLiquid{},
		"ɹ": RetroflexLiquid{Formants: NewFormantState(500, 0.3, 1500, 0.3, 1900, 0.3)},
		"j": Vowel{
			Formants: NewFormantState(280, 0.3, 2250, 0.3, 2890, 0.3),
			Duration: time.Millisecond * 60,
		},
		"w": Vowel{
			Formants: NewFormantState(310, 0.3, 870, 0.3, 2250, 0.3),
			Duration: time.Millisecond * 60,
		},
		"ʔ": GlottalStop{},
		"ɾ": AlveolarPlosive{Voiced: false},
	},
}
