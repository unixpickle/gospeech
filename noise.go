package gospeech

import (
	"math"
	"math/rand"
)

type noiseGenerator struct {
	frequency float64
	spread    float64

	lastT    float64
	argument float64
}

func newNoiseGenerator(freq, spread float64) *noiseGenerator {
	return &noiseGenerator{frequency: freq, spread: spread}
}

func (n *noiseGenerator) nextSample(t float64) float64 {
	elapsed := t - n.lastT
	n.lastT = t

	freq := rand.NormFloat64()*n.spread + n.frequency
	n.argument += math.Pi * 2 * elapsed * freq

	return math.Sin(n.argument)
}
