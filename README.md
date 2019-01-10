## windows

[![GoDoc](https://godoc.org/github.com/mxplusb/windows?status.svg)](https://godoc.org/github.com/mxplusb/windows) [![Build status](https://ci.appveyor.com/api/projects/status/c86d004j0ia1c1fr?svg=true)](https://ci.appveyor.com/project/mxplusb/windows)

This package makes a best-effort attempt at an idiomatic Go interface to making syscalls for Windows.

### Contributing

To add syscalls:

* `go build -v mksyscall_windows.go`
* `cd <top-level-api-folder>`
  * New APIs go under their respective DLL. For example, `GetProcessInformation` belongs to `psapi.dll`, so it goes under `psapi`.
* Add the unexported syscall in this format: `//sys funcName(named type) (namedReturn type)`
  * Hint: `mksyscall_windows.go` contains instructions at the top of the file if you get stuck. File an issue if you really get stuck, I will help you through it.
  * If the syscall has pointer returns to a type, make sure to add the type with appropriate comments. Docs are a must!
* Run `mksyscall_windows.exe -output="z<dest_api>_windows.go" <target_api>_template_windows.go` in the API folder to generate the syscall.
* Write a test for the syscall.
  * If the syscall has pointer returns to a type, make sure to create a method receiver and deep copy it to provide an idiomatic interface.
    * And a test for that method receiver.
  * Write an example if possible!
  * Tests are designed to help you understand what the API feels like, if it feels clunky, then write a wrapper for it.
* Submit a PR! This package can't grow without your help, let's make it awesome!

### Why `github.com/mxplusb/windows` instead of `golang.org/x/sys/windows`?

I have put a few syscalls into the `golang.org/x/sys` library, and it's neither well documented nor easy. The Go Authors feel it's best to leave documentation out, which forces devs to search for them online, and really raises the bar for entry. Also, adding new syscalls is time consuming and not guaranteed due to pre-existing bloat. They have stated it's impossible to get all Windows APIs into `golang.org/x/sys/windows` and they have no interest in really optimising the experience. While I understand their logic, I was thoroughly turned off when I was looking for Windows syscalls, so I felt I can provide a better dev experience with a more Windows developer friendly interface with [blackjack and hookers](https://www.youtube.com/watch?v=5l3ipKcnYlQ). This is not to say the Go Authors' approach is bad, it's just a difference of opinion.

I felt it was more important to have:

* Idiomatic interfaces.
  * Windows types are available. You can use a `DWORD` in a syscall, because that's what you would do if this was Windows-based C++.
  * No pointer return interfaces when possible. Example: `SomeSyscallFunc(handle, &someReturn)`
    * While there are certain packages that do this, such as `encoding/json`, on the whole that's non-standard.
* Method receivers for a type. It should be obvious which interface you're using, and it should feel like Go, not like C++.
* Low bar to entry.
  * You shouldn't have to have years of experience with Windows to be productive.
  * It should be exceedingly obvious what the syscall does.
  * There should be examples when possible! You shouldn't have to go to Stack Overflow just to understand an interface.

There are some tradeoffs when consuming this package instead of `golang.org/x/sys/windows`.

* To provide a simple interface on method receivers, we do a deep copy from the unexported syscall.
  * While not slow, there is a lot of reflection that takes place.
  * Deep copies are `O(n^2)` complexity, so on large structs, there is a time tradeoff for a simple interface.
* Sometimes there is a little indirection under the hood to make it more user friendly.

### Getting Help

If you don't understand how something works, please file an issue asking for help so you can get some! Or feel free to email me from my Github email, I'll help you understand how it works as much as I can.
