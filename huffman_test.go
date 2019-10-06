package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestUnpack(t *testing.T) {
	t.Run("handle basic case", func(t *testing.T) {
		//got := Hello("z x x x r r r r t t")
		//want := []string{"r", "x", "t", "z"}

		assert.Equal(t, Hello(), "Hello")
	})
}
