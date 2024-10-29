// program to read file

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]

	readlFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error openning file: %v\n", err)
		os.Exit(1)
	}

	fileInfo, _ := os.Stat(filename)
	fileName, fileSize := fileInfo.Name(), fileInfo.Size()
	fmt.Printf("\nReading..\n\tFileName : %v\n\tSize     : %v\n\n", fileName, fileSize)

	fileScanner := bufio.NewScanner(readlFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	readlFile.Close()
}
