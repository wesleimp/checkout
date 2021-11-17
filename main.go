package main

import (
	"github.com/ktr0731/go-fuzzyfinder"

	"github.com/wesleimp/checkout/internal/branch"
	"github.com/wesleimp/checkout/internal/git"
)

type Checkout struct {
	from string
	to   string
}

func main() {
	if !git.IsRepo() {
		println("It's not a git repository")
	}

	current_branch, err := git.Run("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		panic(err.Error())
	}

	branches, err := branch.Checkouts(10000)
	if err != nil {
		panic(err.Error())
	}

	if !contains(branches, current_branch) {
		branches = append(branches, current_branch)
	}

	idx, err := fuzzyfinder.Find(branches, func(i int) string {
		return branches[i]
	})
	if err != nil {
		panic(err.Error())
	}

	println("selected branch:", branches[idx])
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
