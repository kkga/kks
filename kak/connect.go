package kak

func Connect(kctx *Context, fp *Filepath) error {
	return Run(kctx, []string{"-c"}, fp)
}
