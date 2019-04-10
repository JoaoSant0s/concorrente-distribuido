-package main

import (
	"fmt"
	"math/rand"
	"net/rpc"
	"shared"
	"strconv"
	"time"
)

func clienteRPCTCP() {
	var reply bool

	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.Prime_Port))
	shared.ChecaErro(err, "Não foi possivel conectar ao servidor")

	defer client.Close()
	var average float64

	start := time.Now()
	for i := 0; i < shared.Iterations; i++ {

		t1 := time.Now()

		x := rand.Intn(5000)

		args := shared.Args{A: x}
		fmt.Printf("Número: %d \n", x)

		client.Call("Prime.Prime", args, &reply)

		fmt.Println(reply)

		t2 := time.Now()
		iterationTime := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Printf("Tempo da iteração: %.2fms \n", iterationTime)

		average = average + iterationTime
	}
	average = average / shared.Iterations
	fmt.Printf("Tempo médio por iteração:  %.2fms\n", average)
	totalTime := time.Since(start)
	fmt.Printf("Tempo: %s \n", totalTime)
}

func main() {

	go clienteRPCTCP()

	fmt.Scanln()
}
