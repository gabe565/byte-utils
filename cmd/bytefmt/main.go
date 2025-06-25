package main

import (
	"errors"
	"fmt"
	"os"

	"gabe565.com/byte-utils/internal/exiterr"
)

func main() {
	if err := New().Execute(); err != nil {
		var exitErr exiterr.ExitErr
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.Code)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
