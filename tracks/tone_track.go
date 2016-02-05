package tracks

import (
	"math"
	"math/rand"
	"time"

	"github.com/unixpickle/wav"
)

// A ToneTrack manages a pure tone with optional
// overlaid noise.
type ToneTrack struct {
	currentTime time.Duration
	segments    []*noiseSegment
}

// NewToneTrack generates a zero-length ToneTrack which
// starts with the given tone parameters.
func NewToneTrack(freq, volume, spread float64) *ToneTrack {
	return &ToneTrack{
		currentTime: 0,
		segments: []*noiseSegment{
			&noiseSegment{
				duration:       0,
				startSpread:    spread,
				startFrequency: freq,
				startVolume:    volume,
				endFrequency:   freq,
				endVolume:      volume,
				endSpread:      spread,
			},
		},
	}
}

func (s *ToneTrack) Duration() (res time.Duration) {
	for _, segment := range s.segments {
		res += segment.duration
	}
	return
}

func (s *ToneTrack) Encode(sampleRate int) []wav.Sample {
	duration := s.Duration()
	var segmentStartTime time.Duration
	var segmentIndex int
	var sampleIndex int

	res := []wav.Sample{}
	var sineArgument float64
	for {
		secondsElapsed := float64(sampleIndex) / float64(sampleRate)
		currentTime := time.Duration(float64(time.Second) * secondsElapsed)
		if currentTime >= duration {
			break
		}

		for currentTime >= segmentStartTime+s.segments[segmentIndex].duration {
			segmentStartTime += s.segments[segmentIndex].duration
			segmentIndex++
		}

		segment := s.segments[segmentIndex]
		freq, volume, spread := segment.infoAtTime(currentTime - segmentStartTime)
		sample := math.Sin(sineArgument) * volume
		res = append(res, wav.Sample(sample))

		freq += rand.NormFloat64() * spread
		sineArgument += math.Pi * 2 * freq / float64(sampleRate)
		for sineArgument > math.Pi*2 {
			sineArgument -= math.Pi * 2
		}

		sampleIndex++
	}

	return res
}

// Continue elongates the tone without modifying it.
func (s *ToneTrack) Continue(duration time.Duration) {
	lastSeg := s.lastSegment()
	if lastSeg.static() {
		lastSeg.duration += duration
	} else {
		seg := &noiseSegment{
			duration:       duration,
			startFrequency: lastSeg.endFrequency,
			startVolume:    lastSeg.endVolume,
			endFrequency:   lastSeg.endFrequency,
			endVolume:      lastSeg.endVolume,
		}
		s.segments = append(s.segments, seg)
	}
}

// Volume returns the tone's current amplitude.
func (s *ToneTrack) Volume() float64 {
	return s.lastSegment().endVolume
}

// AdjustVolume elongates the track while adjusting the tone's frequency.
func (s *ToneTrack) AdjustVolume(newVolume float64, duration time.Duration) {
	s.AdjustAll(s.Frequency(), newVolume, s.Spread(), duration)
}

// Frequency returns the tone's current frequency.
func (s *ToneTrack) Frequency() float64 {
	return s.lastSegment().endFrequency
}

// AdjustFrequency elongates the track while adjusting the tone's frequency.
func (s *ToneTrack) AdjustFrequency(newFrequency float64, duration time.Duration) {
	s.AdjustAll(newFrequency, s.Volume(), s.Spread(), duration)
}

// Spread returns the tone's random spread.
func (s *ToneTrack) Spread() float64 {
	return s.lastSegment().endSpread
}

// AdjustSpread elongates the track while adjusting the tone's random spread.
func (s *ToneTrack) AdjustSpread(spread float64, duration time.Duration) {
	s.AdjustAll(s.Frequency(), s.Volume(), spread, duration)
}

// AdjustAll elongates the track by while adjusting the tone's characteristics.
func (s *ToneTrack) AdjustAll(freq, volume, spread float64, duration time.Duration) {
	lastSeg := s.lastSegment()
	seg := &noiseSegment{
		duration:       duration,
		startSpread:    lastSeg.endSpread,
		startFrequency: lastSeg.endFrequency,
		startVolume:    lastSeg.endVolume,
		endFrequency:   freq,
		endVolume:      volume,
		endSpread:      spread,
	}
	s.segments = append(s.segments, seg)
}

func (s *ToneTrack) lastSegment() *noiseSegment {
	return s.segments[len(s.segments)-1]
}

type noiseSegment struct {
	duration       time.Duration
	startSpread    float64
	startFrequency float64
	startVolume    float64
	endSpread      float64
	endFrequency   float64
	endVolume      float64
}

func (s *noiseSegment) static() bool {
	return s.startFrequency == s.endFrequency &&
		s.startVolume == s.endVolume &&
		s.startSpread == s.endSpread
}

func (s *noiseSegment) infoAtTime(t time.Duration) (freq, vol, spread float64) {
	fracDone := float64(t) / float64(s.duration)
	freq = fracDone*s.endFrequency + (1-fracDone)*s.startFrequency
	vol = fracDone*s.endVolume + (1-fracDone)*s.startVolume
	spread = fracDone*s.endSpread + (1-fracDone)*s.startSpread
	return
}
