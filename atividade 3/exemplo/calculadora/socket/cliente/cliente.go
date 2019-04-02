package main

import (
	"net"
	"fmt"
	"encoding/json"
	"time"
	"strconv"
	"shared"
)

func ClienteCalculadoraUDP() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err,"Não foi possível encontrar o servidor,")

	// connecta ao servidor
	conn, err := net.DialUDP("udp", nil, addr)
	shared.ChecaErro(err,"Não foi possivel criar um handler para o servidor,")
	defer conn.Close()

	var msgToServer []byte
	var rep shared.Reply
	msgFromServer := make([]byte, 1024)

	// loop
	//start := time.Now()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		t1 := time.Now()

		// prepara e serializa request
		req := shared.Request{Op:"add",P1:i,P2:i}
		msgToServer,err = json.Marshal(req)
		shared.ChecaErro(err,"Não foi posível serializar o 'request'")

		// envia request
		_,err := conn.Write(msgToServer)
		shared.ChecaErro(err,"Não foi posível enviar o 'request'")

		// recebe resposta
		n, _, err := conn.ReadFromUDP(msgFromServer)
		shared.ChecaErro(err,"Não foi posível receber a 'resposta'")

		// desserializa resposta
		err = json.Unmarshal(msgFromServer[:n],&rep)
		shared.ChecaErro(err,"Não foi posível desserializar a 'resposta'")

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Println(x)
	}
	//elapsed := time.Since(start)
	//fmt.Printf("Tempo: %s \n", elapsed)
}

func ClienteCalculadoraTCP() {

	// conecta ao servidor
	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err,"Não foi posível estabelecer uma conexão com o servidor")

	// cria um decoder/encoder json
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	var msgFromServer shared.Reply

	// loop
	//start := time.Now()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		t1 := time.Now()

		// prepara request
		msgToServer := shared.Request{"add", i, i}

		// envia request
		err = jsonEncoder.Encode(msgToServer)
		shared.ChecaErro(err,"Não foi posível enviar o 'request'")

		// recebe resposta
		err = jsonDecoder.Decode(&msgFromServer)
		shared.ChecaErro(err,"Não foi posível receber a 'resposta'")

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Println(x)
	}
	//elapsed := time.Since(start)
	//fmt.Printf("Tempo: %s \n", elapsed)
}

func main() {
	//go ClienteCalculadoraUDP()
	go ClienteCalculadoraTCP()

	fmt.Scanln()
}
