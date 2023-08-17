package main

import (
	"errors"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/crashdump/venlock/pkg"
	"github.com/crashdump/venlock/pkg/gomod"
	"github.com/crashdump/venlock/pkg/maven"
	"github.com/crashdump/venlock/pkg/npm"
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

	logger.print("┌─────────────┐")
	logger.print("│ Vendor Lock │")
	logger.print("└─────────────┘")
	logger.print("")

	app := &cli.App{
		Name:     "venlock",
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

					logger.print("Enumerating libraries from source code...")

					logger.printHeader("Go...")
					enumerate[gomod.GoMod[gomod.Library], gomod.Library](path)

					logger.printHeader("Maven...")
					enumerate[maven.Maven[maven.Library], maven.Library](path)

					logger.printHeader("NPM...")
					enumerate[npm.Npm[npm.Library], npm.Library](path)

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
					logger.printHeader("Generating sbom.jsom from source code...")
					panic("not yet implemented")
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
						Required:    true,
						Destination: &flagOutput,
						Aliases:     []string{"c"},
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
					config := cCtx.String("config")

					logger.print("Searching for foreign libraries in source code...")

					var found []string

					logger.printHeader("Go...")
					f, err := enforce[gomod.GoMod[gomod.Library], gomod.Library](config, path)
					if err != nil {
						logger.printFatal(err.Error())
					}
					found = append(found, f...)

					logger.printHeader("Maven...")
					f, err = enforce[maven.Maven[maven.Library], maven.Library](config, path)
					if err != nil {
						logger.printFatal(err.Error())
					}
					found = append(found, f...)

					logger.printHeader("Npm...")
					f, err = enforce[npm.Npm[npm.Library], npm.Library](config, path)
					if err != nil {
						logger.printFatal(err.Error())
					}
					found = append(found, f...)

					if len(found) == 0 {
						logger.print("")
						logger.print("No mismatch, success!")
						return nil
					}
					return errors.New("found unexpected libraries")
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.printFatal(err.Error())
	}
}

func enumerate[S pkg.Scanner[L], L pkg.Library](targetDir string) {
	lg := pkg.NewVenlock[L](targetDir, *new(S))
	results, err := lg.Enumerate()
	if err != nil {
		logger.printFatal(err.Error())
	}
	logger.printfResult("... found %d dependencies.", len(results))
	logger.print("")
	for _, result := range results {
		logger.printResult(result.String())
	}
}

func enforce[S pkg.Scanner[L], L pkg.Library](config string, targetDir string) ([]string, error) {
	lg := pkg.NewVenlock[L](targetDir, *new(S))

	var cfg pkg.Config[L]
	err := cfg.Load(config)
	if err != nil {
		return []string{}, err
	}

	found, err := lg.Enforce(cfg.Catalogue[lg.Scanner.String()])
	if err != nil {
		return []string{}, err
	}

	var results []string
	for _, f := range found {
		results = append(results, f.String())
	}

	return results, nil
}
