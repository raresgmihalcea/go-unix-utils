package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, dir := range pathSplit {
		fullPath := filepath.Join(dir, file)
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}

		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}
	}
}
