package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	calc := NewCalculator()
	if calc == nil {
		t.Fatal("NewCalculator returned nil")
	}
}

func TestCalculatorAdd(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
	}
	
	for _, test := range tests {
		result := calc.Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestCalculatorSubtract(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{0, 0, 0},
		{1, 1, 0},
		{10, 15, -5},
	}
	
	for _, test := range tests {
		result := calc.Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestCalculatorMultiply(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 3, -6},
		{4, 4, 16},
	}
	
	for _, test := range tests {
		result := calc.Multiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("Multiply(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestCalculatorDivide(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		a, b, expected float64
	}{
		{6, 2, 3.0},
		{7, 2, 3.5},
		{0, 5, 0.0},
		{-6, 2, -3.0},
	}
	
	for _, test := range tests {
		result := calc.Divide(test.a, test.b)
		if result != test.expected {
			t.Errorf("Divide(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
		}
	}
}

func TestCalculatorPower(t *testing.T) {
	calc := NewCalculator()
	
	tests := []struct {
		base, exponent, expected int
	}{
		{2, 3, 8},
		{5, 0, 1},
		{1, 10, 1},
		{3, 2, 9},
	}
	
	for _, test := range tests {
		result := calc.Power(test.base, test.exponent)
		if result != test.expected {
			t.Errorf("Power(%d, %d) = %d, expected %d", test.base, test.exponent, result, test.expected)
		}
	}
}

func TestStringProcessor(t *testing.T) {
	sp := NewStringProcessor()
	if sp == nil {
		t.Fatal("NewStringProcessor returned nil")
	}
}

func TestStringProcessorReverse(t *testing.T) {
	sp := NewStringProcessor()
	
	tests := []struct {
		input, expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"12345", "54321"},
	}
	
	for _, test := range tests {
		result := sp.Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestStringProcessorIsPalindrome(t *testing.T) {
	sp := NewStringProcessor()
	
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"aba", true},
		{"abc", false},
	}
	
	for _, test := range tests {
		result := sp.IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%s) = %t, expected %t", test.input, result, test.expected)
		}
	}
}

func TestStringProcessorCountWords(t *testing.T) {
	sp := NewStringProcessor()
	
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 2},
		{"", 0},
		{"single", 1},
		{"  multiple   spaces   ", 2},
		{"one two three four", 4},
	}
	
	for _, test := range tests {
		result := sp.CountWords(test.input)
		if result != test.expected {
			t.Errorf("CountWords(%s) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestStringProcessorRemoveSpaces(t *testing.T) {
	sp := NewStringProcessor()
	
	tests := []struct {
		input, expected string
	}{
		{"hello world", "helloworld"},
		{"  test  ", "test"},
		{"no spaces", "nospaces"},
		{"", ""},
	}
	
	for _, test := range tests {
		result := sp.RemoveSpaces(test.input)
		if result != test.expected {
			t.Errorf("RemoveSpaces(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestSorter(t *testing.T) {
	sorter := NewSorter()
	if sorter == nil {
		t.Fatal("NewSorter returned nil")
	}
}

func TestSorterBubbleSort(t *testing.T) {
	sorter := NewSorter()
	
	input := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	result := sorter.BubbleSort(input)
	
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
		return
	}
	
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("At index %d: expected %d, got %d", i, v, result[i])
		}
	}
}

func TestSorterQuickSort(t *testing.T) {
	sorter := NewSorter()
	
	input := []int{3, 6, 8, 10, 1, 2, 1}
	result := sorter.QuickSort(input)
	
	// Check if result is sorted
	for i := 1; i < len(result); i++ {
		if result[i-1] > result[i] {
			t.Errorf("Array is not sorted: %v", result)
			break
		}
	}
}

func TestSorterMergeSort(t *testing.T) {
	sorter := NewSorter()
	
	input := []int{38, 27, 43, 3, 9, 82, 10}
	result := sorter.MergeSort(input)
	
	// Check if result is sorted
	for i := 1; i < len(result); i++ {
		if result[i-1] > result[i] {
			t.Errorf("Array is not sorted: %v", result)
			break
		}
	}
}

func TestSorterIsSorted(t *testing.T) {
	sorter := NewSorter()
	
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1}, true},
		{[]int{}, true},
		{[]int{1, 1, 1}, true},
		{[]int{1, 3, 2}, false},
	}
	
	for _, test := range tests {
		result := sorter.IsSorted(test.input)
		if result != test.expected {
			t.Errorf("IsSorted(%v) = %t, expected %t", test.input, result, test.expected)
		}
	}
}

func TestSorterFindElement(t *testing.T) {
	sorter := NewSorter()
	
	arr := []int{1, 3, 5, 7, 9, 11}
	
	tests := []struct {
		target, expected int
	}{
		{5, 2},
		{1, 0},
		{11, 5},
		{4, -1},
		{0, -1},
	}
	
	for _, test := range tests {
		result := sorter.FindElement(arr, test.target)
		if result != test.expected {
			t.Errorf("FindElement(%v, %d) = %d, expected %d", arr, test.target, result, test.expected)
		}
	}
}

// Benchmark tests
func BenchmarkCalculatorAdd(b *testing.B) {
	calc := NewCalculator()
	for i := 0; i < b.N; i++ {
		calc.Add(123, 456)
	}
}

func BenchmarkCalculatorMultiply(b *testing.B) {
	calc := NewCalculator()
	for i := 0; i < b.N; i++ {
		calc.Multiply(123, 456)
	}
}

func BenchmarkStringReverse(b *testing.B) {
	sp := NewStringProcessor()
	text := "This is a sample text for benchmarking"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sp.Reverse(text)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	sorter := NewSorter()
	data := []int{64, 34, 25, 12, 22, 11, 90, 5, 77, 30}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testData := make([]int, len(data))
		copy(testData, data)
		b.StartTimer()
		
		sorter.BubbleSort(testData)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	sorter := NewSorter()
	data := []int{64, 34, 25, 12, 22, 11, 90, 5, 77, 30}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testData := make([]int, len(data))
		copy(testData, data)
		b.StartTimer()
		
		sorter.QuickSort(testData)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	sorter := NewSorter()
	data := []int{64, 34, 25, 12, 22, 11, 90, 5, 77, 30}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		testData := make([]int, len(data))
		copy(testData, data)
		b.StartTimer()
		
		sorter.MergeSort(testData)
	}
}