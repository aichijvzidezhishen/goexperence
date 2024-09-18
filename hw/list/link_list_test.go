package list

import "testing"

func TestDelStable(t *testing.T) {
	type args struct {
		pos int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"t1", args{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelStable(tt.args.pos)
		})
	}
}
