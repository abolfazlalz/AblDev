package mathops

type Numeric interface {
	int | uint | float64
}

// Add returns the sum of two numbers.
func Add[T Numeric](a, b T) T {
	return a + b
}

// Subtract returns the difference between two numbers.
func Subtract[T Numeric](a, b T) T {
	return a - b
}

// Multiply returns the product of two numbers.
func Multiply[T Numeric](a, b T) T {
	return a * b
}

// Divide returns the quotient of two numbers.
// If b is 0, it returns 0 to avoid panic.
func Divide[T Numeric](a, b T) T {
	if b == 0 {
		return 0
	}
	return a / b
}
