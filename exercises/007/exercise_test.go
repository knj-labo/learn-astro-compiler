package main

import (
	"reflect"
	"testing"
)

func TestFilterEven(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{[]int{1, 3, 5, 7}, []int{}},
		{[]int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
		{[]int{}, []int{}},
	}
	
	for _, test := range tests {
		result := FilterEven(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("FilterEven(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestMapSquare(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 4, 9, 16}},
		{[]int{0, 1, -2, 3}, []int{0, 1, 4, 9}},
		{[]int{}, []int{}},
		{[]int{5}, []int{25}},
	}
	
	for _, test := range tests {
		result := MapSquare(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapSquare(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestSortByLength(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"apple", "go", "banana", "hi"}, []string{"go", "hi", "apple", "banana"}},
		{[]string{"a", "bb", "ccc"}, []string{"a", "bb", "ccc"}},
		{[]string{"programming", "go", "lang"}, []string{"go", "lang", "programming"}},
		{[]string{}, []string{}},
	}
	
	for _, test := range tests {
		result := SortByLength(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortByLength(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 2, 3, 3, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}},
		{[]int{5, 3, 5, 1, 3, 1}, []int{5, 3, 1}},
	}
	
	for _, test := range tests {
		result := RemoveDuplicates(test.input)
		
		// 順序は保持されるべき（最初の出現順）
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("RemoveDuplicates(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}