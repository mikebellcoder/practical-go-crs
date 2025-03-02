package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	// mapDemo()
	w, err := mostCommon(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(w)
}

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(freqs)
}

// "Who's on first?" -> [Who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

// Q: What is the most common word (ignoring case) in sherlock.txt
// Word frequency
func mapDemo() {
	var stocks map[string]float64 // symbol -> price
	sym := "TTWO"
	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	stocks = map[string]float64{
		sym:    136.73,
		"AAPL": 172.35,
	}

	for k := range stocks { // keys only
		fmt.Println(k)
	}

	for k, v := range stocks { // keys and values
		fmt.Println(k, "->", v)
	}

}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}

	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // split line by words
		// add words to count map
		for _, word := range words {
			freqs[strings.ToLower(word)]++
		}

	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}
