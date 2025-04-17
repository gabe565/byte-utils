package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"gabe565.com/bytefmt/internal/config"
	"gabe565.com/bytefmt/internal/exiterr"
	"gabe565.com/utils/cobrax"
	"gabe565.com/utils/termx"
	"github.com/spf13/cobra"
)

func New(opts ...cobrax.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bytect [file]...",
		Short: "Pretty-print the size of files or stdin",
		RunE:  run,

		SilenceErrors: true,
	}

	cfg := config.NewBytect()
	cfg.RegisterFlags(cmd)
	cmd.SetContext(config.NewContext(context.Background(), cfg))

	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	cfg := config.Load[config.Bytect](cmd)

	cmd.SilenceUsage = true
	encode := cfg.NewEncodeFunc()

	var exitErr exiterr.ExitErr
	if len(args) != 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight)
		for _, arg := range args {
			stat, err := os.Stat(arg)
			if err != nil {
				return err
			}

			if stat.IsDir() {
				exitErr.Code = 1
				fmt.Println("bytect: " + arg + ": is a directory")
				continue
			}

			if _, err := w.Write(
				[]byte(encode(stat.Size()) + "\t  " + arg + "\n"),
			); err != nil {
				return err
			}
		}
		if err := w.Flush(); err != nil {
			return err
		}
	} else {
		if termx.IsTerminal(os.Stdin) {
			return cmd.Usage()
		}

		n, err := io.Copy(io.Discard, os.Stdin)
		if err != nil {
			return err
		}

		fmt.Println(encode(n))
	}
	return exitErr
}
