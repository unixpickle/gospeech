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
	if !strings.HasPrefix(word, realWord) {
		return ipa, nil
	}
	suffix := word[len(realWord):]
	suffixes := map[string]string{"s": "s"}
	if ending, ok := suffixes[suffix]; ok {
		return ipa + ending, nil
	}
	return ipa, nil
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

func findTag(body, open, close string) (string, error) {
	e := open + "(.*?)" + close
	r, err := regexp.Compile(e)
	if err != nil {
		return "", err
	}
	m := r.FindStringSubmatch(body)
	if m == nil {
		return "", errors.New("Not found: " + e)
	}
	return m[1], nil
}
