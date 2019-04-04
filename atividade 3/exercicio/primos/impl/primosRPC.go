package impl

import (
	"math"
	"shared"
)

//PrimosRPC is service to check if number is prime or not
type PrimosRPC struct{}

//Prime is method to check if a number is Prime or Not
func (t *PrimosRPC) Prime(args *shared.Args, reply *bool) error {
	*reply = isPrimo(args.A)

	return nil
}

func isPrimo(n int) bool {
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
