package main

import (
	"fmt"
	"os"
)

type CLI struct {
}

const Usage = `
	check --data DATA
	update --data DATA WHERE
`

func (cli *CLI) Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}

	switch args[1] {
	case "update":
		if len(args) != 5 {
			fmt.Println(Usage)
			return
		}
		cli.check(args[3], args[4])
		break

	default:
		fmt.Println(Usage)

	}
}
