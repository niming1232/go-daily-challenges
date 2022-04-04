package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func walk(path string, lv int) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			if err != nil {
				return err
			}
			fmt.Println(strings.Repeat("-", lv) + entry.Name() + "/")
			walk(path+"/"+entry.Name(), lv+1)
			continue
		}
		fmt.Println(strings.Repeat("-", lv) + entry.Name())
	}
	return nil
}

func main() {
	var lv int
	path := "/Users/xinhe/go-daily-challenges/dir_walk"
	fmt.Println("using my own walk function which follows symbolic links")
	err := walk(path, lv)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("using filepath.Walk which follow symbolic links")
	err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("visited file or dir: %q\n", info.Name())
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", path, err)
		return
	}

	fmt.Println("using filepath.WalkDir which does not follow symbolic links (but it still shows the symbolic link?)")
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("visited file or dir: %q\n", d.Name())
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", path, err)
		return
	}
}
