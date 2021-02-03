package main

import "fmt"

func main() {
	fmt.Println(c())
	a()
	b()
}

// A deferred function's arguments are evaluated when the defer statement is evaluated.
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func b() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

// Deferred functions may read and assign to the returning function's `named` return values.
// will return 2 because of after return statement executed it checked for return variables values
func c() (i int) {
	defer func() { i++ }()
	return 1
}
