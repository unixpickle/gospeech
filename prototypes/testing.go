package main

import (
	"time"

	"github.com/unixpickle/gospeech/tracks"
	"github.com/unixpickle/wav"
)

func main() {
	set := tracks.TrackSet{
		"s": tracks.NewToneTrack(5000, 0, 1000),
		"ɛ": tracks.TrackSet{
			"F1": tracks.NewToneTrack(500, 0, 0),
			"F2": tracks.NewToneTrack(1770, 0, 0),
			"F3": tracks.NewToneTrack(2490, 0, 0),
		},
		"I": tracks.TrackSet{
			"F1": tracks.NewToneTrack(400, 0, 0),
			"F2": tracks.NewToneTrack(1920, 0, 0),
			"F3": tracks.NewToneTrack(2560, 0, 0),
		},
		"Nasal": tracks.TrackSet{
			"F1": tracks.NewToneTrack(250, 0, 0),
			"F2": tracks.NewToneTrack(2500, 0, 0),
			"F3": tracks.NewToneTrack(3250, 0, 0),
		},
	}

	set.Continue(time.Second / 3)
	encodeTest(set)
	encodeIng(set)
	set.Continue(time.Second / 3)

	outputSound := wav.NewPCM8Sound(1, 44100)
	outputSound.SetSamples(set.Encode(44100))
	wav.WriteFile(outputSound, "testing.wav")
}

func encodeTest(set tracks.TrackSet) {
	set["s"].AdjustVolume(0.5, time.Millisecond*3)
	set.EvenOut()
	set.Continue(time.Millisecond * 30)
	set.AdjustVolume(0, time.Millisecond*20)
	set["ɛ"].AdjustVolume(0.3333, time.Millisecond*100)
	set.EvenOut()
	set.Continue(time.Millisecond * 300)
	set.AdjustVolume(0, time.Millisecond*100)
	set["s"].AdjustVolume(0.5, time.Millisecond*100)
	set.EvenOut()
	set.AdjustVolume(0, 0)
	set.Continue(time.Millisecond * 50)
	set["s"].AdjustVolume(0.5, 0)
	set.EvenOut()
	set.AdjustVolume(0, time.Millisecond*30)
}

func encodeIng(set tracks.TrackSet) {
	set["I"].AdjustVolume(0.3, time.Millisecond*50)
	set.EvenOut()
	set.Continue(time.Millisecond * 300)

	newFormants := map[string]float64{"F1": 400, "F2": 1800, "F3": 2860}
	for name, track := range set["I"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustAll(newFormants[string(name)], 0.1, 0, time.Millisecond*100)
	}

	newFormants = map[string]float64{"F1": 250, "F2": 2500, "F3": 3250}
	newAmplitudes := map[string]float64{"F1": 0.3, "F2": 0.1, "F3": 0.1}
	for name, track := range set["I"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustAll(newFormants[string(name)],
			newAmplitudes[string(name)], 0, time.Millisecond*30)
	}

	set.EvenOut()
	set.AdjustVolume(0, time.Millisecond*100)
}
