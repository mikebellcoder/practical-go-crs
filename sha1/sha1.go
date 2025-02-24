package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	s, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("something wrong with sha1Sum: %s", err)
	}
	fmt.Println(s)
}

/*
if file names end with .gz

	$ cat http.log.gz | gunzip | sha1sum

else

	$ cat http.log | sha1sum
*/
func sha1Sum(fileName string) (string, error) {
	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var r io.Reader = file
	// unzip if necessary
	if strings.HasSuffix(fileName, ".gz") {
		// unzip
		f, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer f.Close()
		r = f
	}
	// create sha1sum
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), err
}
