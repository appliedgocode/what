//go:build what || whathappens || whatif || whatfunc || whatslog || whatis || whatpackage

package what

import (
	"os"
	"runtime"
	"strings"
)

var enabled map[string]bool

func isPackageEnabled() bool {
	if len(enabled) == 0 { // all packages enabled
		return true
	}
	pkg, prnt := pkgname(3) // 3 = skip isPackageEnabled and the top-level what.X function
	if enabled[pkg] {
		return true
	}
	if enabled[prnt+"/"+pkg] {
		return true
	}
	return false
}

// pkgname returns the package name and the direct parent in the package path
// of the skip'th caller
func pkgname(skip int) (string, string) {
	pc, _, _, _ := runtime.Caller(skip)
	fn := runtime.FuncForPC(pc).Name()
	// possible fn formats:
	// * /some/path/to/package.(Receiver).Func
	// * /some/path/to/package.Func.func1 (closure)
	// * /some/path/to/package.Func
	// * pathto/package.Func
	// * package.Func (for package main, and if someone hacks
	//   the stdlib and adds `what` calls there.)

	startName, startParent, endName := 0, 0, 0

	// Search the last /, it is the beginning of the package name
	for i := len(fn) - 1; i >= 0; i-- {
		if fn[i] == '/' {
			startName = i + 1
			break
		}
	}
	// post-loop assert: startName is either 0 (for package main) or i+1

	// Search the first dot after the last /, it is the end of the package name
	for i := startName; i < len(fn); i++ {
		if fn[i] == '.' {
			endName = i
			break
		}
	}

	// no leading / found means we are probably in package main.
	if startName == 0 {
		return fn[0:endName], ""
	}

	// Search the second-last /, it is the beginning of the parent's name
	endParent := startName - 1
	for i := endParent - 1; i >= 0; i-- {
		if fn[i] == '/' {
			startParent = i + 1
			break
		}
	}
	// startParent is 0 in case of pathto/package.Func

	return fn[startName:endName], fn[startParent:endParent]
}

func getenvWhat() {
	packages := strings.Split(os.Getenv("WHAT"), ",")
	enabled = map[string]bool{}
	for _, p := range packages {
		if p != "" {
			enabled[p] = true
		}
	}
}

func init() {
	getenvWhat()
}
