package gospeech

import "github.com/unixpickle/wav"

func SynthesizePhones(phones []Phone, v Voice) wav.Sound {
	res := wav.NewPCM8Sound(1, 44100)
	if len(phones) == 0 {
		return res
	}

	wav.Append(res, v["-"+phones[0].Name()])
	for i := 0; i < len(phones)-1; i++ {
		wav.Append(res, getDiphone(phones[i], phones[i+1], v))
	}
	wav.Append(res, v[phones[len(phones)-1].Name()+"-"])
	return res
}

func getDiphone(p1, p2 Phone, v Voice) wav.Sound {
	// If the voice has the diphone, use that.
	name := p1.Name() + "-" + p2.Name()
	if dp, ok := v[name]; ok {
		return dp
	}

	// Generate the diphone manually using p1- + -p2
	edge1 := v[p1.Name()+"-"]
	edge2 := v["-"+p2.Name()]
	res := edge1.Clone()
	wav.Append(res, edge2)
	return res
}
