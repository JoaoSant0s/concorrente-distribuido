package shared

import (
	"log"
)

const Iterations = 5000
const Prime_Port = 8080

type Args struct {
	A int
}

type Request struct {
	Op string
	P1 int
}

type Reply struct {
	Result bool
}

func ChecaErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
	//fmt.Println(msg)
}
