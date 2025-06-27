package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
Exercise 009: リフレクションとカスタム型

このエクササイズでは、Goのリフレクション機能とカスタム型の活用を学びます：

1. StructInfo 関数を実装する
   - 任意の構造体の情報を取得
   - フィールド名、型、タグを表示
   - reflect パッケージを使用

2. DeepCopy 関数を実装する
   - 任意のデータ構造の深いコピーを作成
   - JSON Marshal/Unmarshal を使用

3. ValidateStruct 関数を実装する
   - 構造体のバリデーションを実行
   - "required" タグを持つフィールドの検証
   - 空の値をチェック

4. CustomString 型を実装する
   - カスタムメソッドを持つ文字列型
   - String(), Upper(), Reverse() メソッド

期待される動作:
- StructInfo(person) で構造体の詳細情報を表示
- DeepCopy(original) で完全なコピーを作成
- ValidateStruct(person) でバリデーション実行
- CustomString("hello").Upper() → "HELLO"
*/

func main() {
	fmt.Println("Exercise 009: Reflection and Custom Types")
	
	// テスト用の構造体
	type Person struct {
		Name  string `json:"name" required:"true"`
		Age   int    `json:"age" required:"true"`
		Email string `json:"email"`
		City  string `json:"city" required:"true"`
	}
	
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		City:  "Tokyo",
	}
	
	// 構造体情報の表示
	fmt.Println("=== Struct Info ===")
	StructInfo(person)
	
	// 深いコピーのテスト
	fmt.Println("\n=== Deep Copy ===")
	copyPerson, err := DeepCopy(person)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Original: %+v\n", person)
		fmt.Printf("Copy: %+v\n", copyPerson)
	}
	
	// バリデーションのテスト
	fmt.Println("\n=== Validation ===")
	validPerson := Person{Name: "Bob", Age: 25, City: "Osaka"}
	invalidPerson := Person{Name: "", Age: 25, City: ""}
	
	fmt.Printf("Valid person validation: %v\n", ValidateStruct(validPerson))
	fmt.Printf("Invalid person validation: %v\n", ValidateStruct(invalidPerson))
	
	// カスタム文字列型のテスト
	fmt.Println("\n=== Custom String ===")
	cs := CustomString("hello world")
	fmt.Printf("Original: %s\n", cs)
	fmt.Printf("Upper: %s\n", cs.Upper())
	fmt.Printf("Reverse: %s\n", cs.Reverse())
	fmt.Printf("Length: %d\n", cs.Length())
}

// StructInfo 関数の実装
func StructInfo(v interface{}) {
	// TODO: 実装する
	// ヒント:
	// 1. reflect.ValueOf() と reflect.TypeOf() を使用
	// 2. v.NumField() でフィールド数を取得
	// 3. v.Field(i) と t.Field(i) でフィールド情報を取得
	// 4. フィールド名、型、タグを表示
}

// DeepCopy 関数の実装
func DeepCopy(src interface{}) (interface{}, error) {
	// TODO: 実装する
	// ヒント:
	// 1. json.Marshal() でシリアライズ
	// 2. reflect.TypeOf() で元の型を取得
	// 3. reflect.New() で新しいインスタンスを作成
	// 4. json.Unmarshal() でデシリアライズ
	// 5. reflect.Indirect() で値を取得
	return nil, nil
}

// ValidateStruct 関数の実装
func ValidateStruct(v interface{}) error {
	// TODO: 実装する
	// ヒント:
	// 1. reflect.ValueOf() と reflect.TypeOf() を使用
	// 2. フィールドをループで確認
	// 3. "required" タグをチェック
	// 4. フィールドが空かどうかチェック
	// 5. エラーがあれば適切なメッセージを返す
	return nil
}

// CustomString カスタム型の定義
type CustomString string

// String メソッドの実装
func (cs CustomString) String() string {
	// TODO: 実装する
	return ""
}

// Upper メソッドの実装
func (cs CustomString) Upper() CustomString {
	// TODO: 実装する
	// ヒント: strings.ToUpper() を使用
	return ""
}

// Reverse メソッドの実装
func (cs CustomString) Reverse() CustomString {
	// TODO: 実装する
	// ヒント: 
	// 1. []rune に変換
	// 2. 前後を入れ替え
	// 3. string に戻す
	return ""
}

// Length メソッドの実装
func (cs CustomString) Length() int {
	// TODO: 実装する
	return 0
}