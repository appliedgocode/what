//go:build what || whatis

package what

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

// DebugString() is a pseudo-standard for printing debug info.
// HT to DoltHub: https://www.dolthub.com/blog/2025-01-03-gos-debug-string-pseudo-standard/
type DebugStringer interface {
	DebugString() string
}

// Is pretty-prints data.
func Is(v any) {
	// if v is a DebugStringer, have it print its debug info.
	if ds, ok := v.(DebugStringer); ok {
		spew.Fprint(log.Writer(), ds.DebugString())
		return
	}
	// any non-DebugStringer -> dump the data.
	spew.Fdump(log.Writer(), v)

}

func init() {
	// Default indent of one space is too narrow for my taste
	spew.Config.Indent = "  "
}
