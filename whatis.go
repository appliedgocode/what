// +build what whatis

package what

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

// Is pretty-prints data.
func Is(v interface{}) {
	spew.Fdump(log.Writer(), v)
}
