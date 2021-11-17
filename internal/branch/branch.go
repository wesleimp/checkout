package branch

import (
	"bufio"
	"os/exec"
	"strings"
)

func Checkouts(buffer int) ([]string, error) {
	cmd := exec.Command("git", "reflog")

	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	_, err = cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	counter := 0
	bm := map[string]bool{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() && counter < 10000 {
		line := scanner.Text()
		if strings.Contains(line, "checkout: moving from") {
			branches := strings.Split(strings.Split(line, "checkout: moving from")[1], "to")
			for _, b := range branches {
				trimmed := strings.TrimSpace(b)
				if trimmed != "" {
					bm[trimmed] = true
				}
			}
		}
		counter++
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	branches := []string{}
	for b := range bm {
		branches = append(branches, b)
	}

	return branches, nil
}
