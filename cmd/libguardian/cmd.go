package main

import (
	"errors"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/crashdump/libguardian/pkg"
	"github.com/crashdump/libguardian/pkg/gomod"
)

var logger *logging

const STDOUT = "stdout"

func init() {
	cli.HelpFlag = &cli.BoolFlag{Name: "help"}
	cli.VersionFlag = &cli.BoolFlag{Name: "version", Aliases: []string{"v"}}
}

var errorPathMissing = "You need to specify the path to the source code folder"

func main() {
	logger = newLogger()
	var flagOutput string

	logger.print("┌──────────────┐")
	logger.print("│ Lib Guardian │")
	logger.print("└──────────────┘")
	logger.print("")

	app := &cli.App{
		Name:     "libguardian",
		Usage:    "Walk files in a directory and identifies gaps between project and inventory.",
		Compiled: time.Now(),
		Authors: []*cli.Author{{
			Name:  "Adrien Pujol",
			Email: "ap@cdfr.net",
		}},
		Commands: []*cli.Command{
			{
				Name:    "enumerate",
				Aliases: []string{"e"},
				Usage:   "enumerate all the libraries from source code.",
				Before: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() < 1 {
						logger.print(errorPathMissing)
						os.Exit(1)
					}
					return nil
				},
				Action: func(cCtx *cli.Context) error {
					path := cCtx.Args().First()

					logger.printHeader("Enumerating libraries from source code...")

					lg := pkg.NewLibGuardian[gomod.Library](path, gomod.GoMod[gomod.Library]{})
					results, err := lg.Enumerate()
					if err != nil {
						logger.printFatal(err.Error())
					}
					logger.printfResult("Found %d files.", len(results))
					logger.print("")
					for _, result := range results {
						logger.printResult(result.Module)
					}

					return nil
				},
			},
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "generate a config.json from source code.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "output",
						Aliases:     []string{"o"},
						Destination: &flagOutput,
						Action: func(ctx *cli.Context, v string) error {
							if v == "" {
								return errors.New("missing output filename")
							}
							return nil
						},
					},
				},
				Before: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() < 1 {
						logger.print(errorPathMissing)
						os.Exit(1)
					}
					return nil
				},
				Action: func(cCtx *cli.Context) error {
					//path := cCtx.Args().First()

					logger.printHeader("Generating sbom.jsom from source code...")
					panic("not implemented")
					//results, err := libguardian.Enumerate(path)
					//if err != nil {
					//	logger.printFatal(err.Error())
					//}
					//logger.printfResult("Found %d files.", len(results))
					//logger.print("")

					return nil
				},
			},
			{
				Name:    "enforce",
				Aliases: []string{"v"},
				Usage:   "enforce inventory libraries.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "config",
						Aliases:     []string{"c"},
						Destination: &flagOutput,
						Action: func(ctx *cli.Context, v string) error {
							if v == "" {
								return errors.New("missing config filename")
							}
							return nil
						},
					},
				},
				Before: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() < 1 {
						logger.print(errorPathMissing)
						os.Exit(1)
					}
					return nil
				},
				Action: func(cCtx *cli.Context) error {
					path := cCtx.Args().First()

					logger.printHeader("Searching for foreign libraries in source code...")

					lg := pkg.NewLibGuardian[gomod.Library](path, gomod.GoMod[gomod.Library]{})
					err := lg.Config.Load(cCtx.String("config"))
					if err != nil {
						return err
					}

					results, err := lg.Enforce()
					if err != nil {
						return err
					}

					if len(results) > 0 {
						logger.printfResult("Found %d foreign libraries.", len(results))
						logger.print("")
						for _, result := range results {
							logger.printResult(result.Module)
						}

						logger.printFatal("Failed!")
					}
					logger.print("No mismatch, success!")

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.printFatal(err.Error())
	}
}
