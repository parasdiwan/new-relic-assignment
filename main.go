package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("hello world!")
	listener, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		fmt.Println("there has been a problem")
		fmt.Println(err)
	}
	connection, err := listener.Accept()
	if err != nil {
		fmt.Println("there has been a problemo accepting connections!")
	}
	isInputValid := true
	for isInputValid {
		var inputString, _ = bufio.NewReader(connection).ReadString('\n')
		fmt.Println(inputString)
		// TODO trim input
		isInputValid = checkInputValidity(inputString)
	}
}

func checkInputValidity(inputString string) bool {
	return true
}
