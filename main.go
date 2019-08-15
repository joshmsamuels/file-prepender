package main

import (
	"fmt"
	"os"
)

func main() {
	src, dsts := validateArgs(os.Args)
	if src == "" || len(dsts) == 0 {
		os.Exit(1)
	}

}

func validateArgs(args []string) (string, []string) {
	if args == nil || len(args) < 2 {
		fmt.Println("Incorrect usage")
		fmt.Println("Correct usage: <file to prepend> ...<files/dirs to append to>")

		return "", nil
	}

	var destStrs []string

	return args[0], append(destStrs, args[1:]...)
}
