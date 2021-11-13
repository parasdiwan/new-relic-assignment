package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

const maxNumberOfConnections = 5

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:4000")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// If you want, you can increment a counter here and inject to handleClientRequest below as client identifier
		go handleClientRequest(connection)
	}
}

func handleClientRequest(con net.Conn) {
	clientReader := bufio.NewReader(con)

	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n')

		if err != nil {
			log.Println("error reading client request: %v", err)
			return
		}
		clientRequest = strings.TrimSpace(clientRequest)
		if clientRequest == "terminate" {
			log.Println("client requested server to close the connection so closing")
			return
		}
		if isValidResponse(clientRequest) {
			log.Println(clientRequest)
		} else {
			log.Println("invalid response from the client, closing connection")
			return
		}
	}
}

func isValidResponse(request string) bool {
	_, err := strconv.Atoi(request)
	return err == nil
}
