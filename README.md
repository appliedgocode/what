# What: debug-level logging that vanishes from production code

## What does what do

`what` is a set of simple and easy logging functions, suitable for tracing any kind of activities in your code. `what` can print the current function name, quickly `Printf`-format your data, and pretty-print data structures. 

## Who need this? 

You definitely should give `what` a closer look if you -

* heartily agree to [Dave Cheney's article about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging), or
* want to keep your production code free from *any* log or trace output. (Think security!) 

## How does it work?

### Available functions

```go
what.Happens("Foo: %s", bar) // log.Printf("Foo: %s\n", bar)
what.If(cond, "Foo: %s", bar) // only print if cond is true
what.Func() // Print out the fully qualified function name
what.Is(var) // Pretty-print var 
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


#### Disable what

Nothing easier than that! Without any of the above build tags, all funtions get replace by no-ops, ready for being optimized away entirely (if the compiler decides to do so).

* No log output 
* No bloated binary
* No security leak from chatty binaries.

## Non-features

* Uses only stdlib `log`, no custom logger configurable
* No custom pretty-printer