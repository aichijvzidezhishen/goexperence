package hash

import (
	"testing"
)

func TestMergeTableRecord(t *testing.T) {
	MergeTableRecord()
}

func TestNorepeatNum(t *testing.T) {
	tests := []struct {
		name string
		num  int
	}{
		// TODO: Add test cases.
		{"test1", 10},
		{"test2", 120},
		{"test3", 1022},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NorepeatNum(tt.num)
		})
	}
}

func TestCharNum(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{"abc"}},
		{"t1", args{"abca"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CharNum(tt.args.n)
		})
	}
}

func TestReverseNum(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"t1", args{"dsadasd"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseNum(tt.args.str)
		})
	}
}
