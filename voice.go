package gospeech

import (
	"encoding/json"
	"github.com/unixpickle/wav"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Voice map[string]wav.Sound

func LoadVoice(path string) (Voice, error) {
	res, err := loadRawVoice(path)
	if err != nil {
		return nil, err
	}

	// Load the cuts.json file if possible
	f, err := os.Open(filepath.Join(path, "cuts.json"))
	if err != nil {
		return res, nil
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return res, nil
	}

	// Parse cuts.json
	var cuts map[string]fileCut
	if err := json.Unmarshal(content, cuts); err != nil {
		return res, nil
	}
	for name, cut := range cuts {
		if sound, ok := res[name]; ok {
			start := time.Duration(float64(time.Second) * cut.start)
			end := time.Duration(float64(time.Second) * cut.end)
			wav.Crop(sound, start, end)
		}
	}

	return res, nil
}

func loadRawVoice(path string) (Voice, error) {
	// Read the directory
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	// Load all WAV files
	res := Voice{}
	for _, name := range names {
		if !strings.HasSuffix(name, ".wav") {
			continue
		}
		baseName := name[0 : len(name)-4]
		sound, err := wav.ReadSoundFile(filepath.Join(path, name))
		if err != nil {
			return nil, err
		}
		res[baseName] = sound
	}

	return res, nil
}

type fileCut struct {
	start float64
	end   float64
}
