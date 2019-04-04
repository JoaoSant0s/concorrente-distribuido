package main

import (
"net"
"fmt"
"net/rpc"
	"calculadora/impl"
	"strconv"
	"shared"
)

func servidorRPCTCP(){

	// cria instância da calculadora
	calculadora := new(impl.CalculadoraRPC)

	// cria um novo servidor rpc e registra a calculadora
	server := rpc.NewServer()
	server.RegisterName("Calculadora", calculadora)

	// cria um listen rpc-sender
	l, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err,"Servidor não está pronto")


	// aguarda por chamadas
	fmt.Println("Servidor pronto (RPC TCP) ...")
	server.Accept(l)
}

func main() {

	go servidorRPCTCP()

	fmt.Scanln()
}

