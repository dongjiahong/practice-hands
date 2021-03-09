package jsonrpc

import (
	"context"
	"errors"
	"log"
	"math"

	"github.com/semrush/zenrpc/v2"
)

type ArithService struct{ zenrpc.Service }

func (as ArithService) Sum(ctx context.Context, a, b int) (int, *zenrpc.Error) {
	log.Println("a: ", a, " b: ", b)
	return a + b, nil
}

func (as ArithService) Multiply(a, b int) int {
	return a * b
}

type Quotient struct {
	Quo, Rem int
}

func (as ArithService) Divide(a, b int) (quo *Quotient, err error) {
	if b == 0 {
		return nil, errors.New("divide by zero")
	} else if b == 1 {
		return nil, zenrpc.NewError(401, errors.New("we do not serve 1"))
	}
	return &Quotient{
		Quo: a / b,
		Rem: a % b,
	}, nil
}

//zenrpc:exp=2
func (as ArithService) Pow(base float64, exp float64) float64 {
	return math.Pow(base, exp)
}

//go:generate zenrpc
