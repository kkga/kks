package kak

import (
	"fmt"
	"strconv"
	"strings"
)

type Filepath struct {
	Name   string   `json:"name"`
	Line   int      `json:"line"`
	Column int      `json:"column"`
	Raw    []string `json:"raw"`
}

func NewFilepath(args []string) (*Filepath, error) {
	fp := Filepath{Raw: args}

	name, line, col, err := fp.parse()
	if err != nil {
		return new(Filepath), err
	}

	fp.Name = name
	fp.Line = line
	fp.Column = col

	return &fp, nil
}

func (fp *Filepath) parse() (name string, line int, col int, err error) {
	r := fp.Raw
	name = r[0]

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

	fmt.Println(name, line, col)

	return name, line, col, nil
}
