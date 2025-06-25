package main

import (
	"os"
	"path/filepath"
	"strings"

	"gabe565.com/byte-utils/internal/cmd/bytect"
	"gabe565.com/byte-utils/internal/cmd/bytefmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func main() {
	const output = "./docs"

	if err := os.RemoveAll(output); err != nil {
		panic(err)
	}

	if err := os.MkdirAll(output, 0o755); err != nil {
		panic(err)
	}

	for _, cmd := range []*cobra.Command{bytefmt.New(), bytect.New()} {
		f, err := os.Create(filepath.Join(output, strings.ToLower(cmd.Name())+".md"))
		if err != nil {
			panic(err)
		}

		if err := doc.GenMarkdown(cmd, f); err != nil {
			panic(err)
		}

		if err := f.Close(); err != nil {
			panic(err)
		}
	}
}
