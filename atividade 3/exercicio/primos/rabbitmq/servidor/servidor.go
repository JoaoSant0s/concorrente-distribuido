package main

import (
	"encoding/json"
	"fmt"-
	"log"
	"primos/impl"
	"shared"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao servidor")
	defer conn.Close()

	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o servidor")
	defer ch.Close()

	requestQueue, err := ch.QueueDeclare("request", false, false, false,
		false, nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no servidor")

	replyQueue, err := ch.QueueDeclare("response", false, false, false,
		false, nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no servidor")

	msgsFromClient, err := ch.Consume(requestQueue.Name, "", true, false,
		false, false, nil)
	shared.ChecaErro(err, "Falha ao registrar o consumidor do servidor")

	fmt.Println("Servidor pronto...")
	for d := range msgsFromClient {

		msgRequest := shared.Request{}
		err := json.Unmarshal(d.Body, &msgRequest)
		shared.ChecaErro(err, "Falha ao desserializar a mensagem")

		r := impl.Primo{}.InvocaPrimo(msgRequest)

		replyMsg := shared.Reply{Result: r}
		replyMsgBytes, err := json.Marshal(replyMsg)
		shared.ChecaErro(err, "Não foi possível criar a fila no servidor")
		if err != nil {
			log.Fatalf("%s: %s", "Falha ao serializar mensagem", err)
		}

		err = ch.Publish("", replyQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: []byte(replyMsgBytes)})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o servidor")
	}
}
