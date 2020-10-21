package cmd

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io"
	"io/ioutil"
	"os"
)

func source(opt Options) (io.Reader, error) {
	if opt.piped {
		return os.Stdin, nil
	}

	f, err := os.Open(opt.input)
	if err != nil {
		return nil, err
	}
	return f, nil

}

func out(result []byte, opt Options) error {
	if opt.output == "" {
		fmt.Print(string(result))
	} else {
		file, err := os.Create(opt.output)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write(result)
		if err != nil {
			return err
		}
	}
	return nil
}

func isYaml(in string) bool {
	if len(in) < 3 {
		return false
	}
	if in[len(in)-4:] == "yaml" || in[len(in)-3:] == "yml" {
		return true
	}
	return false
}

func convert(opt Options) ([]byte, error) {
	src, err := source(opt)
	if err != nil {
		return nil, err
	}

	in, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}

	var result []byte
	if isYaml(opt.input) || opt.yaml {
		result, err = yaml.YAMLToJSON(in)
	} else {
		result, err = yaml.JSONToYAML(in)
	}

	return result, err
}
