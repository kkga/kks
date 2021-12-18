package kak

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Filepath struct {
	Name   string
	Line   int
	Column int
	Raw    []string
}

func NewFilepath(args []string) *Filepath {
	fp := &Filepath{Raw: args}

	if len(args) > 0 {
		name, line, col, err := fp.parse()
		if err != nil {
			return nil
		}
		fp.Name = name
		fp.Line = line
		fp.Column = col
	}

	return fp
}

func (fp *Filepath) Dir() (dir string, err error) {
	info, err := os.Stat(fp.Name)
	if err != nil {
		return "", err
	}

	if info.IsDir() {
		dir = fp.Name
	} else {
		dir = path.Dir(fp.Name)
	}

	return
}

func (fp *Filepath) ParseGitDir() string {
	dir, _ := fp.Dir()
	gitOut, err := exec.Command("git", "-C", dir, "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(path.Base(string(gitOut)), ".", "-"))
}

func (fp *Filepath) parse() (absName string, line, col int, err error) {
	r := fp.Raw

	rawName := r[0]

	if filepath.IsAbs(rawName) {
		absName = rawName
	} else {
		cwd, _ := os.Getwd()
		absName = path.Join(cwd, rawName)
	}

	if len(r) > 1 && strings.HasPrefix(r[1], "+") {
		if strings.Contains(r[1], ":") {
			lineStr := strings.ReplaceAll(strings.Split(r[1], ":")[0], "+", "")
			lineInt, err := strconv.Atoi(lineStr)
			if err != nil {
				return "", 0, 0, err
			}
			line = lineInt

			colStr := strings.Split(r[1], ":")[1]
			colInt, err := strconv.Atoi(colStr)
			if err != nil {
				return "", 0, 0, err
			}
			col = colInt
		} else {
			lineStr := strings.ReplaceAll(r[1], "+", "")
			lineInt, err := strconv.Atoi(lineStr)
			if err != nil {
				return "", 0, 0, err
			}
			line = lineInt
		}
	}

	return absName, line, col, err
}
