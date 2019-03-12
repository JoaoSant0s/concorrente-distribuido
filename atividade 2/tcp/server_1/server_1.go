package main

import (
	"fmt"
	"net"
	"bufio"	
	"strings"
	"strconv"
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
	
	conn2, err2 := net.Dial("tcp", "localhost:1313")

	if err2 != nil {
		fmt.Println(err2)
	}

	defer conn2.Close()
		
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')		
		
		fmt.Fprintf(conn2, message + "\n")

		newMessage, _ := bufio.NewReader(conn2).ReadString('\n')

		extractedValue := extractNumberString(newMessage)

		_, err := strconv.Atoi(extractedValue)

		if err != nil {
			b, _ := strconv.ParseBool(extractedValue)
			extractedNumber := extractNumberString(message)
			conn.Write([]byte("Checando... \n"))	
			if b {
				conn.Write([]byte("O número " + extractedNumber + " é primo\n"))			
			}else{
				conn.Write([]byte("O número " + extractedNumber + " não é primo\n"))			
			}
		} else {	
			conn.Write([]byte("Erro: Digite um número Natural válido.\n"))
		}				
	}
}

func extractNumberString(value string) string{
	newValue := strings.ReplaceAll(value, "\r", "")
	newValue = strings.ReplaceAll(newValue, "\n", "")
	return newValue
}