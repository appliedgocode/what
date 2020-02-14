// +build what whatpackage

package what

import (
	"log"
)

func Package() {
	log.Println("Package", pkgname(2))
}
