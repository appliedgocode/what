// +build what whatpackage

package what

func Package() {
	log.Println("Package", pkgname(2))
}
