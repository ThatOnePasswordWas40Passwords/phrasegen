package phrasegen

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Casing string

const (
	UPPER Casing = "upper"
	LOWER Casing = "lower"
)

type CliOptions struct {
	Verbose     bool   `flag:"verbose"`
	Outfile     string `flag:"o"`
	NoStripPunc bool   `flag:"no-strip"`
	JoinStr     string `flag:"join-str"`
	Input       string `flag:"i"`
	Size        int    `flag:"size"`
	Only        bool   `flag:"only"`
	Case        Casing `flag:"casing"`
}

type casingValue struct {
	casing Casing
}

func (casing *casingValue) String() string {
	return string(casing.casing)
}

func (casing *casingValue) Set(s string) error {
	switch s {
	case string(UPPER), string(LOWER), "":
		casing.casing = Casing(s)
		return nil
	default:
		return errors.New("-casing, if provided, must be of: 'upper|lower'")
	}
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
	var casing casingValue

	flag.StringVar(
		&cliOpts.Outfile,
		"o",
		"",
		"Filepath to write output to. If not given, will write to STDOUT (default empty)",
	)
	flag.Var(&casing, "casing", "Casing to force; If not provided, no case mangling will occur (default '')")
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

	cliOpts.Case = casing.casing

	return cliOpts, nil
}
