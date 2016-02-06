package tracks

import (
	"time"

	"github.com/unixpickle/wav"
)

// A Track is a mutable stream of audio data.
//
// At any given time, a Track has a "current sound", which may be
// a tone, a type of pink noise, etc.
// It is possible to elongate a track by producing more of its
// current sound, or to adjust the track's current sound.
//
// Different concrete implementations of Track will produce different
// classes of sounds.
type Track interface {
	Duration() time.Duration
	Encode(sampleRate int) []wav.Sample

	// Continue elongates the track with the current sound.
	Continue(duration time.Duration)

	// Volume returns the average volume of the current sound.
	Volume() float64

	// AdjustVolume elongates the track while simultaneously
	// adjusting the volume of the current sound.
	AdjustVolume(newVolume float64, transitionTime time.Duration)
}

// A TrackID is a string used to identify tracks in a TrackSet.
type TrackID string

// A TrackSet is a track which manages other Tracks in bulk.
type TrackSet map[TrackID]Track

// ExcludeTracks returns a TrackSet that does not contain tracks
// with the given track IDs.
func (t TrackSet) ExcludeTracks(ids ...TrackID) TrackSet {
	res := TrackSet{}
OuterLoop:
	for id, val := range t {
		for _, x := range ids {
			if x == id {
				continue OuterLoop
			}
		}
		res[id] = val
	}
	return res
}

// Duration returns the duration of the longest track in the set.
func (t TrackSet) Duration() (maxDur time.Duration) {
	for _, track := range t {
		if d := track.Duration(); d > maxDur {
			maxDur = d
		}
	}
	return
}

// Encode generates samples by encoding every track in the set and
// summing up the signals.
func (t TrackSet) Encode(sampleRate int) (res []wav.Sample) {
	sampleCount := 0
	encodedTracks := make([][]wav.Sample, 0, len(t))
	for _, track := range t {
		encodedTrack := track.Encode(sampleRate)
		encodedTracks = append(encodedTracks, encodedTrack)
		if len(encodedTrack) > sampleCount {
			sampleCount = len(encodedTrack)
		}
	}

	res = make([]wav.Sample, sampleCount)
	for i := range res {
		for _, enc := range encodedTracks {
			if i >= len(enc) {
				continue
			}
			res[i] += enc[i]
		}
	}

	return
}

// Continue elongates all of the set's tracks by a given duration.
func (t TrackSet) Continue(duration time.Duration) {
	for _, track := range t {
		track.Continue(duration)
	}
}

// Volume returns the average volume of the tracks in this set.
func (t TrackSet) Volume() float64 {
	var sum float64
	for _, track := range t {
		sum += track.Volume()
	}
	return sum / float64(len(t))
}

// AdjustVolume elongates all of the tracks while simultaneously
// adjusting their volumes to the given value.
func (t TrackSet) AdjustVolume(newVolume float64, duration time.Duration) {
	for _, track := range t {
		track.AdjustVolume(newVolume, duration)
	}
}
