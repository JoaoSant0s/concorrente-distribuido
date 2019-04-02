package main

import (
	"time"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
	"shared"
	"calculadora/grpc/calculadora"
)

func main() {
	var idx int32

	// Estabelece conexão com o servidor
	conn, err := grpc.Dial("localhost"+":"+strconv.Itoa(shared.CALCULATOR_PORT), grpc.WithInsecure())
	shared.ChecaErro(err,"Não foi possível se conectar ao servidor")
	defer conn.Close()

	calc := calculadora.NewGreeterClient(conn)

	// contacta o servidor
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for idx = 0; idx < shared.SAMPLE_SIZE; idx++ {
		t1 := time.Now()

		// invoca operação remota
		calc.Add(ctx, &calculadora.Request{Op:"add",P1:idx,P2:idx})

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Printf("%f \n", x)
	}
}


