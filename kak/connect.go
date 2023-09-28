package kak

func Connect(session string, fp *Filepath) error {
	return Run(session, []string{"-c"}, fp)
}
