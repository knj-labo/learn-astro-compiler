package main

import (
	"reflect"
	"strings"
	"testing"
)

// テスト用の構造体
type TestStruct struct {
	Name     string `json:"name" required:"true"`
	Age      int    `json:"age" required:"true"`
	Email    string `json:"email"`
	Optional string `json:"optional"`
}

func TestStructInfo(t *testing.T) {
	ts := TestStruct{
		Name:  "Test",
		Age:   25,
		Email: "test@example.com",
	}
	
	// この関数は出力をテストするのが難しいので、パニックしないことを確認
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("StructInfo panicked: %v", r)
		}
	}()
	
	StructInfo(ts)
}

func TestDeepCopy(t *testing.T) {
	original := TestStruct{
		Name:  "Original",
		Age:   30,
		Email: "original@example.com",
	}
	
	copied, err := DeepCopy(original)
	if err != nil {
		t.Fatalf("DeepCopy failed: %v", err)
	}
	
	copiedStruct, ok := copied.(TestStruct)
	if !ok {
		t.Fatalf("DeepCopy returned wrong type: %T", copied)
	}
	
	// 値が同じかチェック
	if !reflect.DeepEqual(original, copiedStruct) {
		t.Errorf("DeepCopy values don't match: original %+v, copied %+v", original, copiedStruct)
	}
	
	// 異なるメモリアドレスかチェック（ポインタの場合）
	if reflect.ValueOf(original).Pointer() == reflect.ValueOf(copiedStruct).Pointer() {
		t.Error("DeepCopy returned same reference")
	}
}

func TestDeepCopySlice(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	
	copied, err := DeepCopy(original)
	if err != nil {
		t.Fatalf("DeepCopy failed: %v", err)
	}
	
	copiedSlice, ok := copied.([]int)
	if !ok {
		t.Fatalf("DeepCopy returned wrong type: %T", copied)
	}
	
	if !reflect.DeepEqual(original, copiedSlice) {
		t.Errorf("DeepCopy values don't match: original %v, copied %v", original, copiedSlice)
	}
}

func TestValidateStructValid(t *testing.T) {
	valid := TestStruct{
		Name: "Valid",
		Age:  25,
	}
	
	err := ValidateStruct(valid)
	if err != nil {
		t.Errorf("Expected no error for valid struct, got: %v", err)
	}
}

func TestValidateStructInvalid(t *testing.T) {
	invalid := TestStruct{
		Name: "", // required but empty
		Age:  25,
	}
	
	err := ValidateStruct(invalid)
	if err == nil {
		t.Error("Expected error for invalid struct")
	}
}

func TestValidateStructMultipleErrors(t *testing.T) {
	invalid := TestStruct{
		Name: "", // required but empty
		Age:  0,  // required but zero value
	}
	
	err := ValidateStruct(invalid)
	if err == nil {
		t.Error("Expected error for invalid struct")
	}
}

func TestCustomStringString(t *testing.T) {
	cs := CustomString("hello")
	result := cs.String()
	if result != "hello" {
		t.Errorf("Expected 'hello', got '%s'", result)
	}
}

func TestCustomStringUpper(t *testing.T) {
	cs := CustomString("hello world")
	result := cs.Upper()
	expected := CustomString("HELLO WORLD")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestCustomStringReverse(t *testing.T) {
	tests := []struct {
		input    CustomString
		expected CustomString
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"a", "a"},
		{"abcd", "dcba"},
	}
	
	for _, test := range tests {
		result := test.input.Reverse()
		if result != test.expected {
			t.Errorf("Reverse(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestCustomStringLength(t *testing.T) {
	tests := []struct {
		input    CustomString
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"a", 1},
		{"hello world", 11},
	}
	
	for _, test := range tests {
		result := test.input.Length()
		if result != test.expected {
			t.Errorf("Length(%s) = %d, expected %d", test.input, result, test.expected)
		}
	}
}