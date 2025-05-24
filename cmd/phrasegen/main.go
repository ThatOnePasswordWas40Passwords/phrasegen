package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	phrasegen "t1pw40p/tools/phrasegen/internal"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	cliOpts := phrasegen.ParseArgs()
	inp := phrasegen.GetInput(cliOpts)
	splitParts := phrasegen.SplitOnSpace(inp)

	for _, pair := range phrasegen.SlidingWindow(splitParts, cliOpts.Size) {
		fmt.Println(strings.Join(pair, cliOpts.JoinStr))
	}
}
