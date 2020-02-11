package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("\nfindIDHash v1          Olivier Laflamme")
	fmt.Println("\n=========================================")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE: -file <<Filename>> [default: hash.txt] \n\nPut the unknown hash into a file. Only 1 hash at a time and remove domain or usernames in your raw hash.\nThis tool compares hashes by length so usernames with varying length sizes will not match the examples in hashcatexamples.csv. hashcatexamples.csv file is required\n\nExample: go run findIDHash.go -file hash.txt\n\n")
	}
	rows := exampleHashes("hashcathash.csv")
	rows = calculate(rows)
}

//Read Hashcat examples as CSV files (https://hashcat.net/wiki/doku.php?id=example_hashes)
func exampleHashes(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '|'

	// read the entire file. (if its small enough sometimes it fails and i dont understand)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}

//Processing
func calculate(rows [][]string) [][]string {
	wordFile := flag.String("file", "hash.txt", "enter file name")
	flag.Parse()
	txtfile := *wordFile

	// open the file filepath
	fileHandle, _ := os.Open(txtfile)
	defer fileHandle.Close()

	fs, err := ioutil.ReadFile(txtfile)
	if err != nil {
		fmt.Print(err)
	}

	str := string(fs)
	//trim whitespace
	trimmed := strings.TrimSpace(str)

	for i := range rows {
		// column 0 can be removed
		hashExample := rows[i][2]
		hashName := rows[i][1]
		hashType := rows[i][0]

		if len(hashExample) == len(trimmed) {
			fmt.Println(hashName + " [-m " + hashType + "]")
		}
	}
	return rows
}
