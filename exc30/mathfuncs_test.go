package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {

	expected := 2
	actual := add(1, 1)

	assert.Equal(t, expected, actual)
}

func TestSubtract(t *testing.T) {

	expected := 3
	actual := subtract(6, 3)

	assert.Equal(t, expected, actual)

}

func TestDoMath(t *testing.T) {

	anyFunc := func(x int, y int) int {
		return x * y
	}

	expected := 20

	actual := doMath(4, 5, anyFunc)

	assert.Equal(t, expected, actual)
}

func TestDoMathWithZero(t *testing.T) {

	anyFunc := func(x int, y int) int {
		return x * y
	}

	expected := 0

	actual := doMath(0, 5, anyFunc)

	assert.Equal(t, expected, actual)
}
