package gospeech

import (
	"bufio"
	"os"
	"strings"
)

// Dictionary is a phonetic dictionary which maps words to corresponding
// pronunciations.
type Dictionary map[string][]Phone

// LoadDictionary reads a dictionary file and returns it.
func LoadDictionary(path string) (Dictionary, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b := bufio.NewReader(f)

	// Read each line and add it to the dictionary
	res := Dictionary{}
	for {
		// Read the line
		rawLine, _, err := b.ReadLine()
		if err != nil {
			break
		}
		line := string(rawLine)

		// Ignore comments or lines without proper separators
		if strings.HasPrefix(line, ";;;") {
			continue
		}
		idx := strings.Index(line, "  ")
		if idx < 0 {
			continue
		}

		// Parse the entry
		word := line[0:idx]
		rawPhones := strings.TrimLeft(line[idx+2:], " ")
		if parsed, err := ParsePhones(rawPhones); err == nil {
			res[word] = parsed
		}
	}
	return res, nil
}

// Get returns the phonetic pronunciation (or nil) for a given word.
// The case of the word does not matter.
func (d Dictionary) Get(word string) []Phone {
	if res, ok := d[strings.ToUpper(word)]; ok {
		return res
	} else {
		return nil
	}
}

// GetRaw returns the phonetic string representation of a word.
// This is equivalent to calling Get() and joining the resultant phones with
// spaces.
func (d Dictionary) GetRaw(word string) string {
	res := ""
	phones := d.Get(word)
	if phones == nil {
		return ""
	}
	for i, ph := range phones {
		if i != 0 {
			res += " "
		}
		res += ph.String()
	}
	return res
}
