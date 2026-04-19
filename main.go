package main

import (
	"fmt"
	"os"

	"rssc/internal"
)

func main() {
    conn, err := internal.Connect();
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err);
		os.Exit(1);
    }
	defer conn.Close();

    argv := os.Args[1:];

    err = internal.ParseArgs(argv);
    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR: %v\n", err);
		os.Exit(1);
    }
}

