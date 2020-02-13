// +build what whathappens whatif whatfunc

package what

import (
	"log"
	"os"
	"runtime"
	"strings"
)

var enabled map[string]bool

func isPackageEnabled() bool {
	if enabled[pkgname(3)] { // skip isPackageEnabled and the top-level what.X function
		return true
	}
	return false
}

func pkgname(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	fn := runtime.FuncForPC(pc).Name()
	log.Println("fn:", fn)
	// fn looks like /some/path/to/package.(Receiver).Func
	// or like /some/path/to/package.Func.func1 for anonymous funcs inside Func
	// or just /some/path/to/package.Func
	// so the package name is always the string between
	// the last "/" and the first "."

	start, end := 0, 0
	for i := len(fn) - 1; i >= 0; i-- {
		if fn[i] == '/' {
			start = i + 1
			break
		}
	}

	if start == 0 {
		return ""
	}

	for i := start; i < len(fn); i++ {
		if fn[i] == '.' {
			end = i
			break
		}
	}

	return fn[start:end]
}

func init() {
	packages := strings.Split(os.Getenv("WHAT"), ",")
	enabled = map[string]bool{}
	for _, p := range packages {
		enabled[p] = true
	}
}
