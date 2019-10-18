package main

import "fmt"

//将数据更新
func (cli *CLI) check(data string, where string) {
	cachePath := "../../cache/storage"
	file, err := GetFile(cachePath)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(data + ";" + where + "\r\n")
	file.Close()
}
