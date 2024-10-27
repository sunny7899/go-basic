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

	// below variable based on values data type is inferred
	i1 := 10
	f1 := 22.22
	s1 := "Hello there"

	// below variable based on values data type is inferred
	var i2 = 100
	var f2 = 0.0012211221
	var s2 = "Guess "

	fmt.Printf(`
	Type of b:  %v, value: %v
	Type of i:  %v, value: %v
	Type of u:  %v, value: %v
	Type of f:  %v, value: %v
	Type of s:  %v, value: %v
	Type of i1: %v, value: %v
	Type of f1: %v, value: %v
	Type of s1: %v, value: %v
	Type of i2: %v, value: %v
	Type of f2: %v, value: %v
	Type of s2: %v, value: %v
	
	`,
		reflect.TypeOf(b), b,
		reflect.TypeOf(i), i,
		reflect.TypeOf(u), u,
		reflect.TypeOf(f), f,
		reflect.TypeOf(s), s,
		reflect.TypeOf(i1), i1,
		reflect.TypeOf(f1), f1,
		reflect.TypeOf(s1), s1,
		reflect.TypeOf(i2), i2,
		reflect.TypeOf(f2), f2,
		reflect.TypeOf(s2), s2,
	)

}
