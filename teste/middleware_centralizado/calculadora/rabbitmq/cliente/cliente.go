package main

import (
	"github.com/streadway/amqp"
	"encoding/json"
	"shared"
	"time"
	"fmt"
)

func main() {
	// conecta ao servidor de mensageria
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err,"Não foi possível se conectar ao servidor de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err,"Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()

	// declara as filas
	requestQueue, err := ch.QueueDeclare(
		"request", false, false, false, false, nil, )
	shared.ChecaErro(err,"Não foi possível criar a fila no servidor de mensageria")

	replyQueue, err := ch.QueueDeclare(
		"response", false, false, false, false, nil, )
	shared.ChecaErro(err,"Não foi possível criar a fila no servidor de mensageria")

	// cria consumidor
	msgsFromServer, err := ch.Consume(replyQueue.Name, "", true, false,
		false, false, nil, )
	shared.ChecaErro(err,"Falha ao registrar o consumidor servidor de mensageria")

	//start := time.Now()
	for i := 0; i<shared.SAMPLE_SIZE; i++{

		t1 := time.Now()

		// prepara request
		msgRequest := shared.Request{Op:"add",P1:i,P2:i}
		msgRequestBytes,err := json.Marshal(msgRequest)
		shared.ChecaErro(err,"Falha ao serializar a mensagem")

		// publica request
		err = ch.Publish("", requestQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes,})
		shared.ChecaErro(err,"Falha ao enviar a mensagem para o servidor de mensageria")

		// recebe resposta
		<- msgsFromServer

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Println(x)
	}
	//elapsed := time.Since(start)
	//fmt.Printf("Tempo: %s \n", elapsed)
}
