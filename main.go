package main

import(
	"fmt"
	"flag"
)

func main() {
	fileNameFlag := flag.String("file", "problems.csv", "a csv file that has the format of 'question, answer'")
	flag.Parse()
	fmt.Println("flag input is", *fileNameFlag)
}