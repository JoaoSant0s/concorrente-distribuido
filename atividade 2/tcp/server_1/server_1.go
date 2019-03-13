package main

import (
	"fmt"
	"net"
	"bufio"	
	"strings"
	"strconv"
)

func main() {	
	fmt.Println("Initing server 1")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := ln.Accept()

	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Connection with client")
	}

	defer conn.Close()	
	
	conn2, err2 := net.Dial("tcp", "localhost:1313")

	if err2 != nil {
		fmt.Println(err2)
	}else{
		fmt.Println("Connection with server 2")
	}

	defer conn2.Close()

	fmt.Println("Establishing readers")
		
	start(conn, conn2)
}

func start(conn net.Conn, conn2 net.Conn){
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')		
		
		fmt.Fprintf(conn2, message + "\n")

		newMessage, _ := bufio.NewReader(conn2).ReadString('\n')

		extractedValue := TrimString(newMessage)

		_, err := strconv.Atoi(extractedValue)		

		if err != nil {
			fmt.Println("Retornando resultado")
			b, _ := strconv.ParseBool(extractedValue)
			extractedNumber := TrimString(message)			
			if b {
				conn.Write([]byte("O número " + extractedNumber + " é primo\n"))			
			}else{
				conn.Write([]byte("O número " + extractedNumber + " não é primo\n"))			
			}
		} else {
			fmt.Println("Retornando erro")
			conn.Write([]byte("Erro: Digite um número Natural válido.\n"))
		}				
	}
}

func TrimString(value string) string{
	return strings.Trim(value, "\n\r ")		
}