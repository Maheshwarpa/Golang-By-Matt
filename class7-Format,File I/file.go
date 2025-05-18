package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile() {
	fmt.Println("Going to read the file format.txt")
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Println(os.Stderr, err)
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()

			fmt.Println(s)
		}

		file.Close()
	}
}
