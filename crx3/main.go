package main

import (
	"fmt"
	"os"

	"crx/crx3/command"
)

func main() {
	cli := command.New()
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
