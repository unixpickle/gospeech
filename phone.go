package gospeech

import (
	"errors"
	"strings"
)

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

type Phone string

func (p Phone) Name() string {
	return string(p)
}

func (p Phone) String() string {
	return string(p)
}

func (p Phone) Type() int {
	return PhoneTypes[p]
}

func (p Phone) Valid() bool {
	_, ok := PhoneTypes[p]
	return ok
}

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

func ParsePhone(ph string) (Phone, error) {
	res := Phone(ph)
	if !res.Valid() {
		return "", errors.New("Invalid phone: " + ph)
	}
	return res, nil
}
