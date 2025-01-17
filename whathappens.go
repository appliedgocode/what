//go:build what || whathappens

package what

import (
	"log"
)

// Happens logs the current function name and whatever message is passed in. The fmt string can use any formatting verbs as defined by the `fmt` package.
// Exception: If the fmt string is either of DEBUG, INFO, WARN, or ERROR, Happens switches to structured logging with colorization. In this case, the args must follow `log/slog` convention, that is: "message", "key1", value1, "key2", value2, etc.
// Happens prefixes the output with the current time, the log level, and the function name.
func Happens(fmt string, args ...any) {
	if isPackageEnabled() {
		// First, attempt to read a structured log message and colorize the output.
		switch fmt {
		case "DEBUG", "INFO", "WARN", "ERROR":
			log.SetFlags(0)
			log.Print(colorize(fmt, args...))
		default:
			log.Printf(funcname(2)+": "+fmt+"\n", args...)
		}
	}
}

// If works like Happens but only if yes is true.
func If(yes bool, fmt string, args ...any) {
	if yes && isPackageEnabled() {
		log.Printf(funcname(2)+": "+fmt+"\n", args...)
	}
}
