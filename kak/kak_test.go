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
			"/etc/",
			"/etc",
		},
	}
	for i, tt := range tests {
		t.Run("", func(t *testing.T) {
			testSession, err := Start(fmt.Sprintf("kks-test-%d", i))
			if err != nil {
				t.Fatal(err)
			}

			defer func() {
				err = Send(testSession, "", "", "kill", nil)
			}()

			if err := Send(testSession, "", "", fmt.Sprintf("cd %s", tt.kakdir), nil); err != nil {
				t.Fatal(err)
			}

			got, err := SessionDir(testSession)
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
	sessions := []string{
		"kks-test-hey",
		"kks-test-yo",
		"kks-wassup12348fkqwer-qw",
	}
	for _, session := range sessions {
		t.Run("", func(t *testing.T) {
			_, err := Start(session)
			if err != nil {
				t.Fatal(err)
			}

			defer func() {
				err = Send(session, "", "", "kill", nil)
			}()

			got, err := SessionExists(session)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(true, got) {
				t.Errorf("Sessiond.Dir() mismatch:\n%s", cmp.Diff(true, got))
			}
		})
	}
}
