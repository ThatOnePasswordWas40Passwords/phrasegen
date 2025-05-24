package main

import (
	"log"
	"os"
	phrasegen "t1pw40p/tools/phrasegen/internal"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	cliOpts := phrasegen.ParseArgs()
	inp := phrasegen.GetInput(cliOpts)
	splitParts := phrasegen.SplitOnSpace(inp)

	if cliOpts.Only {
		phrasegen.ShowPhrases(splitParts, cliOpts.Size, cliOpts.JoinStr)
		return
	}

	for size := range cliOpts.Size + 1 {
		phrasegen.ShowPhrases(splitParts, size, cliOpts.JoinStr)
	}
}
