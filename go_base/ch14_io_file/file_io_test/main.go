package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func inputFile() {
	inputFile, inputError := os.Open("ch13_interface_more.5/file_io_test/input.txt")
	if inputError != nil {
		fmt.Println("上传文件出现错误", inputError)
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("%s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

func inputFileBuf() {
	inputFile, inputError := os.Open("ch13_interface_more.5/file_io_test/input.txt")
	if inputError != nil {
		fmt.Println("上传文件出现错误", inputError)
		return
	}
	defer inputFile.Close()

	buf := make([]byte, 1024)
	n, err := inputFile.Read(buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		panic(err.Error())
	}
	if n > 0 {
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	inputFile()
	// inputFileBuf()
	//outPutFile()
	//Fscanln()
}

func outPutFile() {
	inputFile := "ch13_interface_more.5/file_io_test/input.txt"
	outPutFile := "ch13_interface_more.5/file_io_test/inputCopy.txt"
	buf, err2 := ioutil.ReadFile(inputFile)
	if err2 != nil {
		_, _ = fmt.Fprintf(os.Stderr, "File error: %s\n", err2)
	}
	fmt.Printf("%s\n", string(buf))
	err3 := ioutil.WriteFile(outPutFile, buf, 0644)
	if err3 != nil {
		panic(err3.Error())
	}
}

func Fscanln() {
	inputFile, inputError := os.Open("ch13_interface_more.5/file_io_test/input.txt")
	if inputError != nil {
		panic(inputError.Error())
	}
	defer inputFile.Close()

	var c1, c2, c3, c4, c5 []string
	for {
		var v1, v2, v3, v4, v5 string
		_, err := fmt.Fscanln(inputFile, &v1, &v2, &v3, &v4, &v5)
		if err != nil {
			break
		}
		c1 = append(c1, v1)
		c2 = append(c2, v2)
		c3 = append(c3, v3)
		c4 = append(c4, v4)
		c5 = append(c5, v5)
	}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
}
