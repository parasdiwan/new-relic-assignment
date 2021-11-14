package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

const maxNumberOfConnections = 1
const portNumber = 4000

var activeConnections = 0

func StartListener() {
	address := fmt.Sprintf("0.0.0.0:%d", portNumber)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for activeConnections <= maxNumberOfConnections {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// If you want, you can increment a counter here and inject to handleClientRequest below as client identifier
		go handleClientRequest(connection)
		activeConnections++
	}
}

func handleClientRequest(con net.Conn) {
	clientReader := bufio.NewReader(con)

	for {
		log.Println("")
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
			activeConnections--
		} else {
			log.Println("invalid response from the client, closing connection")
			return
		}
	}
	activeConnections--
	defer con.Close()
}

func isValidResponse(request string) bool {
	_, err := strconv.Atoi(request)
	return err == nil
}
