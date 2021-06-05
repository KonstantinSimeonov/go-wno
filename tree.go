package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
)

func PrintDir(dir, indent string) {
	file, err := os.Open(dir)
	if err != nil {
		log.Fatalf("Failed %s", err)
	}
	defer file.Close()

	file_list, _ := file.Readdir(0)

	for _, f := range file_list {
		if f.IsDir() {
			fmt.Println(indent + f.Name() + "/")
			PrintDir(filepath.Join(dir, f.Name()), indent + "  ")
		} else {
			fmt.Println(indent + f.Name())
		}
	}
}

func main() {
	PrintDir(os.Args[len(os.Args) - 1], "")
}
