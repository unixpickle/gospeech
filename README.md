# Purpose

I am attempting to use a pronunciation dictionary to synthesize speech.

# Status

Currently, the program can read a pronunciation dictionary and a "voice" (a collection of phonetic WAV files) and produce an audio sentence. The raw performance is somewhat poor as it takes over a second to synthesize 6 syllables. The quality of audio produced is also quite bad: usually it's intelligible, but it's choppy and bad.

Currently, I am using [CMU's pronunciation dictionary](http://svn.code.sf.net/p/cmusphinx/code/trunk/cmudict/cmudict-0.7b).

You can also download [a voice](http://aqnichol.com/downloads/voice.zip) from my website. Right now, it is just a very rough draft and still sounds terrible.
