package gospeech

import (
	"bufio"
	"os"
	"strings"
)

type Dictionary map[string][]Phone

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
		if parsed := parsePhones(rawPhones); parsed != nil {
			res[word] = parsed
		}
	}
	return res, nil
}

func (d Dictionary) Get(word string) []Phone {
	if res, ok := d[strings.ToUpper(word)]; ok {
		return res
	} else {
		return nil
	}
}

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

func parsePhones(raw string) []Phone {
	comps := strings.Split(raw, " ")
	res := make([]Phone, len(comps))
	for i, comp := range comps {
		res[i] = Phone(comp)
		if !res[i].Valid() {
			return nil
		}
	}
	return res
}
