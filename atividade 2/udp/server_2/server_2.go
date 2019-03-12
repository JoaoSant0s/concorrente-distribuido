package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := ln.Accept()

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()		
		
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		newMessage := strings.ToUpper(message)

		conn.Write([]byte(newMessage + "\n"))						
	}
}