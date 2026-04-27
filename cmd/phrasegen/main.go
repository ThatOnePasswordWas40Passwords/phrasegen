package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	phrasegen "t1pw40p/tools/phrasegen/internal"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cliOpts, err := phrasegen.ParseArgs()
	if err != nil {
		return err
	}

	var reader io.Reader
	if _, statErr := os.Stat(cliOpts.Input); errors.Is(statErr, os.ErrNotExist) {
		reader = strings.NewReader(cliOpts.Input)
	} else {
		f, err := os.Open(cliOpts.Input)
		if err != nil {
			log.Printf("Unable to read input file %s: %s\n", cliOpts.Input, err)
			return err
		}
		defer f.Close()
		reader = f
	}

	var outBuffer *bufio.Writer

	if cliOpts.Outfile != "" {
		outFile, err := os.Create(cliOpts.Outfile)
		if err != nil {
			log.Printf("Unable to create %s for writing? %s", cliOpts.Outfile, err)
			return err
		}
		outBuffer = bufio.NewWriterSize(outFile, 256<<10)
		defer outFile.Close()
	} else {
		outBuffer = bufio.NewWriterSize(os.Stdout, 256<<10)
	}

	defer outBuffer.Flush()

	if cliOpts.Dedup {
		return phrasegen.StreamPhrases(
			reader,
			cliOpts.Size,
			cliOpts.Only,
			cliOpts.JoinStr,
			outBuffer,
			cliOpts.Case,
			!cliOpts.NoStripPunc,
		)
	}

	return phrasegen.StreamPhrasesNoDedup(
		reader,
		cliOpts.Size,
		cliOpts.Only,
		cliOpts.JoinStr,
		outBuffer,
		cliOpts.Case,
		!cliOpts.NoStripPunc,
	)
}
