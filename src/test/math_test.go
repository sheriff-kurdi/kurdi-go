package test

import (
	"testing"
)

func TestAdd(test *testing.T) {

	actual := 3
	expected := 3

	if actual != expected {
		test.Errorf("actual %v, expected %v", actual, expected)
	}
}
