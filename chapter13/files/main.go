package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fileW, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer fileW.Close()

	fileW.WriteString("test")

	fileR, err := os.Open("test.txt")
	if err != nil {
		// handle the error
		return
	}
	defer fileR.Close()

	// get the file size
	stat, err := fileR.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = fileR.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
