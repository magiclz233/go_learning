package main

import "fmt"

func main() {
	fmt.Println("Hello Golang!")
}

/*
$ go run hello-world.go
hello world

如果我们想将程序编译成二进制文件（Windows 平台是 .exe 可执行文件）， 可以通过 go build 来达到目的。
$ go build hello-world.go
$ ls
hello-world    hello-world.go

然后我们可以直接运行这个二进制文件。
$ ./hello-world
hello world
*/
