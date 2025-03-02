package main

import "fmt"

func main() {
	var i any
	// go <= 1.18
	// var i interface{}

	i = 7
	fmt.Println(i)

	// i = "hi"
	fmt.Println(i)

	// s := i.(string) // type assertion
	// fmt.Println("s:", s)

	n, ok := i.(int) // won't panic if check if used
	if ok {
		fmt.Println(n)
	} else {
		fmt.Println("could not convert to int")
	}

	switch i.(type) {
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("unknown type: %T\n", i)
	}
}
