package tracks

import (
	"math"
	"time"

	"github.com/unixpickle/wav"
)

const sawtoothHarmonicCount = 50

// A SawtoothParameters represents an instantaneous state of a SawtoothTrack.
type SawtoothParameters struct {
	Volume     float64
	Formants   []float64
	Amplitudes []float64
	Strength   float64
}

// NewSawtoothParameters generates a SawtoothParameters filled in with zero values, but with non-nil
// slices.
func NewSawtoothParameters(formantCount int) *SawtoothParameters {
	return &SawtoothParameters{
		Formants:   make([]float64, formantCount),
		Amplitudes: make([]float64, formantCount),
	}
}

// Copy generates a deep copy of the receiver.
func (s *SawtoothParameters) Copy() *SawtoothParameters {
	res := &SawtoothParameters{
		Volume:     s.Volume,
		Formants:   make([]float64, len(s.Formants)),
		Amplitudes: make([]float64, len(s.Amplitudes)),
		Strength:   s.Strength,
	}
	copy(res.Formants, s.Formants)
	copy(res.Amplitudes, s.Amplitudes)
	return res
}

// powerForFrequency returns a number between 0 and 1 indicating how much of a given frequency
// should be included in the wave, based on the frequency's distance from the nearest formant.
func (s *SawtoothParameters) powerForFrequency(freq float64) float64 {
	minDistance := math.Abs(freq - s.Formants[0])
	for i := 1; i < len(s.Formants); i++ {
		minDistance = math.Min(minDistance, math.Abs(freq-s.Formants[i]))
	}

	// This is somewhat arbitrary, but signifies the fact that, as s.Strength grows, frequencies
	// get cut off more sharply as they diverge from a formant.
	minDistance *= s.Strength

	// I picked this formula semi-randomly, but it has the nice property of starting at 1 and
	// decreasing slowly after that.
	return math.Pow(1+minDistance, -4)
}

// A SawtoothTrack generates a sawtooth wave and filters out certain frequencies in it, acting like
// a bandpass filter that creates certain formants.
type SawtoothTrack struct {
	fundamentalFrequency float64
	amplitudeScale       float64
	parts                []*sawtoothTrackPart
}

// NewSawtoothTrack generates a SawtoothTrack with zero duration and a zero'd set of initial
// parameters.
// The fundFreq argument specifies the fundamental frequency for the wave's fourier series.
func NewSawtoothTrack(fundFreq float64, formantCount int) *SawtoothTrack {
	var maxAmplitude float64
	for i := 1; i <= sawtoothHarmonicCount; i++ {
		freq := float64(i) * fundFreq
		maxAmplitude += 1 / freq
	}
	return &SawtoothTrack{
		fundamentalFrequency: fundFreq,
		amplitudeScale:       1 / maxAmplitude,
		parts: []*sawtoothTrackPart{
			&sawtoothTrackPart{
				start: NewSawtoothParameters(formantCount),
				end:   NewSawtoothParameters(formantCount),
			},
		},
	}
}

func (s *SawtoothTrack) Duration() (duration time.Duration) {
	for _, part := range s.parts {
		duration += part.duration
	}
	return
}

func (s *SawtoothTrack) Encode(sampleRate int) []wav.Sample {
	duration := s.Duration()
	var partStartTime time.Duration
	var partIndex int

	res := []wav.Sample{}
	tempParameters := NewSawtoothParameters(len(s.lastPart().end.Amplitudes))
	for {
		secondsElapsed := float64(len(res)) / float64(sampleRate)
		currentTime := time.Duration(float64(time.Second) * secondsElapsed)
		if currentTime >= duration {
			break
		}

		for currentTime >= partStartTime+s.parts[partIndex].duration {
			partStartTime += s.parts[partIndex].duration
			partIndex++
		}

		part := s.parts[partIndex]
		part.parametersAtTime(tempParameters, currentTime-partStartTime)
		sample := s.sample(tempParameters, secondsElapsed)
		res = append(res, wav.Sample(sample))
	}

	return res
}

// Volume returns the volume of the current parameters.
func (s *SawtoothTrack) Volume() float64 {
	return s.lastPart().end.Volume
}

// AdjustVolume elongates the track while adjusting its volume parameter.
func (s *SawtoothTrack) AdjustVolume(volume float64, d time.Duration) {
	newParams := s.Parameters()
	newParams.Volume = volume
	s.AdjustParameters(newParams, d)
}

// Parameters returns a copy of the current parameters.
func (s *SawtoothTrack) Parameters() *SawtoothParameters {
	return s.lastPart().end.Copy()
}

// AdjustParameters elongates the track while adjusting its parameters.
func (s *SawtoothTrack) AdjustParameters(newParams *SawtoothParameters, d time.Duration) {
	part := &sawtoothTrackPart{
		duration: d,
		start:    s.lastPart().end,
		end:      newParams.Copy(),
	}
	s.parts = append(s.parts, part)
}

func (s *SawtoothTrack) lastPart() *sawtoothTrackPart {
	return s.parts[len(s.parts)-1]
}

func (s *SawtoothTrack) sample(params *SawtoothParameters, time float64) float64 {
	var res float64
	for i := 1; i <= sawtoothHarmonicCount; i++ {
		freq := float64(i) * s.fundamentalFrequency
		sinValue := (1 / freq) * math.Sin(math.Pi*2*freq*time)
		power := params.Volume * params.powerForFrequency(freq)
		res += power * sinValue
	}
	return res * s.amplitudeScale
}

type sawtoothTrackPart struct {
	duration time.Duration
	start    *SawtoothParameters
	end      *SawtoothParameters
}

func (s *sawtoothTrackPart) parametersAtTime(out *SawtoothParameters, t time.Duration) {
	fracDone := float64(t) / float64(s.duration)
	out.Volume = fracDone*s.end.Volume + (1-fracDone)*s.start.Volume
	out.Strength = fracDone*s.end.Strength + (1-fracDone)*s.start.Strength
	for i := range out.Amplitudes {
		out.Amplitudes[i] = fracDone*s.end.Amplitudes[i] + (1-fracDone)*s.start.Amplitudes[i]
		out.Formants[i] = fracDone*s.end.Formants[i] + (1-fracDone)*s.start.Formants[i]
	}
}
