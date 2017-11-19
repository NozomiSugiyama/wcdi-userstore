package main

import (
	"os"
	"fmt"

	"github.com/wcdi/wcdi-userstore/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
