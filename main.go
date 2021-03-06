package main

import (
	"fmt"
	"os"

	"github.com/NozomiSugiyama/wcdi-userstore/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
