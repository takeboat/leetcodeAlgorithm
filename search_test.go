package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHindex(t *testing.T) {
	citations := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ans := hIndex(citations)
	assert.Equal(t, 5, ans)
}

func TestHindex1(t *testing.T) {
	citations := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ans := hIndex(citations)
	assert.Equal(t, 6, ans)
}
