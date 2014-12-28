package cambridge

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func FindFullIPA(word string) (string, error) {
	ipa, realWord, err := FindIPA(word)
	if err != nil {
		return "", err
	}
	if realWord == word {
		return ipa, nil
	}
	suffixes := map[string]string{"ing": "ɪŋ", "s": "z", "er": "ɜr",
		"ers": "ɜrs", "ly": "lē", "ed": "t", "es": "əz"}
	bestSuffix := ""
	bestEnding := ""
	for suffix, ending := range suffixes {
		if !strings.HasSuffix(word, suffix) {
			continue
		} else if len(suffix) < len(bestSuffix) {
			continue
		}
		bestSuffix = suffix
		bestEnding = ending
	}
	return SmartJoin(ipa, bestEnding), nil
}

func FindIPA(word string) (string, string, error) {
	// Read the online URL
	url := "http://dictionary.cambridge.org/us/search/american-english/" +
		"direct/?q=" + word
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	body := string(data)

	// Read the tags
	ipa, err := findTag(body, "<span class=\"ipa\">", "</span>")
	if err != nil {
		return "", "", err
	}
	heading, err := findTag(body,
		"<h2 class=\"di-title cdo-section-title-hw\">", "</h2>")
	if err != nil {
		return "", "", err
	}
	return ipa, heading, nil
}

func IsVowel(phoneme string) bool {
	dipthongs := []string{"juː", "ɑː", "ɒ", "æ", "aɪ", "aʊ", "ɛ", "eɪ", "ɪ", "iː", "ɔː", "ɔɪ",
		"oʊ", "ʊ", "uː", "ʊ", "ə", "ɨ", "ɵ", "ʉ", "i"}
	for _, d := range dipthongs {
		if d == phoneme {
			return true
		}
	}
	return false
}

func SmartJoin(ipa1, ipa2 string) string {
	ph1 := SplitPhonemes(ipa1)
	ph2 := SplitPhonemes(ipa2)
	if len(ph1) == 0 {
		return ipa1
	} else if len(ph2) == 0 {
		return ipa2
	}
	if !IsVowel(ph1[len(ph1)-1]) && !IsVowel(ph2[0]) {
		return ipa1 + "ə" + ipa2
	}
	return ipa1 + ipa2
}

func SplitPhonemes(ipa string) []string {
	// It is important that the three-letter dipthong comes first.
	dipthongs := []string{"juː", "ɑː", "aɪ", "aʊ", "eɪ", "iː", "ɔː", "ɔɪ",
		"oʊ", "uː"}
	for _, dt := range dipthongs {
		idx := strings.Index(ipa, dt)
		if idx >= 0 {
			pre := SplitPhonemes(ipa[0:idx])
			post := SplitPhonemes(ipa[idx+len(dt):])
			res := make([]string, 0, len(pre)+len(post)+1)
			res = append(res, pre...)
			res = append(res, dt)
			res = append(res, post...)
			return res
		}
	}
	runes := []rune(ipa)
	res := make([]string, len(runes))
	for i, r := range runes {
		res[i] = string(r)
	}
	return res
}

func findTag(body, open, close string) (string, error) {
	e := open + "(.*?)" + close
	r, err := regexp.Compile(e)
	if err != nil {
		return "", err
	}
	m := r.FindStringSubmatch(body)
	if m == nil {
		return "", errors.New("Not found.")
	}
	return m[1], nil
}
