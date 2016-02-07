package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/unixpickle/gospeech"
)

var AssetsDir string
var Dictionary gospeech.Dictionary

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Usage: server <dictionary.txt> <assets_dir> <port>")
		os.Exit(1)
	}

	var err error
	Dictionary, err = gospeech.LoadDictionary(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	AssetsDir = os.Args[2]

	if port, err := strconv.Atoi(os.Args[3]); err != nil || port < 0 || port > 65535 {
		fmt.Fprintln(os.Stderr, "Invalid port:", os.Args[3])
		os.Exit(1)
	}

	http.HandleFunc("/synthesize_text", SynthesizeText)
	http.HandleFunc("/synthesize_ipa", SynthesizeIPA)
	http.Handle("/", http.FileServer(http.Dir(AssetsDir)))

	http.ListenAndServe(":"+os.Args[3], nil)
}

func SynthesizeText(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	ipa := Dictionary.TranslateToIPA(text)
	ServeSynthesized(w, r, ipa)
}

func SynthesizeIPA(w http.ResponseWriter, r *http.Request) {
	ipa := r.FormValue("ipa")
	ServeSynthesized(w, r, ipa)
}

func ServeSynthesized(w http.ResponseWriter, r *http.Request, ipa string) {
	wav := gospeech.DefaultVoice.Synthesize(ipa)
	var buf bytes.Buffer
	wav.Write(&buf)
	reader := bytes.NewReader(buf.Bytes())
	w.Header().Set("Content-Type", "audio/x-wav")
	http.ServeContent(w, r, "synth.wav", time.Now(), reader)
}
