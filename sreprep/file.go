package sreprep

import (
	"fmt"
	"os"
	"bufio"
)

func PrintFile (path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if s.Err() != nil {
		return s.Err()
	}

	return nil
}