<div align="center">

# phrasegen

![Contributions welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Generating sliding window ngrams of specified customizations over a given body of text.
Great for generating passphrase candidates with a focus on human memorization from a
given source material.

<br/>

[Why](#why) • [Getting started](#installation) •
[Usage](#usage) • [FAQs, feedback, etc](#faqs)

</div>

## Why

Q: Many other tools for building wordlists from input material already exist, why bother
with this?

A: Sometimes less is more. When you _only_ want generated passphrases that maintain
original apperance ordering, this is the tool you want. A massive keyspace generated
by one of those other tools will contain these values, of course, but they also will
have many, many, many other non-interested in outputs.

## Installation

TODO

## Usage

```bash
phrasegen -h
```

e.g Generating passphrases from a movie script, of exactly 4 phrases each, joined by a
`-` character, without having scrubbed any of the punctuation from the source material
before after joining, forcing all to lowercase, and storing in an output file `gen.txt`:

```bash
phrasegen --join-str "-" \
    --file movie_script.txt \
    --only-size 4 \
    --no-strip-punc \
    --lowercase \
    -o gen.txt
```

Generating from a given string via CLI directly and/or outputting to STDOUT is also possible:

```bash
phrasegen --join-str "-" \
    --input "This is a sentence I'd like to use to generate a set of passphrases against." \
    --max-size 3 \
    --no-strip-punc \
    --lowercase
```

Outputs to STDOUT:

```text
this
is
a
sentence
i'd
like
to
use
to
generate
a
set
of
passphrases
against.
this-is
is-a
a-sentence
sentence-i'd
i'd-like
like-to
to-use
use-to
to-generate
generate-a
a-set
set-of
of-passphrases
passphrases-against.
this-is-a
is-a-sentence
a-sentence-i'd
sentence-i'd-like
i'd-like-to
like-to-use
to-use-to
use-to-generate
to-generate-a
generate-a-set
a-set-of
set-of-passphrases
of-passphrases-against.
```

## FAQs

Please submit an issue against this repo.

## Developement

See [DEVELOPMENT.md]
