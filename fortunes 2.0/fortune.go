package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
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

func seekAndRead(f *os.File, b []byte, offset int64) error {
	_, err := f.Seek(offset, 0)
	if err != nil {
		return err
	}
	_, err = f.Read(b) // Read moves file pointer to next one!!!
	if err != nil {
		return err
	}
	return nil
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
	n := fileInfo.Size()
	rand.Seed(time.Now().UnixNano())
	// Intn returns the half-open interval [0,n) which suites seek perfectly
	offset := rand.Int63n(n)
	b := make([]byte, 1, 1)
	err = seekAndRead(f, b, offset)
	if err != nil {
		return "", err
	}

	for string(b) != "$" && offset != 0 {
		offset -= 1
		err = seekAndRead(f, b, offset)
		if err != nil {
			return "", err
		}
	}

	// Rewind back to the beginning due to pointer move caused by f.Read()
	if offset == 0 {
		_, err = f.Seek(offset, 0)
		if err != nil {
			return "", err
		}
	}

	r := bufio.NewReader(f)
	fortune, err := r.ReadString('$')
	fortune = strings.TrimLeft(fortune, "\n")
	fortune = strings.TrimRight(fortune, "\n$")
	return fortune, nil

	// b := make([]byte, 1, 1)
	// for i := n - 1; i >= 0; i-- {
	// 	_, err = f.Seek(i, 0)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	_, err = f.Read(b)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	fmt.Printf("%v", string(b))
	// }
	// return "", nil
}
