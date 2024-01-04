package util

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	// Test with an empty slice
	result := Unique([]int{})
	if len(result) != 0 {
		t.Errorf("Expected an empty slice, but got %v", result)
	}

	// Test with a slice containing duplicate elements
	result = Unique([]int{1, 2, 3, 2, 4, 1, 5})
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test with a slice containing no duplicate elements
	result = Unique([]int{1, 2, 3, 4, 5})
	expected = []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test with a slice of strings
	resultStr := Unique([]string{"apple", "banana", "orange", "banana", "kiwi"})
	expectedStr := []string{"apple", "banana", "orange", "kiwi"}
	if !reflect.DeepEqual(resultStr, expectedStr) {
		t.Errorf("Expected %v, but got %v", expectedStr, resultStr)
	}
}
