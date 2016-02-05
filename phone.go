package gospeech

import (
	"math"
	"math/rand"

	"github.com/unixpickle/wav"
)

// transitionFraction is the fraction of a stop or fricative's
// time that is spent transitioning between phones.
const transitionFraction = 0.5

// consonantPull is the amount of influence consonants should have on
// their neighboring vowels, measured from 0 to 1.
const consonantPull = 0.2

const maxFormantAmplitude = 0.3
const maxNoiseAmplitude = 0.5

type PhoneType int

const (
	Stop PhoneType = iota
	Fricative
	Nasal
	Liquid
	Glide
	Vowel
)

type Phone struct {
	Type     PhoneType
	Unvoiced bool

	// Duration is the length of the phone in relative units.
	// The average Duration for a given voice should be about 1.
	Duration float64

	// For vowels, these fields determine the phone's sound.
	// For consonants, these determine how vowel-consonant transitions
	// should take place.
	Formants       [3]float64
	FormantVolumes [3]float64

	// These fields describe the noise produced by a fricative consonant.
	NoiseFrequency float64
	NoiseSpread    float64
	NoiseVolume    float64
}

func (p *Phone) Synthesize(last *Phone, sound wav.Sound, phoneRate float64) {
	if last == nil {
		p.synthesizeStatic(sound, phoneRate)
		return
	}
	p.synthesizeTransition(last, sound, phoneRate)
	p.synthesizeStatic(sound, phoneRate)
}

func (p *Phone) synthesizeStatic(sound wav.Sound, phoneRate float64) {
	sampleCount := int(float64(sound.SampleRate()) * p.Duration / phoneRate)
	samples := make([]wav.Sample, sampleCount)
	for i := 0; i < sampleCount; i++ {
		secondsElapsed := float64(i) / float64(sound.SampleRate())

		var s wav.Sample
		for j := 0; j < 3; j++ {
			var frequency, amplitude float64
			if p.Type == Nasal {
				frequency = []float64{250, 2500, 3250}[j]
				amplitude = []float64{0.3, 0.1, 0.05}[j]
			} else {
				frequency = p.Formants[j]
				amplitude = p.FormantVolumes[j] * maxFormantAmplitude
			}
			wavValue := math.Sin(math.Pi * 2 * secondsElapsed * frequency)
			s += wav.Sample(wavValue * amplitude)
		}
		if p.NoiseVolume > 0 {
			// TODO: figure out better way of creating noise here.
			noiseValue := math.Sin(math.Pi * 2 * secondsElapsed * p.NoiseFrequency)
			noiseValue *= rand.Float64() * p.NoiseVolume * maxNoiseAmplitude
			s += wav.Sample(noiseValue)
		}
		samples[i] = s
	}
	sound.SetSamples(append(sound.Samples(), samples...))
}

func (p *Phone) synthesizeTransition(last *Phone, sound wav.Sound, phoneRate float64) {
	transitionTime := math.Min(last.Duration, p.Duration) * transitionFraction
	transitionSamples := int(float64(sound.SampleRate()) * transitionTime / phoneRate)

	samples := make([]wav.Sample, transitionSamples)

	for i := 0; i < transitionSamples; i++ {
		fraction := float64(i) / float64(transitionSamples)
		fractionNew := fraction
		if last.Type != Vowel && p.Type == Vowel {
			fractionNew = consonantPull*fraction + fraction*(1-consonantPull)
		} else if p.Type != Vowel && last.Type == Vowel {
			fractionNew = consonantPull * fraction
		}

		var formants [3]float64
		var formantVolumes [3]float64
		for j := 0; j < 3; j++ {
			lastFormant := last.Formants[j]
			newFormant := p.Formants[j]
			lastVolume := last.FormantVolumes[j]
			newVolume := p.FormantVolumes[j]
			formants[j] = lastFormant*(1-fractionNew) + newFormant*fractionNew
			formantVolumes[j] = lastVolume*(1-fractionNew) + newVolume*fractionNew
		}

		var s wav.Sample
		for j := 0; j < 3; j++ {
			frequency := formants[j]
			amplitude := formantVolumes[j] * maxFormantAmplitude
			secondsElapsed := float64(i) / float64(sound.SampleRate())
			wavValue := math.Sin(math.Pi * 2 * secondsElapsed * frequency)
			s += wav.Sample(wavValue * amplitude)
		}
		samples[i] = s
	}

	sound.SetSamples(append(sound.Samples(), samples...))
}
