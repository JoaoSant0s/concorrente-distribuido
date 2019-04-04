package shared

import (
	"log"
)

const Iterations = 5000
const Prime_Port = 8080

type Args struct {
	A int
}

func ChecaErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
	//fmt.Println(msg)
}
