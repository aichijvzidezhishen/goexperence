package concurrent

import "testing"

func Test_impMutiHttpGet(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impMutiHttpGet()
		})
	}
}
