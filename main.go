package main

import (
	"github.com/wcdi/wcdi-userstore/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
