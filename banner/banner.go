package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)

	s := "G☺"
	fmt.Println("len:", len(s))
	// code point = rune ~= unicode character
	for i, r := range s {
		// rune (int32)
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
		}
	}

	// byte (uint8)
	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)

	// format examples
	x, y := 1, "1"
	fmt.Printf("x=%#v, y=%#v\n", x, y)

	fmt.Printf("%20s!\n", s)

	for _, e := range []string{"g", "go", "gog", "gogo", "g☺g"} {
		fmt.Printf("%#v is a palindrome: %t\n", e, isPalindrome(e))
	}
}

func isPalindrome(s string) bool {
	/*
		// mine
		if len(s) == 1 {
			return true
		}

		rev := ""
		for _, r := range s {
			rev = fmt.Sprintf("%c%s", r, rev)
		}

		return s == rev
	*/

	// his
	rs := []rune(s)
	for i := range len(rs) / 2 {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}

	return true
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2 // BUG: len is in bytes
	// padding := (width - len(text)) / 2 // BUG: len is in bytes
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
