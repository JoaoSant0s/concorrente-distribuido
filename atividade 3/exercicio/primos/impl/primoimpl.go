package impl

import (
	"math"
	"shared"
)

type Primo struct{}

func (Primo) InvocaPrimo(req shared.Request) bool {
	var r bool

	op := req.Op
	p1 := req.P1

	switch op {
	case "prime":
		r = Primo{}.Prime(p1)
	}
	return r
}

func (Primo) Prime(x int) bool {
	return _isPrimo(x)
}

func _isPrimo(n int) bool {
	if n < 5 || n%2 == 0 || n%3 == 0 {
		return (n == 2 || n == 3)
	}

	maxP := math.Sqrt(float64(n)) + 2

	for p := 5; float64(p) < maxP; p += 6 {
		if n%p == 0 || n%(p+2) == 0 {
			return false
		}
	}

	return true
}
