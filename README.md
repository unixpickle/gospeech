# Purpose

I am attempting to use a pronunciation dictionary to synthesize speech.

Currently, the program can synthesize speech using a basic concatenative process. However, I have yet to record a complete set of diphones to fully test it. My guess is that it will sound pretty awful.

# Input formats

The speech synthesizer requires a voice and a pronunciation dictionary. The formats for both are described below.

## Dictionary

The synthesizer can read pronunciation dictionaries in a very specific format. Each line in a dictionary (separated by line feeds) is either a word entry or a comment. Comments begin with three semicolons.

Every line which is not a comment corresponds to a word. The lines follow the format "WORD  PHONE1 PHONE2 ..."&mdash;that is, a word in all caps followed by two spaces followed by a space-delimited list of phones (which are listed below). For the vowel phones (i.e. AA, AE, AH, AO, AW, AY, EH, ER, EY, IH, IY, OW, OY, UH, and UW) a number 0, 1, or 2 may be added to note emphasis.

Here is an example of a dictionary file:

    ;;; Hello reader. This is dictionary version 7.
    ;;; I really like this dictionary. It's nice.
    ABHOR  AE0 B HH AO1 R
    ABHORRED  AH0 B HH AO1 R D
    ABHORRENCE  AH0 B HH AO1 R AH0 N S
    ABHORRENT  AE0 B HH AO1 R AH0 N T
    ;;; Random comments are fun
    ABHORS  AH0 B HH AO1 R Z
    ABIDE  AH0 B AY1 D
    ABIDED  AH0 B AY1 D IH0 D
    ABIDES  AH0 B AY1 D Z

## Voices

A voice is simply a directory containing WAV files. The names of these files determines which phones or diphones they correspond to.

Normal diphones are named using dashes. For example, the diphone between "AA" and "D" would be named "AA-D.wav".

For the beginnings and ends of words, diphones are not sufficient. Edge phones are also necessary. Edge phones occuring at the beginning of words begin with a dash. For example, "-AA.wav" would be an "AA" half-phone at the beginning of a word. The opposite, "AA-.wav", would correspond to an "AA" half-phone at the end of a word.

When a diphone is not available, edge phones can be used to recreate it. For example, if "AA-D.wav" does not exist, the synthesizer can combine "AA-.wav" and "-D.wav". However, this sounds terrible and thus should be employed as infrequently as possible.

For easier editing, voice directories may also contain a "cuts.json" file. Each entry in the cuts.json file specifies a start and end timestamp for a corresponding WAV file. For example, suppose this is cuts.json:

    {"AA-D.wav": {"start": 0.5, "end": 0.9}}

At runtime, the synthesizer will automatically know to use only the part of AA-D.wav from 0.5 seconds to 0.9 seconds. This allows for non-destructive editing and cropping of diphones and edge phones.

## Phones

The phones (individual sounds) that this synthesizer uses are as follows:

| Name | Examples                               |
|------|----------------------------------------|
| AA   | **hon**est, **o**ccupation             |
| AE   | **a**nalyze, **a**ct                   |
| AH   | **a**side, **a**ssist                  |
| AO   | **au**dit, **au**tomatic, **a**lthough |
| AW   | **out**comes, h**ou**r                 |
| AY   | **I**reland, **i**tem                  |
| B    | **b**asic, **b**elong                  |
| CH   | **ch**annel, **ch**oice                |
| D    | **d*ream, **d**esktop                  |
| DH   | **th**ose, **th**em                    |
| EH   | **e**cho, **e**vidence, **e**nhance    |
| ER   | **ear**nings, **ur**ban, **ear**ly     |
| EY   | **eigh*t, **a**ges, **A**sia           |
| F    | **f**ind, **f**acility                 |
| G    | **g**reat, **g**ame                    |
| HH   | **h**ot, **h**azard                    |
| IH   | **i**interview, **i**gnore             |
| IY   | **ea**sy, rep**ea**t                   |
| JH   | **j**oy, **j**aguar                    |
| K    | **c**ooling, **c**ompare, **k**ill     |
| L    | **l**ungs, re**l**ocate                |
| M    | **m**anage, hi**m**                    |
| N    | **n**ame, **N**ovember                 |
| OW   | **ow**n, **o**verview                  |
| OY   | **oi**l                                |
| P    | **p**itch, **p**rimary                 |
| R    | **r**oom, ha**r**m                     |
| S    | **s**ponsor, **c**igarettes            |
| SH   | **sh**e, na**ti**on                    |
| T    | **t**oo, **t**empla**t**e              |
| TH   | **th**irty, **th**ick                  |
| UH   | pl**u**s, sh**oo*k                     |
| UW   | **oo**ps, **U**zbekistan, **oo**zing   |
| V    | **v**ermont, **v**iolin                |
| W    | **w**alnut, **w**orld                  |
| Y    | **y**oga, **y**es                      |
| Z    | **z**one, reali**z**e                  |
| ZH   | **g**enre, presti**g**e                |
