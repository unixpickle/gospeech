package main

import (
	"time"

	"github.com/unixpickle/gospeech/tracks"
	"github.com/unixpickle/wav"
)

func main() {
	set := tracks.TrackSet{
		"i": tracks.TrackSet{
			"F1": tracks.NewToneTrack(280, 0, 0),
			"F2": tracks.NewToneTrack(2250, 0, 0),
			"F3": tracks.NewToneTrack(2890, 0, 0),
		},
		"I": tracks.TrackSet{
			"F1": tracks.NewToneTrack(400, 0, 0),
			"F2": tracks.NewToneTrack(1920, 0, 0),
			"F3": tracks.NewToneTrack(2560, 0, 0),
		},
		"Aspiration": tracks.TrackSet{
			"F1": tracks.NewToneTrack(280, 0, 280),
			"F2": tracks.NewToneTrack(2250, 0, 500),
			"F3": tracks.NewToneTrack(2890, 0, 500),
		},
		"z": tracks.TrackSet{
			"Humm1":      tracks.NewToneTrack(200, 0, 0),
			"Humm2":      tracks.NewToneTrack(300, 0, 0),
			"Humm3":      tracks.NewToneTrack(400, 0, 0),
			"Turbulence": tracks.NewToneTrack(5000, 0, 1000),
		},
		"k": tracks.TrackSet{
			"F1": tracks.NewToneTrack(600, 0, 200),
			"F2": tracks.NewToneTrack(500, 0, 500),
			"F3": tracks.NewToneTrack(800, 0, 500),
		},
		"u": tracks.TrackSet{
			"F1": tracks.NewToneTrack(310, 0, 0),
			"F2": tracks.NewToneTrack(870, 0, 0),
			"F3": tracks.NewToneTrack(2250, 0, 0),
		},
		"l": tracks.NewToneTrack(500, 0, 0),
	}

	set.Continue(time.Second / 3)
	encodeHe(set)
	set.Continue(time.Second / 3)
	encodeIs(set)
	set.Continue(time.Second / 3)
	encodeCool(set)
	set.Continue(time.Second / 3)

	outputSound := wav.NewPCM8Sound(1, 44100)
	outputSound.SetSamples(set.Encode(44100))
	wav.WriteFile(outputSound, "he_is_cool.wav")
}

func encodeHe(set tracks.TrackSet) {
	set["Aspiration"].AdjustVolume(0.1, time.Millisecond*200)
	set.EvenOut()

	set["Aspiration"].AdjustVolume(0, time.Millisecond*100)
	set["i"].AdjustVolume(0.3, time.Millisecond*100)
	set.EvenOut()

	set.Continue(time.Second / 5)
	set.AdjustVolume(0, time.Millisecond*100)
}

func encodeIs(set tracks.TrackSet) {
	set["I"].AdjustVolume(0.3, time.Millisecond*100)
	set.EvenOut()
	set.Continue(time.Millisecond * 100)

	formantTargets := map[string]float64{"F1": 400, "F2": 1920, "F3": 2860}
	for name, formant := range set["I"].(tracks.TrackSet) {
		toneTrack := formant.(*tracks.ToneTrack)
		toneTrack.AdjustAll(formantTargets[string(name)], 0, 0, time.Millisecond*100)
	}
	zTrack := set["z"].(tracks.TrackSet)
	zTrack[tracks.TrackID("Humm1")].Continue(time.Millisecond)
	zTrack[tracks.TrackID("Humm1")].AdjustVolume(0.03, time.Millisecond*50)
	zTrack[tracks.TrackID("Humm2")].Continue(time.Millisecond)
	zTrack[tracks.TrackID("Humm2")].AdjustVolume(0.03, time.Millisecond*50)
	zTrack[tracks.TrackID("Humm3")].Continue(time.Millisecond)
	zTrack[tracks.TrackID("Humm3")].AdjustVolume(0.03, time.Millisecond*50)
	zTrack[tracks.TrackID("Turbulence")].AdjustVolume(0.05, time.Millisecond*150)
	set.EvenOut()

	set.Continue(time.Second / 10)
	set.AdjustVolume(0, time.Millisecond*100)
}

func encodeCool(set tracks.TrackSet) {
	set["k"].AdjustVolume(0.3, time.Millisecond*2)
	set["k"].Continue(time.Millisecond * 10)
	set["k"].AdjustVolume(0, time.Millisecond*100)
	set.EvenOut()

	set["u"].AdjustVolume(0.3, time.Millisecond*130)
	set["u"].Continue(time.Millisecond * 300)
	formantTargets := map[string]float64{"F1": 550, "F2": 1000, "F3": 2490}
	for name, formant := range set["u"].(tracks.TrackSet) {
		toneTrack := formant.(*tracks.ToneTrack)
		toneTrack.AdjustAll(formantTargets[string(name)], 0, 0, time.Millisecond*150)
	}
	set["l"].Continue(time.Millisecond * 350)
	set["l"].AdjustVolume(0.05, time.Millisecond*100)
	set.EvenOut()
	set.AdjustVolume(0, time.Millisecond*100)
}
