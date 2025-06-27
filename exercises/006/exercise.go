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
	// TODO: 実装する
	// ヒント:
	// 1. os.Create() でファイルを作成
	// 2. defer file.Close() でファイルを確実に閉じる
	// 3. file.WriteString() で内容を書き込む
	return nil
}

// ReadFromFile 関数の実装
func ReadFromFile(filename string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. os.Open() でファイルを開く
	// 2. defer file.Close() でファイルを確実に閉じる
	// 3. io.ReadAll() で全内容を読み取る
	// 4. []byte を string に変換して返す
	return "", nil
}

// CopyFile 関数の実装
func CopyFile(src, dst string) error {
	// TODO: 実装する
	// ヒント:
	// 1. os.Open() でソースファイルを開く
	// 2. os.Create() でデスティネーションファイルを作成
	// 3. defer で両方のファイルを閉じる
	// 4. io.Copy() でファイルをコピー
	return nil
}