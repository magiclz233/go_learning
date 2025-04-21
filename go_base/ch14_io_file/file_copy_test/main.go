package main

import (
	"io"
	"os"
)

// copy文件

func main() {

	//CopyFile()
	// 从标准输入读取
	// data, err = ReadFrom(os.Stdin, 11)
	// 从文件读取
	// data1, err = ReadFrom(os.File, 11)
	// 从字符串读取
// data, err = ReadFrom(strings.NewReader("from string"), 12)
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

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n>0 {
		return p[:n], nil
	}
	return p, err
}


