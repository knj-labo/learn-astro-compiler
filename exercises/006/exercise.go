package main

import (
	"fmt"
	"io"
	"os"
)

/*
Exercise 006: ファイル操作とエラーハンドリング

このエクササイズでは、Goでのファイル操作とエラーハンドリングを学びます：

1. WriteToFile 関数を実装する
   - ファイルにテキストを書き込む
   - ファイルが存在しない場合は作成する
   - エラーが発生した場合は適切に処理する

2. ReadFromFile 関数を実装する
   - ファイルからテキストを読み取る
   - ファイルが存在しない場合はエラーを返す
   - 読み取った内容を文字列で返す

3. CopyFile 関数を実装する
   - ソースファイルからデスティネーションファイルにコピーする
   - io.Copy を使用する
   - エラーハンドリングを適切に行う

期待される動作:
- WriteToFile("test.txt", "Hello, World!") でファイルに書き込み
- ReadFromFile("test.txt") でファイルから読み取り
- CopyFile("test.txt", "copy.txt") でファイルをコピー
*/

func main() {
	fmt.Println("Exercise 006: File Operations")
	
	// テスト用のファイル操作
	filename := "test.txt"
	content := "Hello, World!\nThis is a test file."
	
	// ファイルに書き込み
	err := WriteToFile(filename, content)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Printf("Successfully wrote to %s\n", filename)
	
	// ファイルから読み取り
	readContent, err := ReadFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
		return
	}
	fmt.Printf("Read from file: %s\n", readContent)
	
	// ファイルをコピー
	copyFilename := "copy.txt"
	err = CopyFile(filename, copyFilename)
	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}
	fmt.Printf("Successfully copied %s to %s\n", filename, copyFilename)
	
	// 掃除
	os.Remove(filename)
	os.Remove(copyFilename)
}

// WriteToFile 関数の実装
func WriteToFile(filename, content string) error {
	// 1. os.Create() でファイルを作成
	file, err := os.Create(filename)
	if err != nil {
	    return fmt.Errorf("failed to create file: %w", err)
    }

	// 2. defer file.Close() でファイルを確実に閉じる
    defer file.Close()

	// 3. file.WriteString() で内容を書き込む
	_, err = file.WriteString(content)
	if err != nil {
	    return fmt.Errorf("failed to write to file: %w", err)
    }

    // 4. エラーがなければ nil を返す
    return nil
}

// ReadFromFile 関数の実装
func ReadFromFile(filename string) (string, error) {
	// 1. os.Open() でファイルを開く
	file, err := os.Open(filename)
	if err != nil {
	    return "", fmt.Errorf("failed to open file: %w", err)
    }

	// 2. defer file.Close() でファイルを確実に閉じる
	defer file.Close()

	// 3. io.ReadAll() で全内容を読み取る
	data, err := io.ReadAll(file)
	if err != nil {
	    return "", fmt.Errorf("failed to read file: %w", err)
	}

	// 4. []byte を string に変換して返す
	return string(data), nil
}

// CopyFile 関数の実装
func CopyFile(src, dst string) error {
	// 1. os.Open() でソースファイルを開く
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()
	
	// 2. os.Create() でデスティネーションファイルを作成
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}

	// 3. defer で両方のファイルを閉じる
	defer dstFile.Close()
	

	// 4. io.Copy() でファイルをコピー
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	return nil
}