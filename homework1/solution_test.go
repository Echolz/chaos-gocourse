package homework1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolutionSimpleValues(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		m           int
		wantedValue int
	}{
		{"test1", 8, 3, 7},
		{"test2", 11, 5, 8},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantedValue, findWinner(tc.n, tc.m))
		})
	}
}
