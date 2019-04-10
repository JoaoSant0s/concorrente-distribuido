package main

import (
	"fmt"
	"net"
	"net/rpc"
	"primos/impl"
	"shared"
	"strconv"
)

func servidorRPCTCP() {

	// cria instância da calculadora
	primo := new(impl.PrimosRPC)

	// cria um novo servidor rpc e registra a calculadora
	server := rpc.NewServer()
	server.RegisterName("Prime", primo)

	// cria um listen rpc-sender
	l, err := net.Listen("tcp", ":"+strconv.Itoa(shared.Prime_Port))
	shared.ChecaErro(err, "Servidor não está pronto")

	// aguarda por chamadas
	fmt.Println("Servidor pronto (RPC TCP) ...")
	server.Accept(l)
}

func main() {

	go servidorRPCTCP()

	fmt.Scanln()
}
