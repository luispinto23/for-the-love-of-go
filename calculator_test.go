package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

// test divide by zero
func TestDivideInvalid(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "divide by zero",
			args: args{a: 1, b: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Divide(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
			if err == nil {
				t.Error("Divide(), want error for invalid input, got nil")
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "2-2", args: args{a: 2, b: 2}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculator.Subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "2x2", args: args{a: 2, b: 2}, want: 4},
		{name: "2x5", args: args{a: 2, b: 5}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculator.Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	type args struct {
		a float64
		b float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "sqrt(4)", args: args{a: 4}, want: 2},
		{name: "sqrt(9)", args: args{a: 9}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Sqrt(tt.args.a)
			if got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
			if err != nil {
				t.Errorf("Sqrt() = %v, want %v", err, nil)
			}
		})
	}
}
