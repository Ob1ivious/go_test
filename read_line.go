package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	filePath := flag.String("file", "access.log", "read file")
	flag.Parse()

	file, err := os.Open(*filePath)

	if err != nil {
		fmt.Println("err:", err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	line := bufio.NewScanner(file)

	for line.Scan() {
		slice := strings.Split(line.Text(), " ")
		fmt.Printf("%q\n", slice[6])
		fmt.Printf("%q\n", slice[len(slice)-1])

		fmt.Println("-------------")
	}

	err = line.Err()

	if err != nil {
		fmt.Println(err)
	}

}
