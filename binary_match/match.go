package main

import (
	"fmt"
	"strings"
)

func main() {
	var count int
	template := "??"
	for _, c := range template {
		if c == '?' {
			count++
			continue
		}
		if c != '0' && c != '1' {
			fmt.Printf("illegal character %v\n", string(c))
			return
		}
	}
	patterns := make([]string, 0, count)
	// fmt.Println("call recursive to resolve", template)
	findAllPatterns(template, &patterns)
	fmt.Println(patterns)
}

func findAllPatterns(template string, patterns *[]string) {
	var hasWildcard bool
	for _, c := range template {
		if c == '?' {
			hasWildcard = true
			s0 := strings.Replace(template, string(c), "0", 1)
			s1 := strings.Replace(template, string(c), "1", 1)
			// fmt.Println("call recursive to resolve", s0)
			findAllPatterns(s0, patterns)
			// fmt.Println("call recursive to resolve", s1)
			findAllPatterns(s1, patterns)
			break
		}
	}
	if !hasWildcard {
		// fmt.Println(template)
		*patterns = append(*patterns, template)
	}
	return
}
