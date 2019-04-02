package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"calculadora/impl"
	"shared"
)

func ServidorCalculadoraUDP(){
	msgFromClient := make([]byte, 1024)
	var msgToClient []byte
	var req shared.Request

	addr,err := net.ResolveUDPAddr("udp",":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err,"Não foi posível identificar o endereço do servidor")

	conn, err := net.ListenUDP("udp", addr)
	shared.ChecaErro(err,"Não foi posível criar um handler para o servidor")

	fmt.Println("Servidor pronto para receber solicitações (UDP)...")

	for idx := 0; idx < shared.SAMPLE_SIZE; idx++ {

		// recebe e desserializa request
		n, addr, err := conn.ReadFromUDP(msgFromClient)
		shared.ChecaErro(err,"Não foi posível receber o 'request'")

		err = json.Unmarshal(msgFromClient[:n],&req)
		shared.ChecaErro(err,"Não foi posível desserializar o 'request'")

		//processa a solicitação
		r := impl.Calculadora{}.InvocaCalculadora(req)

		// prepara e serializada resposta
		rep := shared.Reply{r}
		msgToClient,err = json.Marshal(rep)
		shared.ChecaErro(err,"Não foi posível serializar a reposta")

		// envia resposta
		_,err = conn.WriteTo(msgToClient,addr)
		shared.ChecaErro(err,"Não foi posível enviar a resposta")
	}
}

func ServidorCalculadoraTCP(){
	var msgFromClient shared.Request

	// escuta na porta tcp 1313
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err,"Não foi posível criar um handler para o servidor")

	// aceita conexões na porta
	conn, err := ln.Accept()
	shared.ChecaErro(err,"Não foi posível aceitar a conexão")

	// fecha o socket
	defer conn.Close()

	// cria um cofificador/decodificador Json
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	fmt.Println("Servidor pronto para receber solicitações (TCP)...")

	for idx := 0; idx < shared.SAMPLE_SIZE; idx++ {

		// recebe solicitações do cliente e decodifica-as
		err := jsonDecoder.Decode(&msgFromClient)
		shared.ChecaErro(err,"Não foi posível receber o 'request'")

		// processa a solicitação
		r := impl.Calculadora{}.InvocaCalculadora(msgFromClient)

		// envia resposta ao cliente
		msgToClient := shared.Reply{r}
		err = jsonEncoder.Encode(msgToClient)
		shared.ChecaErro(err,"Não foi posível enviar a resposta")
	}
}

func main() {

	go ServidorCalculadoraTCP()
	//go ServidorCalculadoraUDP()

	fmt.Scanln()
}
