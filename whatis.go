//go:build what || whatis

package what

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

// Is pretty-prints data.
func Is(v any) {
	spew.Fdump(log.Writer(), v)
}

func init() {
	// Default indent of one space is too narrow for my taste
	spew.Config.Indent = "  "
}
