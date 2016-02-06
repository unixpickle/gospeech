package main

import (
	"time"

	"github.com/unixpickle/gospeech/tracks"
	"github.com/unixpickle/wav"
)

func main() {
	set := tracks.TrackSet{
		"s":  tracks.NewToneTrack(5000, 0, 1000),
		"sd": tracks.NewToneTrack(4500, 0, 1000),
		"dHumm": tracks.TrackSet{
			"Humm1": tracks.NewToneTrack(200, 0, 0),
			"Humm2": tracks.NewToneTrack(300, 0, 0),
			"Humm3": tracks.NewToneTrack(400, 0, 0),
		},
		"a": tracks.TrackSet{
			"F1": tracks.NewToneTrack(710, 0, 0),
			"F2": tracks.NewToneTrack(1100, 0, 0),
			"F3": tracks.NewToneTrack(2540, 0, 0),
		},
		"I": tracks.TrackSet{
			"F1": tracks.NewToneTrack(400, 0, 0),
			"F2": tracks.NewToneTrack(1920, 0, 0),
			"F3": tracks.NewToneTrack(2560, 0, 0),
		},
	}

	set.Continue(time.Second / 3)
	encodeDing(set)
	set.Continue(time.Millisecond * 20)
	encodeDong(set)
	set.Continue(time.Second / 3)

	outputSound := wav.NewPCM8Sound(1, 44100)
	outputSound.SetSamples(set.Encode(44100))
	wav.WriteFile(outputSound, "dingdong.wav")
}

func encodeDing(set tracks.TrackSet) {
	set["sd"].Continue(time.Millisecond * 20)
	set["sd"].AdjustVolume(0.2, time.Millisecond*30)
	set["dHumm"].AdjustVolume(0.1, time.Millisecond*50)
	set["dHumm"].AdjustVolume(0, time.Millisecond*50)
	set["sd"].AdjustVolume(0, time.Millisecond*30)

	set["I"].Continue(time.Millisecond * 80)
	set["I"].AdjustVolume(0.3, time.Millisecond*100)

	set.EvenOut()
	set.Continue(time.Millisecond * 200)

	newFormants := map[string]float64{"F1": 400, "F2": 1800, "F3": 2860}
	for name, track := range set["I"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustAll(newFormants[string(name)], 0.1, 0, time.Millisecond*70)
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

func encodeDong(set tracks.TrackSet) {
	set["sd"].Continue(time.Millisecond * 20)
	set["sd"].AdjustVolume(0.2, time.Millisecond*30)
	set["dHumm"].AdjustVolume(0.1, time.Millisecond*50)
	set["dHumm"].AdjustVolume(0, time.Millisecond*50)
	set["sd"].AdjustVolume(0, time.Millisecond*30)

	set["a"].Continue(time.Millisecond * 80)
	set["a"].AdjustVolume(0.3, time.Millisecond*100)

	set.EvenOut()
	set.Continue(time.Millisecond * 200)

	newFormants := map[string]float64{"F1": 500, "F2": 1000, "F3": 2740}
	for name, track := range set["a"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustAll(newFormants[string(name)], 0.1, 0, time.Millisecond*100)
	}

	newFormants = map[string]float64{"F1": 250, "F2": 2500, "F3": 3250}
	newAmplitudes := map[string]float64{"F1": 0.3, "F2": 0.1, "F3": 0.1}
	for name, track := range set["a"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustAll(newFormants[string(name)],
			newAmplitudes[string(name)], 0, time.Millisecond*50)
	}

	set.EvenOut()
	set.AdjustVolume(0, time.Millisecond*100)
}
