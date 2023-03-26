package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(10, 4)

	if result != 14 {
		t.Errorf("got %d, wanted 14", result)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(10, 6)

	if result != 4 {
		t.Errorf("got %d, wanted 4", result)
	}
}
