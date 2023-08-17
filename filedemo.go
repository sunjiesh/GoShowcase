package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var fileLines []string
	//fileLines = make([]string, 0)
	file, err := os.Open("/tmp/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines,scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	// for _, fileLine := range fileLines {
	// 	fmt.Println(fileLine)
	// }
	fmt.Println(fileLines)
	fmt.Println(fileLines[:3])

}
