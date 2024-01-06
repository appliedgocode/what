//go:build what || whatpackage || whatis || whatfunc || whathappens || whatif

package what

// IMPORTANT: to run these tests, use build tag "what":
//
// go test -tags what

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestOutput(t *testing.T) {
	got := &bytes.Buffer{}
	log.SetOutput(NewTeeWriter(got, os.Stderr)) // write all log output into "got" for later matching
	log.SetFlags(0)                             // no extra decorations
	n := 23

	// no package name set - all packages are enabled
	enabled = map[string]bool{}
	Happens("what.Happens")
	verify(t, got, `appliedgo\.net/what\.TestOutput: what\.Happens`)
	got.Reset()
	// Structured output should get colorized:
	Happens("WARN", "Something is fishy", "code", 23, "err", "Quartotube exceeded max flotic burst")
	verify(t, got, strings.Replace(dim+`\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d`+reset+` `+yellow+`WARN`+reset+" "+blue+`appliedgo.net/what.TestOutput`+reset+` `+white+underscore+`Something is fishy`+reset+" "+cyan+`code`+reset+`=`+magenta+`23`+reset+` `+cyan+`err`+reset+`=`+red+bold+`Quartotube exceeded max flotic burst`+reset+" ", `[`, `\[`, -1))
	got.Reset()
	If(true, "If true")
	verify(t, got, `appliedgo\.net/what\.TestOutput: If true`)
	got.Reset()
	If(false, "If false")
	verify(t, got, ``)
	got.Reset()
	Is(n)
	verify(t, got, `\(int\) 23`)
	got.Reset()
	Func()
	verify(t, got, `Func appliedgo\.net/what\.TestOutput in line \d+ of file .*/what_test\.go`)
	got.Reset()
	Package()
	verify(t, got, `Package appliedgo\.net/what`)
}

func verify(t *testing.T, gotbuf *bytes.Buffer, want string) {
	wantRE := regexp.MustCompile("^" + want + "$")
	got := gotbuf.String()[:max(0, gotbuf.Len()-1)] // trim the trailing \n
	if !wantRE.MatchString(got) {
		t.Errorf("Got: `%s` Want: `%s`", got, want)
	}
}

func TestEnabling(t *testing.T) {
	got := &bytes.Buffer{}
	log.SetOutput(got) // write all log output into "got" for later matching
	log.SetFlags(0)    // no extra decorations

	// no package name set - all packages are enabled
	enabled = map[string]bool{}
	Happens("what.Happens - all packages enabled")

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

	wantRE := regexp.MustCompile(`appliedgo.net/what\.TestEnabling: what\.Happens - all packages enabled
appliedgo\.net/what\.TestEnabling: what.Happens - package 'what' enabled
appliedgo\.net/what\.TestEnabling: what.Happens - package 'appliedgo\.net/what' enabled
`)
	// "got" contains all log output from the above calls
	if !wantRE.Match(got.Bytes()) {
		t.Errorf("Got: %s\n\nWant: %s", got, wantRE)
	}
}

// for pre-Go 1.21 clients
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TeeWriter struct {
	writers []io.Writer
}

// NewTeeWriter creates a new TeeWriter with the given io.Writers.
func NewTeeWriter(writers ...io.Writer) *TeeWriter {
	return &TeeWriter{writers: writers}
}

// Write writes bytes to each of the writers in TeeWriter, returning the number
// of bytes written and an error, if any occurs.
func (t *TeeWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		n, err = w.Write(p)
		if err != nil {
			return n, err
		}
	}
	return len(p), nil
}
