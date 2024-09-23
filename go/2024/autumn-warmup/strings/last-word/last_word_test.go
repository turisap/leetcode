package l_word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastWordNaive(t *testing.T) {
	tcc := []struct {
		input string
		out   int
	}{
		{
			input: "Hello World",
			out:   5,
		},
		{
			input: "   fly me   to   the moon  ",
			out:   4,
		},
	}

	for _, tc := range tcc {
		l := lastWordNaive(tc.input)
		assert.Equal(t, tc.out, l)
	}
}

func TestLastWordIterative(t *testing.T) {
	tcc := []struct {
		input string
		out   int
	}{
		{
			input: "Hello World",
			out:   5,
		},
		{
			input: "   fly me   to   the moon  ",
			out:   4,
		},
	}

	for _, tc := range tcc {
		l := lastWordIterative(tc.input)
		assert.Equal(t, tc.out, l)
	}
}
