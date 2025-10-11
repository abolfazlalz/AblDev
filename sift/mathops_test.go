package mathops

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"two positives", 2, 3, 5},
		{"positive and negative", 5, -3, 2},
		{"two negatives", -2, -4, -6},
		{"zero addition", 0, 5, 5},
		{"float addition", 1.5, 2.3, 3.8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"two positives", 10, 4, 6},
		{"positive minus negative", 5, -2, 7},
		{"negative minus positive", -5, 3, -8},
		{"two negatives", -5, -3, -2},
		{"zero minus number", 0, 3, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"two positives", 3, 5, 15},
		{"positive and negative", 4, -2, -8},
		{"two negatives", -3, -3, 9},
		{"multiply by zero", 7, 0, 0},
		{"floats", 1.5, 2.0, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"normal division", 10, 2, 5},
		{"divide by one", 9, 1, 9},
		{"divide by negative", 9, -3, -3},
		{"negative divided by positive", -8, 2, -4},
		{"two negatives", -9, -3, 3},
		{"divide by zero", 10, 0, 0},
		{"zero divided by number", 0, 5, 0},
		{"float division", 7.5, 2.5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Divide(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
