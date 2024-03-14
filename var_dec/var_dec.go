package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

// vardec\var_dec.go

var c, python, java bool

func add(x int, y int) int {
	return x + y
}

func multiply(x, y float32) float32 {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// return x, y
	return
}

var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var (
	k int
	l bool
	m string
	n float32
)

func main() {

	fmt.Println("My favorite number is ", rand.Intn(19), ".")
	fmt.Println(add(2, 5))
	fmt.Println(multiply(2, 5))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var age int
	age = 19
	fmt.Println(age)

	fmt.Println(split(age))

	fmt.Println(i, c, python, java)

	var d, e, f = true, false, "no!"

	k := 3
	fmt.Println(i, j, k, d, e, f)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("\t", k, l, m, n)
	fmt.Printf("\t--- %v --- %v --- %v --- %v | %q\n", k, l, m, n)

	var z uint = uint(n)
	fmt.Println(z)
}
