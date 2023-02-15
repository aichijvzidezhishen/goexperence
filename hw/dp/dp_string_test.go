package main

import "testing"

func TestDpAloLevenshtein(t *testing.T) {
	type args struct {
		param []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"t1", args{[]string{"abcdefg", "abcdefg"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DpAloLevenshtein(tt.args.param)
		})
	}
}
