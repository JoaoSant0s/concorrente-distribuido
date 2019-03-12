package main

import (
	"fmt"
	"net"
	"bufio"	
	"os"
)

func main() {
	fmt.Println("Verificador de números primos")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	} else {				
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Digite um número: ")
			text, _ := reader.ReadString('\n')

			fmt.Fprintf(conn, text + "\n")			

			message, _ := bufio.NewReader(conn).ReadString('\n')

			fmt.Print(message)							
		}						
	}
	conn.Close()	
}