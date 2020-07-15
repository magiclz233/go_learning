package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

var inputReader *bufio.Reader
var in string
var err error

func main() {
	//fmt.Println("输入你的名字:")
	//_, _ = fmt.Scanln(&firstName, &lastName)
	//fmt.Printf("Hi %s %s!\n", firstName, lastName)
	//_, _ = fmt.Sscanf(input, format, &f, &i, &s)
	//fmt.Println("we read:", f, i, s)

	//inputReader = bufio.NewReader(os.Stdin)
	//fmt.Println("输入: ")
	//in, err = inputReader.ReadString('\n')
	//if err == nil {
	//	fmt.Printf("The input was: %s\n",in)
	//}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	in1, err1 := inputReader.ReadString('\n')
	if err1 != nil {
		fmt.Println("error!")
		return
	}

	fmt.Printf("Your name is: %s\n", in1)
	switch in1 {
	case "Magic\n":
		fmt.Println("Welcome Magic!")
	case "Chris\n":
		fmt.Println("Welcome Chris!")
	case "Mike\n":
		fmt.Println("Welcome Mike!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}
}
