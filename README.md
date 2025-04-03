# GO WITH TEST

This repository works as a learning guide where we make every step with a TDD approach.

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

