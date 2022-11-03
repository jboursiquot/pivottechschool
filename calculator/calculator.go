package calculator

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

type ErrDivideByZero struct{}

// Error implements error
func (ErrDivideByZero) Error() string {
	return "divide by zero"
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero{}
	}
	return a / b, nil
}
