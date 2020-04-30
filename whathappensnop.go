// +build !what,!whathappens

package what

// Happens logs the current function name and whatever message is passed in.
func Happens(fmt string, args ...interface{}) {}

// If works like Happens but only if yes is true.
func If(yes bool, fmt string, args ...interface{}) {}
