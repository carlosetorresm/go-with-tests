# GO WITH TEST

This repository works as a learning guide where we make every step with a TDD approach.

**GO DOCUMENTATION**

We can install the pkgsite in order to view the documentation locally with the follow command:

`go install golang.org/x/pkgsite/cmd/pkgsite@latest`
`pkgsite -open .`

This command will download the source files from that repository and build them into an executable binary, for default in the *$HOME/go/bin* for Linux and MacOS.

We will be able to acces the documentation navigating to.

`http://localhost:8080/`


**PLACEHOLDERS**

When printing, or logging we can make use of the *FMT Package* placeholders

`fmt.Println("this value is: %b", 20)`

`%q` -> wraps value in double quotes.
`%b` -> representation of a base 2 *integer*.
`%T` -> representation of the type of the value.
`%d` -> representation of an *integer* value.


**TEST FUNCTIONS**

Every test are supposed to be part of the package it's testing, also it must have the **_test.go** suffix.

This functions execute the **test_xxx** functions where **xxx** should be the name of the function to test.

It also execute testable Examples, this are functions where we add a commented output.

***To execute the tests***
`go test`

***To execute the tests with Info about each one***
`go test -v`


**BENCHMARK FUNCTIONS**

Benchmarks are run sequentially and can be inaccurate because of Go optimizations.

An option to avoid that would be to record the result
in order to prevent the compiler from eliminating the function call.

***To execute the benchmarks***
`go test -bench=.`

***To execute the benchmarks with memory allocation info***
`go test -bench=. -benchmem`

**COVERAGE**

The coverage tool can help identify areas of the code not covered by tests.

***To execute a coverage  analisis***
`go test -cover`