package bytefmt

import (
	"context"
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

		SilenceErrors:     true,
		DisableAutoGenTag: true,
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
	var exitErr exiterr.ExitError
	for _, arg := range args {
		arg = strings.TrimSpace(arg)

		if cfg.Invert {
			v, err := bytefmt.Decode(arg)
			if err != nil {
				cmd.PrintErrln(cmd.ErrPrefix(), "bytefmt: "+err.Error())
				exitErr.Code = 1
				continue
			}

			cmd.Println(v)
		} else {
			v, err := strconv.ParseInt(arg, 10, 64)
			if err != nil {
				var err2 error
				if v, err2 = bytefmt.Decode(arg); err2 != nil {
					cmd.PrintErrln(cmd.ErrPrefix(), "bytefmt: "+err.Error())
					exitErr.Code = 1
					continue
				}
			}

			cmd.Println(encode(v))
		}
	}

	return exitErr
}
