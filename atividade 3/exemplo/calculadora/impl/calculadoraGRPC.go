package impl

import (
	"context"
	"calculadora/grpc/calculadora"
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
