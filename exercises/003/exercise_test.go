package main

import (
	"reflect"
	"testing"
)

func TestExercise003(t *testing.T) {
	want := map[int]int{
		1: 1,
		2: 4,
		3: 9,
		4: 16,
		5: 25,
		6: 36,
		7: 49,
		8: 64,
	}
	got := Exercise003(8)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Exercise003(8) = %v, want %v", got, want)
	}
}
