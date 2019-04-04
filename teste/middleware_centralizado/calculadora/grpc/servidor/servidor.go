package main

import (
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"strconv"
	"fmt"
	"calculadora/grpc/calculadora"
	"shared"
)

type servidorCalculadora struct{}

func (s *servidorCalculadora) Add(ctx context.Context, in *calculadora.Request) (*calculadora.Reply, error) {
	return &calculadora.Reply{N: in.P1+in.P2}, nil
}

func (s *servidorCalculadora) Sub(ctx context.Context, in *calculadora.Request) (*calculadora.Reply, error) {
	return &calculadora.Reply{N: in.P1-in.P2}, nil
}

func (s *servidorCalculadora) Mul(ctx context.Context, in *calculadora.Request) (*calculadora.Reply, error) {
	return &calculadora.Reply{N: in.P1*in.P2}, nil
}

func (s *servidorCalculadora) Div(ctx context.Context, in *calculadora.Request) (*calculadora.Reply, error) {
	return &calculadora.Reply{N: in.P1/in.P2}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err,"Não foi possível criar o listener")

	servidor := grpc.NewServer()
	calculadora.RegisterGreeterServer(servidor, &servidorCalculadora{})

	fmt.Println("Servidor pronto ...")

	// Register reflection service on gRPC servidor.
	reflection.Register(servidor)

	err = servidor.Serve(conn);
	shared.ChecaErro(err,"Falha ao servir")
}
