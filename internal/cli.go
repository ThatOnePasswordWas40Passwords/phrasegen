package phrasegen

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type CliOptions struct {
	Verbose     bool   `flag:"verbose"`
	Outfile     string `flag:"o"`
	NoStripPunc bool   `flag:"no-strip"`
	JoinStr     string `flag:"join-str"`
	Input       string `flag:"i"`
	Size        int    `flag:"size"`
	Only        bool   `flag:"only"`
}

func (opts CliOptions) String() string {
	return fmt.Sprintf(
		"CliOptions(verbose=%t, outfile=%s, join-str=%s, no-strip=%t, input=%s, size=%d)",
		opts.Verbose,
		opts.Outfile,
		opts.JoinStr,
		opts.NoStripPunc,
		opts.Input,
		opts.Size,
	)
}

func ParseArgs() (CliOptions, error) {
	var cliOpts CliOptions

	flag.StringVar(
		&cliOpts.Outfile,
		"o",
		"",
		"Filepath to write output to. If not given, will write to STDOUT (default empty)",
	)
	flag.StringVar(
		&cliOpts.Input,
		"i",
		"",
		"Path to local file to generate phrases from OR, if not a valid filepath, a raw string to operate against",
	)
	flag.IntVar(&cliOpts.Size, "size", 3, "Number of words to include in generated phrases (default 3)")
	flag.StringVar(&cliOpts.JoinStr, "join-str", "", "Char/String to join words with (default \"\")")
	flag.BoolVar(
		&cliOpts.NoStripPunc,
		"no-strip",
		false,
		"Don't strip punctuation from source while joining (default false)",
	)
	flag.BoolVar(
		&cliOpts.Only,
		"only",
		false,
		"ONLY show phrases of specified -size (default false; i.e, phrases of size [1, size] will be generated)",
	)
	flag.BoolVar(
		&cliOpts.Verbose,
		"verbose",
		false,
		"Print additional debug/verbose message during execution (default false)",
	)
	flag.Parse()

	if cliOpts.Verbose {
		log.Println(cliOpts)
	}

	if cliOpts.Input == "" {
		return CliOptions{}, errors.New("must specify either a local file path or a raw string via -i '...'")
	}

	return cliOpts, nil
}
