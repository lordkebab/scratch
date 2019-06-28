package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {

	fname := "/export/home/new_cadcam_file"

	fh, err := os.Open(fname)
	if err != nil {
		fmt.Printf("%s does not exist\n", fname)
	}
	
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		file := scanner.Text()

		_, err := os.Stat(file)
		if err != nil {
			fmt.Printf("%s does not exist\n", file)
		}
	}
}
