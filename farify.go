/*
Copyright © 2021 Akbar Fauzi <fauzia84@ymail.com>

*/
package main

import (
	"os"

	"github.com/akbarfa49/farify/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
