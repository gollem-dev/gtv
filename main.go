package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	if err := newCommand().Run(context.Background(), os.Args); err != nil {
		slog.Error("command failed", slog.Any("error", err))
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func newCommand() *cli.Command {
	return &cli.Command{
		Name:  "gtv",
		Usage: "Gollem Trace Viewer - start a web server to browse gollem traces",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "addr",
				Value:   ":18900",
				Sources: cli.EnvVars("GTV_ADDR"),
				Usage:   "Server listen address",
			},
			&cli.StringFlag{
				Name:    "dir",
				Sources: cli.EnvVars("GTV_DIR"),
				Usage:   "Local directory containing trace JSON files",
			},
			&cli.StringFlag{
				Name:    "gs",
				Sources: cli.EnvVars("GTV_GS"),
				Usage:   "Google Cloud Storage URI (e.g. gs://bucket/prefix/)",
			},
			&cli.BoolFlag{
				Name:    "no-browser",
				Sources: cli.EnvVars("GTV_NO_BROWSER"),
				Usage:   "Do not open browser automatically",
			},
		},
		Action: runView,
	}
}
