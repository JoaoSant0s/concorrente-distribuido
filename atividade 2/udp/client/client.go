package main

import (
	"fmt"
	"net"
	"bufio"	
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	} else {				
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Digite uma mensagem: ")
			text, _ := reader.ReadString('\n')

			fmt.Fprintf(conn, text + "\n")			

			message, _ := bufio.NewReader(conn).ReadString('\n')

			fmt.Print("Mensagem do Servidor: " + message)							
		}						
	}
	conn.Close()	
}