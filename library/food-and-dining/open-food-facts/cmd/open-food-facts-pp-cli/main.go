// Copyright 2026 Dhilip Subramanian and contributors. Licensed under Apache-2.0.
// Modified by bricenice17 to add store-aware search support.

package main

import (
	"fmt"
	"os"

	"github.com/bricenice17/printing-press-library/library/food-and-dining/open-food-facts/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(cli.ExitCode(err))
	}
}
