package main

import (
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	filename := "test_write.txt"
	content := "Hello, World!"
	
	// テスト後にファイルを削除
	defer os.Remove(filename)
	
	err := WriteToFile(filename, content)
	if err != nil {
		t.Fatalf("WriteToFile failed: %v", err)
	}
	
	// ファイルが存在するかチェック
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("File was not created")
	}
	
	// 内容をチェック
	readContent, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	
	if string(readContent) != content {
		t.Errorf("Expected %s, got %s", content, string(readContent))
	}
}

func TestReadFromFile(t *testing.T) {
	filename := "test_read.txt"
	content := "Hello, World!"
	
	// テスト用ファイルを作成
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// テスト後にファイルを削除
	defer os.Remove(filename)
	
	readContent, err := ReadFromFile(filename)
	if err != nil {
		t.Fatalf("ReadFromFile failed: %v", err)
	}
	
	if readContent != content {
		t.Errorf("Expected %s, got %s", content, readContent)
	}
}

func TestReadFromFileNotExist(t *testing.T) {
	_, err := ReadFromFile("nonexistent.txt")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

func TestCopyFile(t *testing.T) {
	srcFile := "test_src.txt"
	dstFile := "test_dst.txt"
	content := "Hello, World!"
	
	// ソースファイルを作成
	err := os.WriteFile(srcFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}
	
	// テスト後にファイルを削除
	defer os.Remove(srcFile)
	defer os.Remove(dstFile)
	
	err = CopyFile(srcFile, dstFile)
	if err != nil {
		t.Fatalf("CopyFile failed: %v", err)
	}
	
	// デスティネーションファイルが存在するかチェック
	if _, err := os.Stat(dstFile); os.IsNotExist(err) {
		t.Error("Destination file was not created")
	}
	
	// 内容をチェック
	readContent, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}
	
	if string(readContent) != content {
		t.Errorf("Expected %s, got %s", content, string(readContent))
	}
}

func TestCopyFileNotExist(t *testing.T) {
	err := CopyFile("nonexistent.txt", "dst.txt")
	if err == nil {
		t.Error("Expected error for non-existent source file")
	}
}