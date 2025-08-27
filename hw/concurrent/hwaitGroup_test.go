package concurrent

import (
	"testing"
)

func Test_proessMutiItems(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"t1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proessMutiItems()
		})
	}
}

func TestNestingWaitGroup(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"t1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NestingWaitGroup()
		})
	}
}
