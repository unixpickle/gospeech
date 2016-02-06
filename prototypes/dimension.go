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
		"ɛ": tracks.TrackSet{
			"F1": tracks.NewToneTrack(500, 0, 0),
			"F2": tracks.NewToneTrack(1770, 0, 0),
			"F3": tracks.NewToneTrack(2490, 0, 0),
		},
		"sh": tracks.TrackSet{
			"1": tracks.NewToneTrack(5000, 0, 1000),
			"2": tracks.NewToneTrack(4000, 0, 2000),
		},
		"ə": tracks.TrackSet{
			"F1": tracks.NewToneTrack(500, 0, 0),
			"F2": tracks.NewToneTrack(1500, 0, 0),
			"F3": tracks.NewToneTrack(2490, 0, 0),
		},
	}

	set.Continue(time.Second / 3)
	encodeDaIm(set)
	encodeEn(set)
	encodeShun(set)
	set.Continue(time.Second / 3)

	outputSound := wav.NewPCM8Sound(1, 44100)
	outputSound.SetSamples(set.Encode(44100))
	wav.WriteFile(outputSound, "dimension.wav")
}

func encodeDaIm(set tracks.TrackSet) {
	set["sd"].Continue(time.Millisecond * 20)
	set["sd"].AdjustVolume(0.2, time.Millisecond*30)
	set["dHumm"].AdjustVolume(0.1, time.Millisecond*50)
	set["dHumm"].AdjustVolume(0, time.Millisecond*50)
	set["sd"].AdjustVolume(0, time.Millisecond*30)

	set["a"].Continue(time.Millisecond * 80)
	set["a"].AdjustVolume(0.3, time.Millisecond*100)

	set.EvenOut()

	set["a"].AdjustVolume(0.3, time.Millisecond*150)

	newFormants := map[string]float64{"F1": 450, "F2": 1700, "F3": 2560}
	for name, track := range set["a"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustFrequency(newFormants[string(name)], time.Millisecond*200)
	}

	newFormants = map[string]float64{"F1": 250, "F2": 2500, "F3": 3250}
	newAmplitudes := map[string]float64{"F1": 0.4, "F2": 0.13, "F3": 0.13}
	for name, track := range set["a"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustAll(newFormants[string(name)],
			newAmplitudes[string(name)], 0, time.Millisecond*30)
	}

	set.EvenOut()
	set.Continue(time.Millisecond * 50)
}

func encodeEn(set tracks.TrackSet) {
	set.Continue(time.Millisecond * 150)
	set["a"].AdjustVolume(0, time.Millisecond*100)
	set["ɛ"].AdjustVolume(0.3, time.Millisecond*100)
	set.EvenOut()
	set.Continue(time.Millisecond * 150)

	newFormants := map[string]float64{"F1": 480, "F2": 1830, "F3": 2560}
	for name, track := range set["ɛ"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustFrequency(newFormants[string(name)], time.Millisecond*70)
	}

	newFormants = map[string]float64{"F1": 250, "F2": 2500, "F3": 3250}
	newAmplitudes := map[string]float64{"F1": 0.3, "F2": 0.1, "F3": 0.1}
	for name, track := range set["ɛ"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustAll(newFormants[string(name)],
			newAmplitudes[string(name)], 0, time.Millisecond*20)
	}

	set.EvenOut()
}

func encodeShun(set tracks.TrackSet) {
	set.Continue(time.Millisecond * 130)
	set["ɛ"].AdjustVolume(0, time.Millisecond*50)
	set["sh"].Continue(time.Millisecond * 10)
	set["sh"].AdjustVolume(0.2, time.Millisecond*70)
	set["sh"].AdjustVolume(0, time.Millisecond*100)
	set["ə"].Continue(time.Millisecond * 60)
	set["ə"].AdjustVolume(0.3, time.Millisecond*150)
	set.EvenOut()

	newFormants := map[string]float64{"F1": 480, "F2": 1400, "F3": 2560}
	for name, track := range set["ə"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustFrequency(newFormants[string(name)], time.Millisecond*70)
	}

	newFormants = map[string]float64{"F1": 250, "F2": 2500, "F3": 3250}
	newAmplitudes := map[string]float64{"F1": 0.3, "F2": 0.1, "F3": 0.1}
	for name, track := range set["ə"].(tracks.TrackSet) {
		track.(*tracks.ToneTrack).AdjustAll(newFormants[string(name)],
			newAmplitudes[string(name)], 0, time.Millisecond*20)
	}

	set.EvenOut()
	set.Continue(time.Millisecond * 100)
	set.AdjustVolume(0, time.Millisecond*100)
}
