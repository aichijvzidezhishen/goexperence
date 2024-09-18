package base

import (
	"reflect"
	"testing"
)

func TestPointRes(t *testing.T) {
	tests := []struct {
		name string
		want *Point
	}{
		// TODO: Add test cases.
		{"t1", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointRes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointRes() = %v, want %v", got, tt.want)
			}
		})
	}
}
