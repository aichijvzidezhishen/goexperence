package string_alo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPointMove(t *testing.T) {
	x, y := PointMove("A10;W10")

	assert.Equal(t, x, 10, "Point move error")
	assert.Equal(t, y, 10, "Point move error")

}

func TestIdentifyIp(t *testing.T) {
	ipstr := "0.70.44.68~255.254.255.0"
	IdentifyIp(ipstr)
}

func TestCharStatistic(t *testing.T) {
	CharStatistic()
}

func TestGetWeek(t *testing.T) {
	t1 := time.Date(2023, 4, 16, 4, 0, 0, 0, time.Local)
	t2 := time.Date(2023, 4, 17, 4, 0, 0, 0, time.Local)

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"n1", args{t1}, ""},
		{"n1", args{t2}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeek(tt.args.t); got != tt.want {
				t.Errorf("GetWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
