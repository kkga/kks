package kak

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
