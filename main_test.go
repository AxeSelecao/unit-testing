package main

import "testing"

func TestMax(t *testing.T) {
	numbers := []int{1, 2, 3, -10, 150, 160, 6, 9, 99, -5, 0}
	expected := 150

	result := Max(numbers)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

}
