//go:build what || whathappens || whatif || whatfunc || whatis || whatpackage

// Based on https://github.com/spotlightpa/almanack/blob/master/pkg/almlog/colorize.go by Carlana Johnson.

package what

import (
	"bytes"
	"fmt"
	"time"
)

const (
	reset      = "\033[0m"
	bold       = "\033[1m"
	dim        = "\033[2m"
	standout   = "\033[3m"
	underscore = "\033[4m"
	blink      = "\033[5m"
	blinkmore  = "\033[6m"
	invert     = "\033[7m"
	hide       = "\033[8m"
	del        = "\033[9m"
	black      = "\033[30m"
	red        = "\033[31m"
	green      = "\033[32m"
	yellow     = "\033[33m"
	blue       = "\033[34m"
	magenta    = "\033[35m"
	cyan       = "\033[36m"
	white      = "\033[37m"
	purple     = magenta + bold
)

func colorize(level string, args ...any) string {
	// Build the prefix: time, level, func, msg...
	logArgs := []any{"time", time.Now().UTC().Format("2006-01-02 15:04:05"), "level", level, "func", funcname(3), "msg", args[0]}

	// ...then append the key/value pairs
	logArgs = append(logArgs, args[1:]...)

	var buf bytes.Buffer

	// Iterate over the logArgs and apply formatting
	for i := 0; i < len(logArgs); i += 2 {
		key := fmt.Sprint(logArgs[i])
		val := fmt.Sprint(logArgs[i+1])

		keyColor := cyan
		valColor := magenta
		withKey := true

		switch key {
		case "time":
			withKey = false
			valColor = dim
		case "level":
			withKey = false
			switch level {
			case "DEBUG":
				valColor = dim
			case "INFO":
				valColor = green
			case "WARN":
				valColor = yellow
			case "ERROR":
				valColor = red + bold
			}
		case "func":
			withKey = false
			valColor = blue
		case "msg":
			withKey = false
			valColor = white + underscore
		case "err":
			valColor = red + bold
		}

		// Coloring for value based on level

		if withKey {
			// Format as `key=value`
			buf.WriteString(keyColor)
			buf.WriteString(key)
			buf.WriteString(reset)
			buf.WriteString("=")
		}
		buf.WriteString(valColor)
		buf.WriteString(val)
		buf.WriteString(reset)
		buf.WriteString(" ")
	}

	return buf.String()
}
