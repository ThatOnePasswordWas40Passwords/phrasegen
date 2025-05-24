package phrasegen

import (
	"flag"
	"fmt"
	"log"
)

type CliOptions struct {
	Verbose     bool   `flag:"verbose"`
	Outfile     string `flag:"outfile"`
	NoStripPunc bool   `flag:"no-strip"`
	JoinStr     string `flag:"join-str"`
	InputFile   string `flag:"infile"`
	Input       string `flag:"input"`
	Size        int    `flag:"size"`
	Only        bool   `flag:"only"`
}

func (opts CliOptions) String() string {
	return fmt.Sprintf(
		"CliOptions(verbose=%t, outfile=%s, join-str=%s, no-strip=%t, infile=%s, input=%s, size=%d)",
		opts.Verbose,
		opts.Outfile,
		opts.JoinStr,
		opts.NoStripPunc,
		opts.InputFile,
		opts.Input,
		opts.Size,
	)
}

func ParseArgs() CliOptions {
	var cliOpts CliOptions

	flag.StringVar(
		&cliOpts.Outfile,
		"outfile",
		"",
		"Filepath to write output to. If not given, will write to STDOUT (default empty)",
	)
	flag.StringVar(
		&cliOpts.InputFile,
		"infile",
		"",
		"Path to local file to generate phrases from; MUST specify exactly one of this or -input  (default \"\")",
	)
	flag.StringVar(
		&cliOpts.Input,
		"input",
		"",
		"Raw string generate phrases from; MUST specify exactly one of this or -infile (default \"\")",
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
		"ONLY show phrases of specified -size (default false; i.e, phrases of size 1->-size will be generated)",
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

	if cliOpts.Input == "" && cliOpts.InputFile == "" {
		log.Fatal(
			"Must specify exactly one of either '-infile=/some/path' OR '-input=\"Some text to generate against\"",
		)
	}

	return cliOpts
}
