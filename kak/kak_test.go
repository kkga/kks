package kak

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSessionDir(t *testing.T) {
	tests := []struct {
		kakdir string
		want   string
	}{
		{
			"/home/kkga",
			"/home/kkga",
		},
		{
			"~/downloads/",
			"/home/kkga/downloads",
		},
		{
			"/tmp/",
			"/tmp",
		},
	}
	for i, tt := range tests {
		t.Run("", func(t *testing.T) {
			testSession, err := Start(fmt.Sprintf("kks-test-%d", i))
			if err != nil {
				t.Fatal(err)
			}

			kctx := &Context{}
			kctx.Session = Session{testSession}

			defer Send(kctx, "kill")

			if err := Send(kctx, fmt.Sprintf("cd %s", tt.kakdir)); err != nil {
				t.Fatal(err)
			}

			got, err := kctx.Session.Dir()
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(tt.want, got) {
				t.Errorf("Sessiond.Dir() mismatch:\n%s", cmp.Diff(tt.want, got))
			}

		})
	}
}

func TestSessionExists(t *testing.T) {
	tests := []struct {
		session Session
	}{
		{
			Session{"kks-test-hey"},
		},
		{
			Session{"kks-test-yo"},
		},
		{
			Session{"kks-wassup12348fkqwer-qw"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			_, err := Start(tt.session.Name)
			if err != nil {
				t.Fatal(err)
			}
			kctx := &Context{}
			kctx.Session = tt.session
			defer Send(kctx, "kill")

			got, err := kctx.Session.Exists()
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(true, got) {
				t.Errorf("Sessiond.Dir() mismatch:\n%s", cmp.Diff(true, got))
			}

		})
	}
}
