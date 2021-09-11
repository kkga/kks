package cmd

import (
	"strconv"
	"strings"
)

type Filepath struct {
	Name   string   `json:"name"`
	Line   int      `json:"line"`
	Column int      `json:"column"`
	Raw    []string `json:"raw"`
}

func NewFilepath(args []string, cmdWd string, kakWd string) (*Filepath, error) {
	fp := &Filepath{Raw: args}

	if len(args) > 0 {
		name, line, col, err := fp.parse()
		if err != nil {
			return nil, err
		}
		fp.Name = name
		fp.Line = line
		fp.Column = col
	}

	return fp, nil
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
	return name, line, col, nil
}

// 	// TODO: this path resolution needs to happen in Edit

// 	// if strings.Contains(arg, "buflist") {
// 	// 	cwd, err := os.Getwd()
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	fmt.Println("CWD:", cwd)

// 	// kakwd, err := kak.Get("%sh{pwd}", context.session, context.client)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	fmt.Println("KAKWD:", kakwd[0])

// 	// 	relPath, _ := filepath.Rel(cwd, kakwd[0])
// 	// 	if strings.HasPrefix(relPath, "home/") {
// 	// 		relPath = strings.Replace(relPath, "home/", "~/", 1)
// 	// 	}
// 	// 	fmt.Println("rel path:", relPath)
// 	// 	fmt.Println()

// 	// 	for i, buf := range out {
// 	// 		// if !strings.HasPrefix(buf, "~") && !strings.HasPrefix(buf, "*") {
// 	// 		// }
// 	// 		if !filepath.IsAbs(buf) && !strings.HasPrefix(buf, "*") {
// 	// 			out[i] = filepath.Join(relPath, buf)
// 	// 		} else {
// 	// 			out[i] = buf
// 	// 		}
// 	// 	}
// 	// }
