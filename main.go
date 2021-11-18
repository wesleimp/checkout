package main

import (
	"errors"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/urfave/cli/v2"

	"github.com/wesleimp/checkout/internal/branch"
	"github.com/wesleimp/checkout/internal/git"
)

var (
	version = "v0.1.0"
)

func main() {
	app := &cli.App{
		Name:      "checkout",
		Usage:     "Fuzzy-finding for checkouted branches in reflog",
		UsageText: "checkout",
		Version:   version,
		Authors: []*cli.Author{{
			Name:  "Weslei Juan Novaes Pereira",
			Email: "wesleimsr@gmail.com",
		}},
		Action: func(c *cli.Context) error {
			if !git.IsRepo() {
				return errors.New("It's not a git repository")
			}

			current_branch, err := git.Run("rev-parse", "--abbrev-ref", "HEAD")
			if err != nil {
				return err
			}

			branches, err := branch.Checkouts(10000)
			if err != nil {
				return err
			}

			if !contains(branches, current_branch) {
				branches = append(branches, current_branch)
			}

			idx, err := fuzzyfinder.Find(branches, func(i int) string {
				return branches[i]
			})
			if err != nil {
				if err.Error() != "abort" {
					return err
				}
				return nil
			}

			branch := branches[idx]
			out, err := git.Run("checkout", branch)
			if err != nil {
				return err
			}

			println(out)
			println("switched to", branch)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
