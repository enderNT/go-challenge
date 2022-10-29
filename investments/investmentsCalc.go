package investments

import (
	"errors"
	"fmt"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type Credit struct {
	Q1 int32
	Q2 int32
	Q3 int32
}

func (credit *Credit) Assign(investment int32) (int32, int32, int32, error) {
	credit.Q1 = 300
	credit.Q2 = 500
	credit.Q3 = 700
	totalAmount := credit.Q1 + credit.Q2 + credit.Q3

	if investment <= 0 {
		return 0, 0, 0, errors.New("the investment introduced is invalid")
	}

	if investment % totalAmount == 0 {
		result := investment / totalAmount
		fmt.Printf(
			"%d de 300, %d de 500, %d de 700",
			result, result, result,
		)
		return result, result, result, nil
	}

	for i := 0.0; i < float64(investment/credit.Q1); i++ {
		for j := 0.0; j < float64(investment/credit.Q2); j++ {
			for k := 0.0; k < float64(investment/credit.Q3); k++ {
				if 300 * i + 500 * j + 700 * k == float64(investment) {
					return int32(i), int32(j), int32(k), nil
				}
			}
		}
	}

	return 0, 0, 0, errors.New("without solution")
}
