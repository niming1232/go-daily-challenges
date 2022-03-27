package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	path := "./data.txt"
	fortune, err := tellFortune(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fortune)
}

func tellFortune(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().Unix())
	r := rand.Intn(int(fileInfo.Size()))
	_, err = f.Seek(int64(r), 0)
	if err != nil {
		return "", err
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	var c int
	var fortune string
	for s.Scan() {
		char := s.Text()
		if char != "$" && c == 1 {
			fortune += char
			continue
		}
		if char == "$" {
			c++
		}
		if c == 2 {
			break
		}
	}
	return fortune, nil
}
