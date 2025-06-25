package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"gabe565.com/byte-utils/internal/config"
	"gabe565.com/byte-utils/internal/exiterr"
	"gabe565.com/utils/bytefmt"
	"gabe565.com/utils/cobrax"
	"gabe565.com/utils/termx"
	"github.com/spf13/cobra"
)

func New(opts ...cobrax.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bytefmt [number]...",
		Short: "Pretty-print bytes",
		RunE:  run,

		SilenceErrors: true,
	}

	cfg := config.NewBytefmt()
	cfg.RegisterFlags(cmd)
	cmd.SetContext(config.NewContext(context.Background(), cfg))

	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	cfg := config.Load[config.Bytefmt](cmd)
	cmd.SilenceUsage = true

	if len(args) == 0 {
		if termx.IsTerminal(os.Stdin) {
			return cmd.Usage()
		}

		b, err := io.ReadAll(io.LimitReader(os.Stdin, 64))
		if err != nil {
			return err
		}

		args = []string{string(b)}
	}

	encode := cfg.NewEncodeFunc()
	var exitErr exiterr.ExitErr
	for _, arg := range args {
		arg = strings.TrimSpace(arg)

		if v, encodeErr := strconv.ParseInt(arg, 10, 64); encodeErr == nil {
			fmt.Println(encode(v))
		} else {
			v, decodeErr := bytefmt.Decode(arg)
			if decodeErr != nil {
				fmt.Println("bytefmt: input could not be encoded or decoded: encode: " + encodeErr.Error() + "; decode: " + decodeErr.Error())
				exitErr.Code = 1
				continue
			}

			fmt.Println(v)
		}
	}

	return exitErr
}
