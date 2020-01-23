// +build what whatfunc

package what

import (
	"log"
)

// Func logs the current function name, line number, and file name.
// Useful for tracing function calls
func Func() {
	name, file, line := funcinfo(2)
	log.Printf("Func %s in line %d of file %s\n", name, line, file)
}
