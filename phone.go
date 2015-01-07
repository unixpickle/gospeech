package gospeech

import (
	"errors"
	"strings"
)

// These are the types of oral sounds that an English speaker can produce.
const (
	TypeAffricate = iota
	TypeAspirate  = iota
	TypeFricative = iota
	TypeLiquid    = iota
	TypeNasal     = iota
	TypeSemivowel = iota
	TypeStop      = iota
	TypeVowel     = iota
)

// PhoneTypes associates each phone with its numeric type.
var PhoneTypes = map[Phone]int{
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
	"M": TypeNasal, "N": TypeNasal, "NG": TypeNasal,
}

// Phone is a phonetic sound.
type Phone string

// Name returns a string cast of the phone.
func (p Phone) Name() string {
	return string(p)
}

// String returns a string cast of the phone.
func (p Phone) String() string {
	return string(p)
}

// Type returns the type of sound produced by this phone.
func (p Phone) Type() int {
	return PhoneTypes[p]
}

// Valid returns whether or not the phone is valid.
// Phones may be invalid if they were cast directly from strings.
func (p Phone) Valid() bool {
	_, ok := PhoneTypes[p]
	return ok
}

// AllPhones returns an array of every basic phone in the English language.
func AllPhones() []Phone {
	str := "B D G K P T AA AE AH AO AW AY EH ER EY IH IY OW OY UH UW CH JH " +
		"DH F S SH TH V Z ZH HH L R W Y M N NG"
	comps := strings.Split(str, " ")
	res := make([]Phone, len(comps))
	for i, s := range comps {
		res[i] = Phone(s)
	}
	return res
}

// ParsePhone parses a phone from its string representation.
// This method automatically removes "0", "1", or "2" from the end of the string
// before parsing it.
func ParsePhone(ph string) (Phone, error) {
	if len(ph) == 0 {
		return "", errors.New("Empty phone is invalid.")
	}
	if ph[len(ph)-1] >= '0' && ph[len(ph)-1] <= '2' {
		ph = ph[0 : len(ph)-1]
	}
	res := Phone(ph)
	if !res.Valid() {
		return "", errors.New("Invalid phone: " + ph)
	}
	return res, nil
}

// ParsePhones parses a space-delimited list of phones.
func ParsePhones(raw string) ([]Phone, error) {
	comps := strings.Split(raw, " ")
	res := make([]Phone, len(comps))
	for i, comp := range comps {
		var err error
		res[i], err = ParsePhone(comp)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
