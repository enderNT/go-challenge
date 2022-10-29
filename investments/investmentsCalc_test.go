package investments

import (
	"fmt"
	"testing"
)

type investment struct {
	inv int32
	q1  int32
	q2  int32
	q3  int32
	err error
}

func TestHello (t *testing.T) {
	investmentTest := []investment{
		{
			10500, 7, 7, 7, nil,
		},
		{
			-500, 0, 0, 0,
			fmt.Errorf("the investment introduced is invalid"),
		},
		{
			3000, 2, 2, 2, nil,
		},
		{
			4900, 0, 7, 2, nil,
		},
		{
			80900, 0, 5, 112, nil,
		},
		{
			6750, 0, 0, 0,
			fmt.Errorf("without solution"),
		},
		{
			12860, 0, 0, 0,
			fmt.Errorf("without solution"),
		},
	}

	for _, tt := range investmentTest {
		c := new(Credit)
		q1, q2, q3, _ := c.Assign(tt.inv)

		if tt.q1 != q1 || tt.q2 != q2 || tt.q3 != q3 {
			t.Fatalf(
				"Se esperaba: %d, %d, %d\n pero se obtuvo: %d, %d, %d",
				tt.q1, tt.q2, tt.q3,
				q1, q2, q3,
			)
		}

	}
}

