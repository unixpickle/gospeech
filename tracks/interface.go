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

type TrackID string

// A TrackSet facilitates bulk manipulation of Tracks.
type TrackSet map[TrackID]Track

// ExcludeTracks returns a TrackSet that does not contain tracks
// with the given track IDs.
func (t TrackSet) ExcludeTracks(ids []TrackID) TrackSet {
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

// Continue elongates all of the set's tracks by a given duration.
func (t TrackSet) Continue(duration time.Duration) {
	for _, track := range t {
		track.Continue(duration)
	}
}

// AdjustVolume elongates all of the tracks while simultaneously
// adjusting their volumes.
func (t TrackSet) AdjustVolume(newVolume float64, duration time.Duration) {
	for _, track := range t {
		track.AdjustVolume(newVolume, duration)
	}
}
