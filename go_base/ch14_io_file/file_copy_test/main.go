package main

import (
	"io"
	"os"
)

// copy文件

func main() {

	//CopyFile()
}

func CopyFile(oldName, newName string) (written int64, err error) {
	oldFile, err1 := os.Open(oldName)
	if err1 != nil {
		return
	}
	defer oldFile.Close()

	newFile, err2 := os.OpenFile(newName, os.O_WRONLY|os.O_CREATE, 0644)
	if err2 != nil {
		return
	}
	defer newFile.Close()

	return io.Copy(newFile, oldFile)
}
