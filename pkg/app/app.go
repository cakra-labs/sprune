package app

import (
	"fmt"
	"os"
)

func Run() {
	cli := newCLI()

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
