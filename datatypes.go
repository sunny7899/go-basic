package main

// Program for GO data types

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("\nData Types\n" + strings.Repeat("-", 10))

	var b bool = true // Boolean

	var u uint8 // 8 bits / 1 byte

	var i int = 5        // Integer
	var f float32 = 3.14 // Floating point number
	var s string = "Hi!" // String

	fmt.Printf(`
	Type of b: %v, value: %v
	Type of i: %v, value: %v
	Type of u: %v, value: %v
	Type of f: %v, value: %v
	Type of s: %v, value: %v
	`,
		reflect.TypeOf(b), b,
		reflect.TypeOf(i), i,
		reflect.TypeOf(u), u,
		reflect.TypeOf(f), f,
		reflect.TypeOf(s), s)

}
