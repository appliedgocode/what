package what

// IMPORTANT: to run these tests, use build tag "what":
//
// go test -tags what

import (
	"bytes"
	"log"
	"regexp"
	"testing"
)

func TestAll(t *testing.T) {
	got := &bytes.Buffer{}
	log.SetOutput(got)
	log.SetFlags(0)
	n := 23

	// no package name set - all packages are enabled
	enabled = map[string]bool{}
	Happens("what.Happens - all packages enabled")
	If(true, "If true")
	If(false, "If false")
	Is(n)
	Func()
	Package()

	enabled = map[string]bool{
		"what": true,
	}
	Happens("what.Happens - package 'what' enabled")

	enabled = map[string]bool{
		"appliedgo.net/what": true,
	}
	Happens("what.Happens - package 'appliedgo.net/what' enabled")

	enabled = map[string]bool{
		"someotherpackage": true,
	}
	Happens("what.Happens - package 'what' NOT enabled") // this should not print

	wantRE := regexp.MustCompilePOSIX(`appliedgo.net/what\.TestAll: what\.Happens - all packages enabled
appliedgo\.net/what\.If: If true
\(int\) 23
Func appliedgo.net/what\.TestAll in line \d\+ of file .*/what_test.go
Package appliedgo.net/what
appliedgo\.net/what\.TestAll: what.Happens - package 'what' enabled
appliedgo\.net/what\.TestAll: what.Happens - package 'appliedgo\.net/what' enabled
`)
	if !wantRE.Match(got.Bytes()) {
		t.Errorf("Got: %s\n\nWant: %s", got, wantRE.Find(got.Bytes()))
	}
}
