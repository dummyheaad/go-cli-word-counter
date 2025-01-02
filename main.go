package main

import (
	"bufio"
	"flag"
	"fmt"
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

	res := getCount(scanner)

	return res
}

func countFromFile(filenames []string, countLines bool, countBytes bool) (map[string]int, error) {

	res := make(map[string]int)

	for _, filename := range filenames {
		// Read the file
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		if countBytes {
			scanner.Split(bufio.ScanBytes)
		} else if countLines {
			scanner.Split(bufio.ScanLines)
		} else {
			scanner.Split(bufio.ScanWords)
		}

		result := getCount(scanner)

		res[filename] = result
	}

	return res, nil
}

func getCount(scanner *bufio.Scanner) int {
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
	file := flag.String("file", "", "Read from file")
	// Parsing the flags provided by the user
	flag.Parse()

	filenames := flag.Args()
	if *file != "" {
		filenames = append([]string{*file}, filenames...)
	}

	if *file != "" {
		// Read input from file
		res, err := countFromFile(filenames, *lines, *bytes)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// TODO: process the res
		switch {
		case *lines:
			fmt.Println("Line counting result for each file:")
		case *bytes:
			fmt.Println("Byte counting result for each file:")
		default:
			fmt.Println("Word counting result for each file:")
		}
		for fn, c := range res {
			fmt.Printf("filename: %s, Count: %d\n", fn, c)
		}
	} else {
		// Calling the count function to count the number of words
		// received from the Standard Input and printing it out
		fmt.Println(count(os.Stdin, *lines, *bytes))
	}
}
