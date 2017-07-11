

package main

import "fmt"

func main() {
    fmt.Println("hello world")
}

func valuesTest() {

// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.

    // Strings, which can be added together with `+`.
    fmt.Println("go" + "lang")

    // Integers and floats.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Booleans, with boolean operators as you'd expect.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}

func variablesTest() {
// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

    // `var` declares 1 or more variables.
    var a string = "initial"
    fmt.Println(a)

    // You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding
    // initialization are _zero-valued_. For example, the
    // zero value for an `int` is `0`.
    var e int
    fmt.Println(e)

    // The `:=` syntax is shorthand for declaring and
    // initializing a variable, e.g. for
    // `var f string = "short"` in this case.
    f := "short"
    fmt.Println(f)
}