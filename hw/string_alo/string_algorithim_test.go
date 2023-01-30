package string_alo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointMove(t *testing.T) {
	x, y := PointMove("A10;W10")
	// Slice = append(Slice, a, b)
	assert.Equal(t, x, 10, "Point move error")
	assert.Equal(t, y, 10, "Point move error")
}

func TestIdentifyIp(t *testing.T) {
	ipstr := "0.70.44.68~255.254.255.0"
	IdentifyIp(ipstr)
}
