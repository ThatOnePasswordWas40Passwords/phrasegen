package main

import (
	"bufio"
	"log"
	"os"
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

	inp, err := phrasegen.GetInput(cliOpts)
	if err != nil {
		return err
	}

	splitParts := phrasegen.SplitOnSpace(inp)

	var outBuffer *bufio.Writer

	if cliOpts.Outfile != "" {
		outFile, err := os.Create(cliOpts.Outfile)
		if err != nil {
			log.Printf("Unable to create %s for writing? %s", cliOpts.Outfile, err)
			return err
		}
		outBuffer = bufio.NewWriter(outFile)
		defer outFile.Close()
	} else {
		outBuffer = bufio.NewWriter(os.Stdout)
	}

	defer outBuffer.Flush()

	return phrasegen.ShowPhrases(splitParts, cliOpts.Size, cliOpts.Only, cliOpts.JoinStr, outBuffer, cliOpts.Case)
}
