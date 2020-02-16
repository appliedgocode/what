// +build what whatpackage

package what

import (
	"log"
)

func Package() {
	pkg, prnt := pkgname(2)
	if len(pkg) > 0 {
		log.Printf("Package %s/%s", prnt, pkg)
	} else {
		log.Printf("Package %s", pkg)
	}
}
