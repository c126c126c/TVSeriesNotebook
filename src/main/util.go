package main

import "os"

func GetFile(filePath string) (f *os.File, e error) {
	_, err := os.Stat(filePath)
	if err != nil {
		newFile, e := os.Create(filePath)
		return newFile, e

	} else {
		file, e := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		return file, e
	}
}
