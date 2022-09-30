package test

import (
	"kurdi-go/math"
	"testing"
)

func TestAdd(test *testing.T) {

	actual := math.Add(1, 2)
	expected := 3

	if actual != expected {
		test.Errorf("actual %v, expected %v", actual, expected)
	}
}
