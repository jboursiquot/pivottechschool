package calculator_test

import (
	"testing"

	"github.com/jboursiquot/pivottechschool/calculator"
)

func TestCalculator(t *testing.T) {
	tests := map[string]struct {
		a, b, want int
		op         func(int, int) int
		opWithErr  func(int, int) (int, error)
		err        error
	}{
		"Add":          {a: 1, b: 2, want: 3, op: calculator.Add},
		"Subtract":     {a: 1, b: 2, want: -1, op: calculator.Subtract},
		"Multiply":     {a: 2, b: 3, want: 6, op: calculator.Multiply},
		"Divide":       {a: 6, b: 3, want: 2, opWithErr: calculator.Divide},
		"DivideByZero": {a: 6, b: 0, want: 0, opWithErr: calculator.Divide, err: calculator.ErrDivideByZero{}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.op != nil {
				got := test.op(test.a, test.b)
				if got != test.want {
					t.Errorf("got %d, want %d", got, test.want)
				}
				return
			}

			got, err := test.opWithErr(test.a, test.b)

			if test.err != nil {
				if err == nil {
					t.Error("expected error, got nil")
				}
				if err.Error() != test.err.Error() {
					t.Errorf("got %q, want %q", err, test.err)
				}
			}

			if test.err == nil && err != nil {
				t.Errorf("got %q, want nil", err)
			}

			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
