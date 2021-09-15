package kak

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDir(t *testing.T) {
	tests := []struct {
		fp   Filepath
		want string
	}{
		{
			Filepath{Name: "/home/kkga"},
			"/home/kkga",
		},
		{
			Filepath{Name: "/home/kkga/README.md"},
			"/home/kkga",
		},
		{
			Filepath{Name: "/home/kkga/projects/kks/"},
			"/home/kkga/projects/kks/",
		},
		{
			Filepath{Name: "/home/kkga/projects/kks/cmd/attach.go"},
			"/home/kkga/projects/kks/cmd",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := tt.fp.Dir()
			if err != nil {
				t.Fatal("can't get Filepath.Dir(): ", err)
			}
			if !cmp.Equal(tt.want, got) {
				t.Errorf("Filepath.Dir() mismatch:\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestParseGitDir(t *testing.T) {
	tests := []struct {
		fp   Filepath
		want string
	}{
		{
			Filepath{Name: "/home/kkga"},
			"",
		},
		{
			Filepath{Name: "/home/kkga/README.md"},
			"",
		},
		{
			Filepath{Name: "/home/kkga/projects/kks/"},
			"kks",
		},
		{
			Filepath{Name: "/home/kkga/projects/kks/cmd/attach.go"},
			"kks",
		},
		{
			Filepath{Name: "/home/kkga/projects/foot.kak/rc/foot.kak"},
			"foot-kak",
		},
		{
			Filepath{Name: "/home/kkga/repos/kakoune/rc/detection/editorconfig.kak"},
			"kakoune",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.fp.ParseGitDir()
			if !cmp.Equal(tt.want, got) {
				t.Errorf("Filepath.ParseGitDir() mismatch:\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestNewFilepath(t *testing.T) {
	tests := []struct {
		args []string
		want Filepath
	}{
		{
			[]string{"file"},
			Filepath{Name: "/home/kkga/projects/kks/kak/file",
				Raw: []string{"file"}},
		},
		{
			[]string{"../file.kak", "+22"},
			Filepath{Name: "/home/kkga/projects/kks/file.kak",
				Line: 22,
				Raw:  []string{"../file.kak", "+22"}},
		},
		{
			[]string{"/etc/readme", "+10:2"},
			Filepath{Name: "/etc/readme",
				Line: 10, Column: 2,
				Raw: []string{"/etc/readme", "+10:2"}},
		},
		{
			[]string{"../../../downloads/readme", ":2"},
			Filepath{Name: "/home/kkga/downloads/readme",
				Raw: []string{"../../../downloads/readme", ":2"}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := NewFilepath(tt.args)
			if err != nil {
				t.Fatal("can't get Filepath: ", err)
			}
			if !cmp.Equal(tt.want, *got) {
				t.Errorf("Filepath mismatch:\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}
