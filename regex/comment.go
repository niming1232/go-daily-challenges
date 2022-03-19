package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")
	in := bufio.NewReader(os.Stdin)
	s, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("ReadString", err)
		return
	}
	if isComment(strings.TrimSuffix(s, "\n")) {
		fmt.Println("This is a comment of C!")
		return
	}
	fmt.Println("This is NOT a comment of C.")
}

func isComment(s string) bool {
	legalRe := regexp.MustCompile(`^/\*.*\*/$`)
	illegalRe := regexp.MustCompile(`^/\*.*/\*.*\*/.*\*/$`)
	if legalRe.MatchString(s) && !illegalRe.MatchString(s) {
		return true
	}
	return false
}
