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

	pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println(err)
	}	

	defer pc.Close()	
	
	conn2, err2 := net.Dial("udp", "localhost:1313")

	if err2 != nil {
		fmt.Println(err2)
	}else{
		fmt.Println("Connection with server 2")
	}

	defer conn2.Close()

	fmt.Println("Establishing readers")
		
	start(pc, conn2)
}

func start(pc net.PacketConn, conn2 net.Conn){
	for {		
		buffer := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buffer)

		if err != nil{			
			fmt.Println(err)	
		}else{					
			s := string(buffer[:n])
			message := TrimString(s)					

			fmt.Fprintf(conn2, message + "\n")

			newMessage, _ := bufio.NewReader(conn2).ReadString('\n')

			extractedValue := TrimString(newMessage)

			_, err := strconv.Atoi(extractedValue)		

			if err != nil {
				fmt.Println("Retornando resultado")
				b, _ := strconv.ParseBool(extractedValue)
				extractedNumber := TrimString(message)			
				if b {
					pc.WriteTo([]byte("O número " + extractedNumber + " é primo\n"), addr)			
				}else{
					pc.WriteTo([]byte("O número " + extractedNumber + " não é primo\n"), addr)			
				}
			} else {
				fmt.Println("Retornando erro")
				pc.WriteTo([]byte("Erro: Digite um número Natural válido.\n"), addr)
			}
		}				
	}
}

func TrimString(value string) string{
	return strings.Trim(value, "\n\r ")	
}