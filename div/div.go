package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(div(1, 0))
	fmt.Println(safeDiv(1, 0))
}

func safeDiv(a, b int) (q int, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
