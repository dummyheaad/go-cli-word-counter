package main

import (
	"bufio"
	"fmt"
	"flag"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if countLines {
		scanner.Split(bufio.ScanLines)
	} else {
		scanner.Split(bufio.ScanWords)
	}
	

	// Defining a counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
