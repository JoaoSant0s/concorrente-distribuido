package main

import (
	"fmt"
	"net"	
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Initing server 2")

	pc, err := net.ListenPacket("udp", ":1313")
	if err != nil {
		fmt.Println(err)
	}	

	defer pc.Close()	
	
	fmt.Println("Establishing readers")

	start(pc)
}

func start(pc net.PacketConn){
	for {
		buffer := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buffer)

		if err != nil{			
			fmt.Println(err)	
		}else{						
			s := string(buffer[:n])
			message := TrimString(s)	
			
			i, err := strconv.Atoi(message)

			fmt.Println("Checking input ... " + message)

			if err != nil {
				pc.WriteTo([]byte("-1\n"), addr)
			} else {	
				prime := isPrimo(i)				
				pc.WriteTo([]byte(strconv.FormatBool(prime) + "\n"), addr)	
			}			
		}			
	}
}

func TrimString(value string) string{
	return strings.Trim(value, "\n\r ")	
}

func isPrimo(n int) bool {
	if n<5 || n%2==0 || n%3==0{
		return (n==2||n==3)
	}
	
	maxP := math.Sqrt(float64(n)) + 2

	for p := 5; float64(p) < maxP ; p+=6 {
		if(n%p==0||n%(p+2)==0){
			return false
		}
	}

	return true
}