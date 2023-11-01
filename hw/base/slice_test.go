package base

import (
	"reflect"
	"testing"
)

func TestEntrance1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Entrance1()
		})
	}
}

func TestEntrance2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Entrance2()
		})
	}
}

func TestEntrance3(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Entrance3()
		})
	}
}

func TestEvents_Refresh(t *testing.T) {
	es1 := Events{
		[]Event{{
			Eventid: 1,
			Idlist:  []string{"aa", "bb"},
		},
			{
				Eventid: 2,
				Idlist:  []string{"cc", "dd"},
			},
		},
	}
	for k, v := range es1.List {
		PrintSliceStruct(&v.Idlist, "construct replicate ")
		PrintSliceStruct(&es1.List[k].Idlist, "construct index ")
	}

	es1.Refresh()

}

func TestDeepCopy(t *testing.T) {
	type args struct {
		dist []int32
	}
	tests := []struct {
		name   string
		args   args
		wantS1 []int32
	}{
		// TODO: Add test cases.
		// fmt.Println("")
		{"t1", args{[]int32{1, 2, 3}}, []int32{1, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS1 := DeepCopy(tt.args.dist); !reflect.DeepEqual(gotS1, tt.wantS1) {
				t.Errorf("DeepCopy() = %v, want %v", gotS1, tt.wantS1)
			}
		})
	}
}
