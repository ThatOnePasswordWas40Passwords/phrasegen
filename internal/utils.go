package phrasegen

import (
	"bufio"
	"errors"
	"io"
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

func Clean(s []byte) string {
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

	if _, err = os.Stat(opts.Input); errors.Is(err, os.ErrNotExist) {
		inp = opts.Input
	} else {
		inp, err = LoadFile(opts.Input)
		if err != nil {
			return "", err
		}
	}

	if !opts.NoStripPunc {
		inp = Clean([]byte(inp))
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

func ForceCasing(word string, casing Casing) string {
	switch casing {
	case UPPER:
		return strings.ToUpper(word)
	case LOWER:
		return strings.ToLower(word)
	default:
		return word
	}
}

func ShowPhrases(words []string, size int, only bool, joinStr string, buf *bufio.Writer, casing Casing) error {
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
			maybeAddPhrase(seen, ForceCasing(strings.Join(words[i:i+size], joinStr), casing), buf)
		} else {
			for j := 1; j <= size; j++ {
				if i+j < wordsLen {
					maybeAddPhrase(seen, ForceCasing(strings.Join(words[i:i+j], joinStr), casing), buf)
				} else {
					maybeAddPhrase(seen, ForceCasing(strings.Join(words[i:], joinStr), casing), buf)
				}
			}
		}
	}
	return nil
}

func cleanWord(word string) string {
	n := 0
	for i := 0; i < len(word); i++ {
		b := word[i]
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
			n++
		}
	}
	if n == len(word) {
		return word
	}
	if n == 0 {
		return ""
	}
	buf := make([]byte, 0, n)
	for i := 0; i < len(word); i++ {
		b := word[i]
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
			buf = append(buf, b)
		}
	}
	return string(buf)
}

func buildPhrase(words []string, joinStr string, b *strings.Builder) string {
	b.Reset()
	for i, w := range words {
		if i > 0 {
			b.WriteString(joinStr)
		}
		b.WriteString(w)
	}
	return b.String()
}

func writeIfNew(seen map[string]struct{}, phrase string, buf *bufio.Writer) {
	if _, exists := seen[phrase]; !exists {
		seen[phrase] = struct{}{}
		_, _ = buf.WriteString(phrase)
		_ = buf.WriteByte('\n')
	}
}

func writePhrase(words []string, joinStr string, buf *bufio.Writer) {
	for i, w := range words {
		if i > 0 {
			_, _ = buf.WriteString(joinStr)
		}
		_, _ = buf.WriteString(w)
	}
	_ = buf.WriteByte('\n')
}

func emitPhrases(words []string, size int, only bool, joinStr string, builder *strings.Builder, seen map[string]struct{}, buf *bufio.Writer) {
	if only {
		if len(words) >= size {
			writeIfNew(seen, buildPhrase(words[:size], joinStr, builder), buf)
		}
	} else {
		maxJ := size
		if maxJ > len(words) {
			maxJ = len(words)
		}
		for j := 1; j <= maxJ; j++ {
			writeIfNew(seen, buildPhrase(words[:j], joinStr, builder), buf)
		}
	}
}

func emitPhrasesNoDedup(words []string, size int, only bool, joinStr string, buf *bufio.Writer) {
	if only {
		if len(words) >= size {
			writePhrase(words[:size], joinStr, buf)
		}
	} else {
		maxJ := size
		if maxJ > len(words) {
			maxJ = len(words)
		}
		for j := 1; j <= maxJ; j++ {
			writePhrase(words[:j], joinStr, buf)
		}
	}
}

func StreamPhrases(reader io.Reader, size int, only bool, joinStr string, buf *bufio.Writer, casing Casing, stripPunc bool) error {
	return streamPhrases(reader, size, only, joinStr, buf, casing, stripPunc, false)
}

func StreamPhrasesNoDedup(reader io.Reader, size int, only bool, joinStr string, buf *bufio.Writer, casing Casing, stripPunc bool) error {
	return streamPhrases(reader, size, only, joinStr, buf, casing, stripPunc, true)
}

func streamPhrases(reader io.Reader, size int, only bool, joinStr string, buf *bufio.Writer, casing Casing, stripPunc bool, noDedup bool) error {
	if size <= 0 {
		return nil
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	var seen map[string]struct{}
	var builder strings.Builder
	if !noDedup {
		seen = make(map[string]struct{})
	}
	words := make([]string, 0, size+1)
	count := 0

	for scanner.Scan() {
		word := scanner.Text()
		if stripPunc {
			word = cleanWord(word)
			if word == "" {
				continue
			}
		}
		switch casing {
		case UPPER:
			word = strings.ToUpper(word)
		case LOWER:
			word = strings.ToLower(word)
		}
		count++
		words = append(words, word)

		if len(words) > size {
			if noDedup {
				emitPhrasesNoDedup(words, size, only, joinStr, buf)
			} else {
				emitPhrases(words, size, only, joinStr, &builder, seen, buf)
			}
			copy(words, words[1:])
			words[size] = ""
			words = words[:size]
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if count < size {
		return nil
	}

	if only {
		if len(words) == size {
			if noDedup {
				emitPhrasesNoDedup(words, size, only, joinStr, buf)
			} else {
				emitPhrases(words, size, only, joinStr, &builder, seen, buf)
			}
		}
	} else {
		for len(words) > 0 {
			if noDedup {
				emitPhrasesNoDedup(words, size, only, joinStr, buf)
			} else {
				emitPhrases(words, size, only, joinStr, &builder, seen, buf)
			}
			copy(words, words[1:])
			words[len(words)-1] = ""
			words = words[:len(words)-1]
		}
	}

	return nil
}
