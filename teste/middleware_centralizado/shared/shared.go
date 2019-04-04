package shared

import (
	"log"
)

const SAMPLE_SIZE = 5000
const CALCULATOR_PORT = 1313

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Request struct {
	Op string;
	P1 int;
	P2 int;
}

type Reply struct {
	Result int;
}

func ChecaErro(err error, msg string){
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
	//fmt.Println(msg)
}