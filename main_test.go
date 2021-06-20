package main

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func compareArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestSimple(t *testing.T) {
	numbers := []int{2, -1, 5, 6, 88, 5, 25, 85, 99}
	result := mergeSort(numbers)
	expectedResult := []int{-1, 2, 5, 5, 6, 25, 85, 88, 99}
	testResult := compareArray(result, expectedResult)
	assertEqual(t, testResult, true, "The should be true.")
}
