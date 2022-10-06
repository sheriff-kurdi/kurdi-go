package test

import (
	"github.com/stretchr/testify/assert"
	"kurdi-go/math"
	"testing"
)

func TestAdd(test *testing.T) {
	assert := assert.New(test)

	actual := math.Add(1, 2)
	expected := 3

	assert.Equal(expected, actual, "The two numbers should be equals.")
}

func TestAdd2(test *testing.T) {
	assert := assert.New(test)

	actual := math.Add(1, 2)
	expected := 3

	assert.Equal(expected, actual, "The two numbers should be equals.")

}

func TestAdd3(test *testing.T) {
	assert := assert.New(test)

	actual := math.Add(1, 2)
	expected := 3

	assert.Equal(expected, actual, "The two numbers should be equals.")

}
