# What: debug-level logging that vanishes from production code

## How to import the package

```go
import "appliedgo.net/what"
```
(Do not use the direct path to the repo.)

## What does what do

`what` is a set of simple and easy logging functions, suitable for tracing any kind of activities in your code. `what` can print the current function name, quickly `Printf`-format your data, and dumps data structures. 

And last not least, no `what` calls reach your production binary (unless you want it so). Debug-level logging is for developers only.


## Who need this? 

You definitely should give `what` a closer look if you -

* heartily agree to [Dave Cheney's article about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging), or
* want to keep your production code free from *any* log or trace output. (Think security!) 

## How does it work?

First of all, `what` is intended for debug-level logging *only*. So,

* Use `what` for tracing and debugging your code. ("Does my code do what I intended? Does this variable contain what I expect? Why does the loop not stop when the break condition *should* be fulfilled?...")
* Use `log` for user-facing log output. ("What was the app doing before it said, 'cannot connect to server'? Did that service already sync or is it still waiting for other services?...")

You have to explicitly enable `what` logging through build flags (see below).

### Available functions

```go
what.Happens("Foo: %s", bar) // log.Printf("Foo: %s\n", bar)
what.If(cond, "Foo: %s", bar) // only print if cond is true
what.Func() // Print out the fully qualified function name
what.Is(var) // dump the structure and contents of var 
```

Spread these calls across your code, especially in places you want to observe closer. 

Debug-level logging is useful alongside unit testing as well as using a debugger. It does not attempt to replace any of these concepts.

### Enabling and disabling

#### Enable all functions

Simply pass the `what` tag to `go build`, `go install`, `go test` etc:

```sh
go build -tags what
```

And now just lean back and see your code talking about what it does.

#### Enable specific functions

To reduce the noise, you can decide to compile only specific parts of `what`:

* `whathappens` only enables `what.Happens()` and `what.If()`.
* `whatfunc` only enables `what.Func()`.
* `whatis` only enables `what.Is()`.

All disabled functions get replaced by no-ops.

Example:

```sh
go build -tags whathappens
```

You can also choose a combination of the above, for example: `go build -tags whathappens,whatis`


#### Enable debug logging for specific packages only

Go's build tag mechanism cannot help here, so this is done through an environment variable called "WHAT".

To enable specific packages for debug logging, set `WHAT` to a package name, or a list of package names.



#### Disable what

Nothing easier than that! Without any of the above build tags, all funtions get replace by no-ops, ready for being optimized away entirely (if the compiler decides to do so).

* No log output 
* No bloated binary
* No security leak from chatty binaries.

## Non-features

* Uses only stdlib `log`, no custom logger configurable
* No custom variable dumper/pretty-printer. At the moment, `what` uses `github.com/davecgh/go-spew`. See Spew's docs about the syntax used for printing a variable.

## Restrictions

Although `go run` should recognize all build flags that `go build` recognizes (including `-tags`), it seems that `go run main.go -tags what` does not consider the tag. Use `go build -tags what && ./main` instead.
