package phrasegen

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func LoadFile(fpath string) string {
	data, err := os.ReadFile(fpath)
	if err != nil {
		log.Printf("Unable to read input file %s: %s\n", fpath, err)
		os.Exit(1)
	}
	return strings.TrimSuffix(string(data), "\n")
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

func GetInput(opts CliOptions) string {
	var inp string
	if opts.Input != "" {
		inp = opts.Input
	} else {
		inp = LoadFile(opts.InputFile)
	}
	if !opts.NoStripPunc {
		inp = clean([]byte(inp))
	}
	return inp
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

func SlidingWindow(words []string, size int) [][]string {
	if size > len(words) || size <= 0 {
		return nil
	}

	result := [][]string{}
	for i := 0; i <= len(words)-size; i++ {
		result = append(result, words[i:i+size])
	}
	return result
}

func ShowPhrases(parts []string, size int, joinStr string) {
	for _, pair := range SlidingWindow(parts, size) {
		fmt.Println(strings.Join(pair, joinStr))
	}
}
