package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBitValue(t *testing.T) {
	assert.Equal(t, 1, BitCount1(1), "equal err")
}

func TestSetBitValue1(t *testing.T) {
	assert.Equal(t, 2, BitCount1(2))
}
