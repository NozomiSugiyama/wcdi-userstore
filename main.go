package main

import (
	"github.com/wcdi/wcdi-auth/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
