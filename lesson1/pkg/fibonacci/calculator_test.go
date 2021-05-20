package fibonacci_test

import (
	fibo "go-core/lesson1/pkg/fibonacci"
	"testing"
)

func TestCalculate(t *testing.T) {
	testData := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Fibo(0)",
			n:    0,
			want: 0,
		},
		{
			name: "Fibo(0)",
			n:    1,
			want: 1,
		},
		{
			name: "Fibo(0)",
			n:    2,
			want: 1,
		},
		{
			name: "Fibo(0)",
			n:    3,
			want: 2,
		},
		{
			name: "Fibo(0)",
			n:    4,
			want: 3,
		},
		{
			name: "Fibo(0)",
			n:    5,
			want: 5,
		},
		{
			name: "Fibo(0)",
			n:    10,
			want: 55,
		},
		{
			name: "Fibo(0)",
			n:    15,
			want: 610,
		},
		{
			name: "Fibo(0)",
			n:    20,
			want: 6765,
		},
	}
	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			if got := fibo.Calculate(test.n); got != test.want {
				t.Errorf("Calculate() = %v, want %v", got, test.want)
			}
		})
	}
}
