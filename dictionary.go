package gospeech

import (
	"bufio"
	"os"
	"strings"
)

type Dictionary map[string]string

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
		rawLine, _, err := b.ReadLine()
		if err != nil {
			break
		}
		line := string(rawLine)
		if strings.HasPrefix(line, ";;;") {
			continue
		}
		idx := strings.Index(line, "  ")
		if idx < 0 {
			continue
		}
		word := line[0:idx]
		phonetics := line[idx+2:]
		for strings.HasPrefix(phonetics, " ") {
			phonetics = phonetics[1:]
		}
		res[word] = phonetics
	}
	return res, nil
}

func (d Dictionary) Get(word string) []*Phone {
	str, ok := d[strings.ToUpper(word)]
	if !ok {
		return nil
	}
	return parsePhones(str)
}

func parsePhones(raw string) []*Phone {
	parts := strings.Split(raw, " ")
	res := make([]*Phone, len(parts))
	for i, part := range parts {
		var err error
		res[i], err = ParsePhone(part)
		if err != nil {
			return nil
		}
	}
	return res
}
