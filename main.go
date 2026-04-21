package main

import (
	"fmt"
	"os"

	"rssc/internal/cli"
)

func main() {
    argv := os.Args[1:];

	err := cli.ParseArgs(argv);
    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %v\n", err);
		os.Exit(1);
    }
}

