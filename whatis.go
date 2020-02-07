// +build what whatis

package what

import (
	"github.com/k0kubun/pp"
)

// Is pretty-prints data.
func Is(v interface{}) {
	pp.Print(v)
}
