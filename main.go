package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	dir := flag.Args()

	fmt.Println(dir)
	err := filepath.Walk(dir[0],
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".jpg" {
				fmt.Println(path)
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
}
