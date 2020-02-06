// +build what whathappens whatfunc

package what

import "runtime"

// funcinfo returns func name, line, and file name of a
// caller, after skipping skip callers.
func funcinfo(skip int) (name, file string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		name = runtime.FuncForPC(pc).Name()
	}
	return name, file, line
}

// funcname returns the name of a
// caller, after skipping skip callers.
func funcname(skip int) string {
	name, _, _ := funcinfo(skip + 1)
	return name
}
