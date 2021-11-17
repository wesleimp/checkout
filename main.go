package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

type Checkout struct {
	from string
	to   string
}

func main() {
	cmd := exec.Command("git", "reflog")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		println(bufio.NewScanner(stdout).Text())
		println(bufio.NewScanner(stderr).Text())
		os.Exit(1)
	}

	counter := 0
	bm := map[string]bool{}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() && counter < 10000 {
		line := scanner.Text()
		if strings.Contains(line, "checkout: moving from") {
			bb := strings.Split(strings.Split(line, "checkout: moving from")[1], "to")
			for _, b := range bb {
				trimmed := strings.TrimSpace(b)
				if trimmed != "" {
					bm[trimmed] = true
				}
			}
		}
		counter++
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}

	branches := []string{}
	for b := range bm {
		branches = append(branches, b)
	}

	if len(branches) == 1 {
		println("No branches found")
		os.Exit(0)
	}

	idx, err := fuzzyfinder.Find(branches, func(i int) string {
		return branches[i]
	})
	if err != nil {
		panic(err.Error())
	}
	println("selected %v", idx)

}
