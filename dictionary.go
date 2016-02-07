package gospeech

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// A Dictionary maps lowercase words to their IPA representations.
type Dictionary map[string]string

// LoadDictionary reads a dictionary file.
// The file must be CSV with two columns: the word and the word's IPA representation.
func LoadDictionary(path string) (Dictionary, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(contents), "\n")
	res := Dictionary{}
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		comps := strings.Split(line, ",")
		if len(comps) != 2 {
			return nil, errors.New("unexpected string at line: " + strconv.Itoa(i))
		}
		if _, ok := res[comps[0]]; ok {
			return nil, errors.New("repeated entry: " + comps[0])
		}
		res[comps[0]] = comps[1]
	}
	return res, nil
}

// TranslateToIPA uses the dictionary to convert the words in a block of text into IPA.
// This will ignore punctuation, capitalization, etc.
func (d Dictionary) TranslateToIPA(text string) string {
	text = strings.ToLower(text)
	text = strings.Replace(text, "'", "", -1)
	text = strings.Replace(text, ".", " ", -1)
	text = strings.Replace(text, "?", " ", -1)
	text = strings.Replace(text, ";", " ", -1)
	text = strings.Replace(text, "-", " ", -1)
	text = strings.Replace(text, ",", " ", -1)
	text = strings.Replace(text, "\n", " ", -1)

	res := []string{}
	for _, word := range strings.Fields(text) {
		if ipa, ok := d[word]; ok {
			res = append(res, ipa)
		}
	}

	return strings.Join(res, " ")
}
