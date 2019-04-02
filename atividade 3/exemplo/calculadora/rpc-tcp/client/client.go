package main
import (
	"net/rpc"
	"fmt"
	"time"
	"shared"
	"strconv"
)

func clienteRPCTCP(){
	var reply int

	// conectar ao servidor
	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err,"Servidor não está pronto")

	defer client.Close()

	// loop
	//start := time.Now()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		t1 := time.Now()

		// prepara request
		args := shared.Args{A: i, B: i}

		// invoca request
		client.Call("Calculadora.Add", args, &reply)

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Println(x)
	}
	//elapsed := time.Since(start)
	//fmt.Printf("Tempo: %s \n", elapsed)
}

func main() {

	go clienteRPCTCP()

	fmt.Scanln()
}

