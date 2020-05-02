/* Package what provides a quick and convenient way of adding development-only log statements to your code.

Instead of firing up the debugger and stepping through each line one-by-one, spread a few what calls across the code you want to inspect, then run the code and watch the output.

Log output from what must be enabled through build tags.
This ensures that your debug logging does not leak into production code and involuntarily exposes sensitive data to prying eyes. Use the log package for generating user-facing log output from production code.

Enable what by passing "what" as a build tag:

	go build -tags what

Enable only parts of what by passing the respective build tag: whathappens, whatis, whatfunc, or whatpackage. (Good for reducing noise, e.g by muting what.Func().)

Functions that are not enabled by a build tag become no-ops.

Enable what for particular packages only by setting the environment variable WHAT to a package name or a comma-separated list of package names:

	export WHAT=pkg1,pkg2

(Also good for reducing noise.)
*/
package what
