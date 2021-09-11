package cmd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TODO: add tests for resolving paths
func TestNewFilepath(t *testing.T) {
	tests := []struct {
		args  []string
		cwd   string
		kakwd string
		want  Filepath
	}{
		{
			[]string{"file"},
			"/home/k/p/kks/",
			"/home/k/p/kks/",
			Filepath{Name: "file", Raw: []string{"file"}},
		},
		{
			[]string{"file.kak", "+22"},
			"/home/k/p/",
			"/home/k/p/kks/",
			Filepath{Name: "file.kak", Line: 22, Raw: []string{"file.kak", "+22"}},
		},
		{
			[]string{"readme", "+10:2"},
			"/home/k/p/kks/",
			"/home/k/p/kks/",
			Filepath{Name: "readme", Line: 10, Column: 2, Raw: []string{"readme", "+10:2"}},
		},
		{
			[]string{"readme", ":2"},
			"/home/k/p/kks/",
			"/home/k/p/kks/",
			Filepath{Name: "readme", Raw: []string{"readme", ":2"}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := NewFilepath(tt.args, "", "")
			if err != nil {
				t.Fatal("can't get Filepath: ", err)
			}
			if !cmp.Equal(tt.want, *got) {
				t.Errorf("Filepath mismatch:\n%s", cmp.Diff(tt.want, got))
			}
		})
	}
}
