package wordlist

import (
	"io"
	"os"
	"sort"
	"strings"
)

type Mapping map[string]string

func ReadMapping(path string) (Mapping, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	res := Mapping{}
	for _, line := range lines {
		idx := strings.Index(line, "|")
		if idx < 0 {
			continue
		}
		res[line[0:idx]] = line[idx+1:]
	}
	return res, nil
}

func (m Mapping) Write(w io.Writer) error {
	keys := make([]string, 0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		_, err := w.Write([]byte(key + "|" + m[key] + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Mapping) WriteFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return m.Write(f)
}
