package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World")

	// Values Session 2 ----------------------
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(false && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Variable Session 3 ----------------------
    var a = "initial"
    fmt.Println(a)
    var b, c int = 1, 2
    fmt.Println(b, c)
	var d = true
    fmt.Println(d)
	var e int
    fmt.Println(e)
	f := "apple"
    fmt.Println(f)

	// Constants ----------------------
	const s string = "constant"
	const n = 500000000
	const t = 3e20 / n
    fmt.Println(d)
	fmt.Println(int64(t))
	fmt.Println(math.Sin(n))
}

// Learning-Golang
// Learning Golang by building and testing command.