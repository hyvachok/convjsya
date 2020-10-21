package cmd

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() {
	opt, err := parse()
	check(err)

	result, err := convert(opt)
	check(err)

	err = out(result, opt)
	check(err)
}
