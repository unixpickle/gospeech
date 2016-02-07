package gospeech

import (
	"time"

	"github.com/unixpickle/gospeech/tracks"
)

type Phone interface {
	// EncodeBeginning encodes the beginning of the phone into the given vocal system.
	// This should also encode the transition from the previous phone, if applicable.
	// The lastPhone argument will be nil if this is the first phone in the word.
	// The nextPhone argument will be nil if this is the last phone in the word.
	EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone)

	// FormantPull tells the next phone how its formants should be modified initially.
	// It returns the initial formants for the next phone, given the steady-state formants of said
	// phone.
	FormantPull(nextFormant FormantState) FormantState

	// TransitionTime is the amount of time that the formants of the next phone should take to reach
	// steady state after starting at the return value of FormantPull().
	TransitionTime() time.Duration
}

type Vowel struct {
	Formants FormantState
	Duration time.Duration
}

func (v Vowel) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	if lastPhone != nil {
		startFormant := lastPhone.FormantPull(v.Formants)
		system.AdjustFormants(startFormant, v.Duration/6)
		system.AdjustFormants(v.Formants, lastPhone.TransitionTime())
	} else {
		startFormants := v.Formants
		startFormants.Volumes = [3]float64{}
		system.AdjustFormants(startFormants, v.Duration/6)
		system.AdjustFormants(v.Formants, v.Duration/2)
	}
	system.Turbulence().AdjustVolume(0, v.Duration/4)
	system.ConsonantVoice().AdjustVolume(0, v.Duration/3)
	system.Liquid().AdjustVolume(0, v.Duration/3)
	system.FormantsTrack().Continue(v.Duration / 3)
	system.EvenOut()
}

func (v Vowel) FormantPull(nextFormant FormantState) FormantState {
	return v.Formants
}

func (v Vowel) TransitionTime() time.Duration {
	return v.Duration / 2
}

// A BilabialPlosive represents a "b" or "p" sound.
type BilabialPlosive struct {
	Voiced bool
}

func (b BilabialPlosive) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	if system.FormantsTrack().Volume() > 0 {
		endFormant := b.previousFormantPull(system.Formants())
		system.AdjustFormants(endFormant, time.Millisecond*30)
	}
	system.Turbulence().AdjustVolume(0, time.Millisecond*30)
	system.ConsonantVoice().AdjustVolume(0, time.Millisecond*30)
	system.Liquid().AdjustVolume(0, time.Millisecond*30)
	system.Continue(time.Millisecond * 50)
	if b.Voiced {
		system.ConsonantVoice().AdjustVolume(0.1, time.Millisecond*10)
	}
	turbulence := system.Turbulence()[tracks.TrackID("P")]
	turbulence.AdjustVolume(0.3, time.Millisecond*3)
	turbulence.Continue(time.Millisecond * 30)
	system.EvenOut()
}

func (b BilabialPlosive) FormantPull(end FormantState) FormantState {
	for i := 0; i < 3; i++ {
		end.Volumes[i] *= 1
		end.Frequencies[i] *= 0.9
	}
	return end
}

func (v BilabialPlosive) TransitionTime() time.Duration {
	return time.Millisecond * 10
}

func (b BilabialPlosive) previousFormantPull(last FormantState) FormantState {
	for i := 0; i < 3; i++ {
		last.Volumes[i] *= 0
		last.Frequencies[i] *= 0.9
	}
	return last
}

// An AlveolarPlosive represents a "t" or "d" sound.
type AlveolarPlosive struct {
	Voiced bool

	// ContinueToNext indicates that the next phone is an "s" or something like that, in which case
	// this phone needn't terminate its sound.
	ContinueToNext bool
}

func (a AlveolarPlosive) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	if system.FormantsTrack().Volume() > 0 {
		system.AdjustFormants(a.FormantPull(system.Formants()), time.Millisecond*50)
	}
	system.Turbulence().AdjustVolume(0, time.Millisecond*50)
	system.ConsonantVoice().AdjustVolume(0, time.Millisecond*50)
	system.Liquid().AdjustVolume(0, time.Millisecond*50)

	system.Continue(time.Millisecond * 10)
	if a.Voiced {
		system.ConsonantVoice().AdjustVolume(0.3, time.Millisecond*50)
		if !a.ContinueToNext {
			system.ConsonantVoice().AdjustVolume(0, time.Millisecond*50)
		}
	}
	turbulence := system.Turbulence()[tracks.TrackID("S")]
	turbulence.Continue(time.Millisecond * 20)
	turbulence.AdjustVolume(0.3, time.Millisecond*3)
	turbulence.Continue(time.Millisecond * 20)
	if !a.ContinueToNext {
		turbulence.AdjustVolume(0, time.Millisecond*20)
	}
	system.EvenOut()
}

func (a AlveolarPlosive) FormantPull(end FormantState) FormantState {
	// TODO: figure out some values here. It probably varies per vowel.
	end.Volumes = [3]float64{}
	return end
}

func (v AlveolarPlosive) TransitionTime() time.Duration {
	return time.Millisecond * 75
}

// A VelarPlosive represents a "k" or "g" sound.
type VelarPlosive struct {
	Voiced bool
}

func (v VelarPlosive) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	if system.FormantsTrack().Volume() > 0 {
		system.AdjustFormants(v.FormantPull(system.Formants()), time.Millisecond*50)
	}
	system.Turbulence().AdjustVolume(0, time.Millisecond*50)
	system.ConsonantVoice().AdjustVolume(0, time.Millisecond*50)
	system.Liquid().AdjustVolume(0, time.Millisecond*50)

	if v.Voiced {
		system.ConsonantVoice().AdjustVolume(0.1, time.Millisecond*20)
	}
	turbulence := system.Turbulence()[tracks.TrackID("K")]
	turbulence.Continue(time.Millisecond * 20)
	turbulence.AdjustVolume(0.2, time.Millisecond*5)
	turbulence.Continue(time.Millisecond * 20)
	turbulence.AdjustVolume(0, time.Millisecond*10)
	system.EvenOut()
}

func (v VelarPlosive) FormantPull(end FormantState) FormantState {
	res := end
	res.Volumes = [3]float64{}
	res.Frequencies[0] = end.Frequencies[0]*0.9 + end.Frequencies[1]*0.1
	res.Frequencies[1] = end.Frequencies[1]*0.9 + end.Frequencies[0]*0.1
	return res
}

func (v VelarPlosive) TransitionTime() time.Duration {
	return time.Millisecond * 30
}

// A Nasal represents an "n", "m", or "ng" sound.
type Nasal struct {
	// Type is either "n", "m", or "ng", and dictates the formant pull technique.
	Type     string
	Formants FormantState
}

func (n Nasal) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	if system.FormantsTrack().Volume() > 0 {
		system.AdjustFormants(n.FormantPull(system.Formants()), time.Millisecond*50)
	}
	system.Turbulence().AdjustVolume(0, time.Millisecond*50)
	system.ConsonantVoice().AdjustVolume(0, time.Millisecond*50)
	system.Liquid().AdjustVolume(0, time.Millisecond*50)

	system.AdjustFormants(n.Formants, time.Millisecond*50)
	system.EvenOut()
	system.Continue(time.Millisecond * 100)
}

func (n Nasal) FormantPull(end FormantState) FormantState {
	// TODO: figure out a more accurate pull technique for nasals.
	var pullFrequencies [3]float64
	switch n.Type {
	case "ng":
		return VelarPlosive{}.FormantPull(end)
	case "n":
		pullFrequencies = [3]float64{500, 1500, 2490}
	case "m":
		pullFrequencies = [3]float64{550, 1000, 2490}
	}
	end.Volumes = [3]float64{}
	for i := 0; i < 3; i++ {
		end.Frequencies[i] = end.Frequencies[i]*0.9 + pullFrequencies[i]*0.1
	}
	return end
}

func (v Nasal) TransitionTime() time.Duration {
	return time.Millisecond * 20
}

// A Fricative represents any consonant which is characterized by turbulent airflow.
type Fricative struct {
	// Type is "F", "TH", "S", "SH", or "H", indicating which kind of fricative this is.
	Type string

	Voiced bool
}

func (f Fricative) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	if system.FormantsTrack().Volume() > 0 {
		system.AdjustFormants(f.FormantPull(system.Formants()), time.Millisecond*50)
	}
	system.Turbulence().ExcludeTracks(tracks.TrackID(f.Type)).AdjustVolume(0, time.Millisecond*50)
	system.Liquid().AdjustVolume(0, time.Millisecond*50)

	if f.Voiced {
		system.ConsonantVoice().AdjustVolume(0.3, time.Millisecond*100)
	} else {
		system.ConsonantVoice().AdjustVolume(0, time.Millisecond*50)
	}

	turbulence := system.Turbulence()[tracks.TrackID(f.Type)]
	turbulence.AdjustVolume(0.3, time.Millisecond*100)
	system.EvenOut()
}

func (f Fricative) FormantPull(end FormantState) FormantState {
	end.Volumes = [3]float64{}
	return end
}

func (v Fricative) TransitionTime() time.Duration {
	return time.Millisecond * 100
}

// An RetroflexLiquid represents an "r" sound (without trill).
type RetroflexLiquid struct {
	Formants FormantState
}

func (r RetroflexLiquid) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	vowel := Vowel{Formants: r.Formants, Duration: time.Millisecond * 150}
	vowel.EncodeBeginning(system, lastPhone, nextPhone)
}

func (r RetroflexLiquid) FormantPull(end FormantState) FormantState {
	return r.Formants
}

func (v RetroflexLiquid) TransitionTime() time.Duration {
	return time.Millisecond * 100
}

// A LateralLiquid represents an "l" sound.
type LateralLiquid struct{}

func (l LateralLiquid) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	if system.FormantsTrack().Volume() > 0 {
		system.AdjustFormants(l.FormantPull(system.Formants()), time.Millisecond*50)
	}
	system.Turbulence().AdjustVolume(0, time.Millisecond*50)
	system.ConsonantVoice().AdjustVolume(0, time.Millisecond*50)
	system.Liquid().AdjustVolume(0.3, time.Millisecond*50)
	system.EvenOut()
}

func (l LateralLiquid) FormantPull(end FormantState) FormantState {
	end.Volumes = [3]float64{}
	pullFrequencies := [3]float64{450, 1030, 2380}
	for i := 0; i < 3; i++ {
		end.Frequencies[i] = end.Frequencies[i]*0.9 + pullFrequencies[i]*0.1
	}
	return end
}

func (v LateralLiquid) TransitionTime() time.Duration {
	return time.Millisecond * 100
}

// A GlottalStop represents a pause, like in "uh-oh" or "Batman".
type GlottalStop struct{}

func (g GlottalStop) EncodeBeginning(system VocalSystem, lastPhone, nextPhone Phone) {
	system.AdjustVolume(0, time.Millisecond*50)
	system.Continue(time.Millisecond * 50)
}

func (g GlottalStop) FormantPull(end FormantState) FormantState {
	end.Volumes = [3]float64{}
	return end
}

func (v GlottalStop) TransitionTime() time.Duration {
	return time.Millisecond * 30
}
