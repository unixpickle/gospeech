# Generators

These two scripts will help you generate words for recording diphones. For example, if you need a record the "AA-B" sound, the scripts might suggest the word "<b>ob</b>servation".

Both scripts take two inputs: a pronunciation dictionary and a common words list. The scripts use the pronunciation dictionary to find words which use a given sound. They use the common words list to sort their suggestions.

# Pre-generated lists.

In this directory, I have included a generated list of [diphones](diphones.txt) and [edge phones](edge_phones.txt). These should suffice for most purposes.

# Where to get files

I use CMU's pronuncation dictionary. Currently, [this link](http://svn.code.sf.net/p/cmusphinx/code/trunk/cmudict/cmudict-0.7b) takes you straight to the right file, but I fear that it may break in the future.

For common words, I use [this wordlist](https://github.com/first20hours/google-10000-english).
