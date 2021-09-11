package cmd

import (
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Filepath struct {
	Name   string   `json:"name"`
	Line   int      `json:"line"`
	Column int      `json:"column"`
	Raw    []string `json:"raw"`
}

func NewFilepath(args []string) (fp *Filepath, err error) {
	fp = &Filepath{Raw: args}

	if len(args) > 0 {
		name, line, col, err := fp.parse()
		if err != nil {
			return nil, err
		}
		fp.Name = name
		fp.Line = line
		fp.Column = col
	}

	return
}

func (fp *Filepath) parse() (absName string, line int, col int, err error) {
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
