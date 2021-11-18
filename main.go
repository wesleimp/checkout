package main

import (
	"os"

	"github.com/ktr0731/go-fuzzyfinder"

	"github.com/wesleimp/checkout/internal/branch"
	"github.com/wesleimp/checkout/internal/git"
)

func main() {
	if !git.IsRepo() {
		println("It's not a git repository")
		os.Exit(1)
	}

	current_branch, err := git.Run("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	branches, err := branch.Checkouts(10000)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if !contains(branches, current_branch) {
		branches = append(branches, current_branch)
	}

	idx, err := fuzzyfinder.Find(branches, func(i int) string {
		return branches[i]
	})
	if err != nil {
		if err.Error() != "abort" {
			println(err.Error())
		}
		os.Exit(0)
	}

	branch := branches[idx]
	out, err := git.Run("checkout", branch)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println(out)
	println("switched to", branch)
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
