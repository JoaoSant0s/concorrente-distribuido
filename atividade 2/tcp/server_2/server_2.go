package main

import (
	"fmt"
	"net"
	"bufio"	
	"math"
	"strconv"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":1313")
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

		value := extractNumberString(message)		
		
		i, err := strconv.Atoi(value)

		if err != nil {
			conn.Write([]byte("-1\n"))
		} else {	
			prime := isPrimo(i)				
			conn.Write([]byte(strconv.FormatBool(prime) + "\n"))	
		}
	}
}

func extractNumberString(value string) string{
	newValue := strings.ReplaceAll(value, "\r", "")
	newValue = strings.ReplaceAll(newValue, "\n", "")
	return newValue
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