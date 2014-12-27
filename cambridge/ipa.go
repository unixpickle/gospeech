package cambridge

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

func FindIPA(word string) (string, error) {
	e := "<span class=\"ipa\">.*?</span>"
	url := "http://dictionary.cambridge.org/us/dictionary/american-english/" +
		word
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	match := regexp.MustCompile(e).FindString(string(data))
	if match == "" {
		return "", errors.New("Not found.")
	}
	return regexp.MustCompile("<.*?>").ReplaceAllString(match, ""), nil
}
