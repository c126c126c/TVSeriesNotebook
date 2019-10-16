package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("start...")

	cachePath := "../cache/storage"

	fmt.Println("Input content:")

	value := ""
	fmt.Scanln(&value)

	file, err := getFile(cachePath)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(value + "\r\n")

	file.Close()
}

func getFile(filePath string) (f *os.File, e error) {
	_, err := os.Stat(filePath)
	if err != nil {
		newFile, e := os.Create(filePath)
		return newFile, e

	} else {
		file, e := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		return file, e
	}
}
