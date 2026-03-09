package main

import (
	"errors"
	"os"

	"gabe565.com/byte-utils/internal/cmd/bytefmt"
	"gabe565.com/byte-utils/internal/exiterr"
	"gabe565.com/utils/cobrax"
)

func main() {
	cmd := bytefmt.New(cobrax.WithVersion(""))
	if err := cmd.Execute(); err != nil {
		if err, ok := errors.AsType[exiterr.ExitError](err); ok {
			os.Exit(err.Code)
		}
		cmd.PrintErrln(cmd.ErrPrefix(), err)
		os.Exit(1)
	}
}
