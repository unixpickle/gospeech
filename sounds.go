package gospeech

import (
	"errors"
	"github.com/unixpickle/wav"
	"os"
	"path/filepath"
	"strings"
)

type Sounds map[string]wav.Sound

func ReadSounds(dirPath string) (Sounds, error) {
	// Read the directory
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	// Load all WAV files
	res := Sounds{}
	for _, name := range names {
		if !strings.HasSuffix(name, ".wav") {
			continue
		}
		baseName := name[0 : len(name)-4]
		sound, err := wav.ReadSoundFile(filepath.Join(dirPath, name))
		if err != nil {
			return nil, err
		}
		res[baseName] = sound
	}

	// Make sure no phonemes are missing
	for _, phoneme := range AllPhonemes() {
		name := phoneme.String()
		if _, ok := res[phoneme.String()]; ok {
			continue
		}
		basePhon := *phoneme
		basePhon.Emphasis = -1
		baseName := basePhon.String()
		if _, ok := res[baseName]; ok {
			res[name] = res[baseName]
			continue
		}
		return nil, errors.New("Missing phoneme: " + name)
	}

	return res, nil
}

func (s Sounds) Get(p *Phoneme) wav.Sound {
	return s[p.String()]
}
