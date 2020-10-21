package cmd

import (
	"errors"
	"flag"
	"os"
)

type Options struct {
	input  string
	output string
	piped  bool
	yaml   bool
}

func parse() (Options, error) {
	var input = flag.String("i", "", "Defines the input file")
	var output = flag.String("o", "", "Defines the output file")
	var yaml2json = flag.Bool("yaml", false, "Flag indicating if to convert YAML to JSON, used only in piped mode")
	var piped = false
	flag.Parse()

	if *input == "" && len(os.Args) > 1 {
		*input = os.Args[1]
	}
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 { // whether piped mode or console output is used
		piped = true
	}
	if *input == "" && len(os.Args) <= 1 && !piped {
		return Options{}, errors.New("no input source")
	}

	return Options{
		input:  *input,
		output: *output,
		piped:  piped,
		yaml:   *yaml2json,
	}, nil
}
