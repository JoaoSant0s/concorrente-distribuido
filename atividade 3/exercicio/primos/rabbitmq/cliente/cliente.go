package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"shared"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// conecta ao servidor de mensageria
	conn, err := amqp.Dial("amqp://guest:guest@localhost:8080/")
	shared.ChecaErro(err, "Não foi possível se conectar ao servidor de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()

	// declara as filas
	requestQueue, err := ch.QueueDeclare(
		"request", false, false, false, false, nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no servidor de mensageria")

	replyQueue, err := ch.QueueDeclare(
		"response", false, false, false, false, nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no servidor de mensageria")
	
	// cria consumidor
	msgsFromServer, err := ch.Consume(replyQueue.Name, "", true, false,
		false, false, nil)
	shared.ChecaErro(err, "Falha ao registrar o consumidor servidor de mensageria")

	var average float64

	start := time.Now()

	//start := time.Now()
	for i := 0; i < shared.Iterations; i++ {

		t1 := time.Now()
		y := rand.Intn(5000)

		// prepara request
		msgRequest := shared.Request{Op: "prime", P1: y}
		msgRequestBytes, err := json.Marshal(msgRequest)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		// publica request
		err = ch.Publish("", requestQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o servidor de mensageria")

		fmt.Println(<-msgsFromServer)

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Printf("Tempo da iteração: %.2fms \n", x)
		average = average + x
	}
	average = average / shared.Iterations
	fmt.Printf("Tempo médio por iteração:  %.2fms\n", average)
	totalTime := time.Since(start)
	fmt.Printf("Tempo: %s \n", totalTime)
}
