package phrasegen

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func LoadFile(fpath string) (string, error) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		log.Printf("Unable to read input file %s: %s\n", fpath, err)
		return "", err
	}
	return strings.TrimSuffix(string(data), "\n"), nil
}

func clean(s []byte) string {
	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' || b == '\n' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}

func GetInput(opts CliOptions) (string, error) {
	var inp string
	var err error
	if opts.Input != "" {
		inp = opts.Input
	} else {
		inp, err = LoadFile(opts.InputFile)
		if err != nil {
			return "", err
		}
	}
	if !opts.NoStripPunc {
		inp = clean([]byte(inp))
	}
	return inp, nil
}

// SplitOnNonLetters splits a string on non-letter runes.
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

func SplitOn(s string, splitOn string) []string {
	return strings.Split(s, splitOn)
}

func SplitOnSpace(s string) []string {
	return strings.Fields(s)
}

func SlidingWindow(words []string, size int, only bool) [][]string {
	if size > len(words) || size <= 0 {
		return nil
	}

	result := [][]string{}
	for i := 0; i <= len(words)-size; i++ {
		if only {
			result = append(result, words[i:i+size])
		} else {
			for j := 1; j <= size; j++ {
				result = append(result, words[i:i+j])
			}
		}
	}
	return result
}

func maybeAddPhrase(seen map[string]struct{}, phrase string, buf *bufio.Writer) {
	if _, exists := seen[phrase]; !exists {
		seen[phrase] = struct{}{}
		_, _ = buf.WriteString(phrase + "\n")
	}
}

func ShowPhrases(words []string, size int, only bool, joinStr string, buf *bufio.Writer) error {
	if size > len(words) || size <= 0 {
		return nil
	}

	wordsLen := len(words)
	if only {
		wordsLen -= size
	}

	seen := make(map[string]struct{})
	for i := 0; i <= wordsLen; i++ {
		if only {
			maybeAddPhrase(seen, strings.Join(words[i:i+size], joinStr), buf)
		} else {
			for j := 1; j <= size; j++ {
				if i+j < wordsLen {
					maybeAddPhrase(seen, strings.Join(words[i:i+j], joinStr), buf)
				} else {
					maybeAddPhrase(seen, strings.Join(words[i:], joinStr), buf)
				}
			}
		}
	}
	return nil
}
