package gospeech

import (
	"time"

	"github.com/unixpickle/gospeech/tracks"
)

// FormantState represents an instantaneous state of three formants.
type FormantState struct {
	Frequencies [3]float64
	Volumes     [3]float64
}

func NewFormantState(f1, v1, f2, v2, f3, v3 float64) FormantState {
	return FormantState{
		Frequencies: [3]float64{f1, f2, f3},
		Volumes:     [3]float64{v1, v2, v3},
	}
}

// A VocalSystem manages speech-like qualities in a TrackSet.
type VocalSystem struct {
	tracks.TrackSet
}

// NewVocalSystem creates a VocalSystem that is currently silent.
func NewVocalSystem() VocalSystem {
	return VocalSystem{tracks.TrackSet{
		"Formants": tracks.TrackSet{
			"F1": tracks.NewToneTrack(400, 0, 0),
			"F2": tracks.NewToneTrack(1000, 0, 0),
			"F3": tracks.NewToneTrack(2000, 0, 0),
		},
		"Turbulence": tracks.TrackSet{
			"S":  tracks.NewToneTrack(5000, 0, 1000),
			"SH": tracks.NewToneTrack(4700, 0, 1000),
			"TH": tracks.NewToneTrack(3500, 0, 1000),
			"B":  tracks.NewToneTrack(4250, 0, 1000),
			"F":  tracks.NewToneTrack(4250, 0, 1000),
			"K":  tracks.NewToneTrack(4250, 0, 1000),
			"H": tracks.TrackSet{
				"F1": tracks.NewToneTrack(1000, 0, 500),
				"F2": tracks.NewToneTrack(2250, 0, 500),
				"F3": tracks.NewToneTrack(2890, 0, 500),
			},
		},
		"ConsonantVoice": tracks.TrackSet{
			"Humm1": tracks.NewToneTrack(200, 0, 0),
			"Humm2": tracks.NewToneTrack(300, 0, 0),
			"Humm3": tracks.NewToneTrack(400, 0, 0),
		},
		"Liquid": tracks.NewToneTrack(500, 0, 0),
	}}
}

// FormantsTrack returns the track corresponding to the formants as a whole.
func (v VocalSystem) FormantsTrack() tracks.TrackSet {
	return v.TrackSet[tracks.TrackID("Formants")].(tracks.TrackSet)
}

// Formants returns the current formant state.
func (v VocalSystem) Formants() FormantState {
	freqs := map[string]float64{}
	volumes := map[string]float64{}
	for name, track := range v.FormantsTrack() {
		tone := track.(*tracks.ToneTrack)
		freqs[string(name)] = tone.Frequency()
		volumes[string(name)] = tone.Volume()
	}
	return FormantState{
		Frequencies: [3]float64{freqs["F1"], freqs["F2"], freqs["F3"]},
		Volumes:     [3]float64{volumes["F1"], volumes["F2"], volumes["F3"]},
	}
}

// AdjustFormants adjusts the formant state over a period of time.
func (v VocalSystem) AdjustFormants(state FormantState, d time.Duration) {
	freqs := map[string]float64{"F1": state.Frequencies[0], "F2": state.Frequencies[1],
		"F3": state.Frequencies[2]}
	volumes := map[string]float64{"F1": state.Volumes[0], "F2": state.Volumes[1],
		"F3": state.Volumes[2]}
	for name, track := range v.FormantsTrack() {
		n := string(name)
		track.(*tracks.ToneTrack).AdjustAll(freqs[n], volumes[n], 0, d)
	}
}

// Turbulence returns the set of track that correspond to different kinds of turbulent airflow,
// such as the hissing sound in "s" or the brief expulsion of air in "t".
func (v VocalSystem) Turbulence() tracks.TrackSet {
	return v.TrackSet[tracks.TrackID("Turbulence")].(tracks.TrackSet)
}

// ConsonantVoice returns the track that corresponds to the humming heard in voiced consonants.
func (v VocalSystem) ConsonantVoice() tracks.Track {
	return v.TrackSet[tracks.TrackID("ConsonantVoice")]
}

// Liquid returns the track that corresponds to the "L" sound.
func (v VocalSystem) Liquid() tracks.Track {
	return v.TrackSet[tracks.TrackID("ConsonantVoice")]
}
