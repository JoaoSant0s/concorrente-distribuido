package main

import (
	"fmt"
	"net"
	"bufio"	
	"os"
	"time"
)

func main() {
	fmt.Println("Verificador de números primos")

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	} else {				
		start(conn)
	}
	defer conn.Close()	
}

func start(conn net.Conn){
	for {
		reader := bufio.NewReader(os.Stdin)	
		fmt.Print("Digite um número: ")	
		text, _ := reader.ReadString('\n')

		start := time.Now()
		fmt.Fprintf(conn, text + "\n")			

		message, _ := bufio.NewReader(conn).ReadString('\n')

		t := time.Now()
		elapsed := t.Sub(start)
		
		fmt.Print(message)
		fmt.Printf("The call took %v to run.\n\n", elapsed)		
	}	
}