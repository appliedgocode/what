// +build what whatpackage

package what

import (
	"log"
)

// Package prints the code's package name, including the
// parent path entry if any.
func Package() {
	pkg, prnt := pkgname(2)
	if len(prnt) > 0 {
		log.Printf("Package %s/%s", prnt, pkg)
	} else {
		log.Printf("Package %s", pkg)
	}
}
