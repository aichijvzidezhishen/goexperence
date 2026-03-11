package concurrent

import (
	"testing"
)

func TestOddAndEven(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"OddAndEven"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OddAndEven()
		})
	}
}

func Test_generatePrimes(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"generatePrimes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GeneratePrimes()
		})
	}
}

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Fibonacci", args{20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fibonacci(tt.args.n)
		})
	}
}

func TestCountLinesInFiles(t *testing.T) {
	type args struct {
		files []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"CountLinesInFiles", args{
			files: []string{"countLine1.txt", "countLine2.txt"},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountLinesInFiles(tt.args.files)
		})
	}
}

func TestFastestResponse(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"FastestResponse"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FastestResponse()
		})
	}
}

func TestConcurrentErrorHandling(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"ConcurrentErrorHandling"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConcurrentErrorHandling()
		})
	}
}
