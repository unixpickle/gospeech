package gospeech

import (
	"strconv"
	"strings"
)

const (
	TypeAffricate = iota
	TypeAspirate  = iota
	TypeFricative = iota
	TypeLiquid    = iota
	TypeSemivowel = iota
	TypeStop      = iota
	TypeVowel     = iota
)

type Phoneme struct {
	Emphasis int
	Letters  string
	Type     int
}

func AllPhonemes() []*Phoneme {
	lists := map[int]string{
		TypeStop:      "B D G K P T",
		TypeVowel:     "AA AE AH AO AW AY EH ER EY IH IY OW OY UH UW",
		TypeAffricate: "CH JH",
		TypeFricative: "DH F S SH TH V Z ZH",
		TypeAspirate:  "HH",
		TypeLiquid:    "L R",
		TypeSemivowel: "W Y",
	}
	res := make([]*Phoneme, 0)
	for t, v := range lists {
		list := strings.Split(v, " ")
		for _, s := range list {
			res = append(res, &Phoneme{-1, s, t})
			if t == TypeVowel {
				// Append phonemes with the 3 emphases
				for i := 0; i < 3; i++ {
					res = append(res, &Phoneme{i, s, t})
				}
			}
		}
	}
	return res
}

func ParsePhoneme(ph string) (*Phoneme, error) {
	types := map[string]int{
		"R": TypeLiquid, "W": TypeSemivowel, "Y": TypeSemivowel,
		"B": TypeStop, "D": TypeStop, "G": TypeStop,
		"K": TypeStop, "P": TypeStop, "T": TypeStop,
		"AA": TypeVowel, "AE": TypeVowel, "AH": TypeVowel,
		"AO": TypeVowel, "AW": TypeVowel, "AY": TypeVowel,
		"EH": TypeVowel, "ER": TypeVowel, "EY": TypeVowel,
		"IH": TypeVowel, "IY": TypeVowel, "OW": TypeVowel,
		"OY": TypeVowel, "UH": TypeVowel, "UW": TypeVowel,
		"CH": TypeAffricate, "JH": TypeAffricate, "DH": TypeFricative,
		"F": TypeFricative, "S": TypeFricative, "SH": TypeFricative,
		"TH": TypeFricative, "V": TypeFricative, "Z": TypeFricative,
		"ZH": TypeFricative, "HH": TypeAspirate, "L": TypeLiquid,
	}
	emphasis := -1
	if strings.HasSuffix(ph, "0") || strings.HasSuffix(ph, "1") ||
		strings.HasSuffix(ph, "2") {
		emphasis = int(ph[len(ph)-1] - '0')
		ph = ph[0 : len(ph)-1]
	}
	theType, ok := types[ph]
	if !ok {
		return nil, ErrInvalidPhoneme
	}
	return &Phoneme{emphasis, ph, theType}, nil
}

func (p *Phoneme) String() string {
	if p.Emphasis < 0 {
		return p.Letters
	}
	return p.Letters + strconv.Itoa(p.Emphasis)
}
