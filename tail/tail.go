package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var path = flag.String("path", "./temp.txt", "The path of the file.")
var n = flag.Int("n", 3, "The number of lines.")

func main() {
	flag.Parse()

	lines, err := tail(*path, *n)
	if err != nil {
		log.Fatal(err)
	}
	for i := range lines {
		fmt.Println(lines[i])
	}
}

func tail(path string, n int) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	if n >= len(lines) {
		return lines, nil
	}
	return lines[len(lines)-n:], nil
}
