package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Passe o número inteiro de testes a serem executados")
	} else {
		numberOftests, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println("Passe o número inteiro válido")
		} else {
			fmt.Println("Verificador de números primos")

			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println(err)
			} else {
				start(conn, numberOftests)
			}
			defer conn.Close()
		}
	}
}

func start(conn net.Conn, numberOftests int) {
	primes := 0
	noPrimes := 0

	var totalTime time.Duration

	for index := 0; index < numberOftests; index++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		newNumber := r1.Intn(4294967295)
		numberString := strconv.Itoa(newNumber)
		fmt.Println("Número gerado: " + numberString)

		start := time.Now()
		fmt.Fprintf(conn, numberString+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')

		t := time.Now()
		elapsed := t.Sub(start)
		totalTime += elapsed

		if strings.Contains(message, "não é") {
			noPrimes++
		} else {
			primes++
		}

		fmt.Print(message)
		fmt.Printf("The call took %v to run.\n\n", elapsed)
	}

	primesString := strconv.Itoa(primes)
	noPrimeString := strconv.Itoa(noPrimes)

	fmt.Println("Número de primos: " + primesString)
	fmt.Println("Número de não primos: " + noPrimeString)
	fmt.Println(totalTime)
}
