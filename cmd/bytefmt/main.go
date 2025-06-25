package main

import (
	"errors"
	"os"

	"gabe565.com/byte-utils/internal/exiterr"
)

func main() {
	cmd := New()
	if err := cmd.Execute(); err != nil {
		var exitErr exiterr.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.Code)
		}
		cmd.PrintErrln(cmd.ErrPrefix(), err)
		os.Exit(1)
	}
}
