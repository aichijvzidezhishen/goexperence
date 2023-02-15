package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 近似
func TestApproximate(t *testing.T) {
	assert.Equal(t, 1, Approximate(1.25))
	assert.Equal(t, 2, Approximate(1.85))
}

//
