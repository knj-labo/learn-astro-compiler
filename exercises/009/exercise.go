package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	// 1. reflect.ValueOf() と reflect.TypeOf() を使用
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	
	// ポインタの場合は実際の値を取得
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	// 構造体でない場合は処理しない
	if val.Kind() != reflect.Struct {
		fmt.Printf("Type: %s (not a struct)\n", typ.Name())
		return
	}
	
	fmt.Printf("Type: %s\n", typ.Name())
	fmt.Printf("Fields: %d\n", val.NumField())
	
	// 2. v.NumField() でフィールド数を取得
	for i := 0; i < val.NumField(); i++ {
	    // 3. v.Field(i) と t.Field(i) でフィールド情報を取得
		field := typ.Field(i)
		value := val.Field(i)

	    // 4. フィールド名、型、タグを表示
		fmt.Printf("  Field %d: %s\n", i+1, field.Name)
		fmt.Printf("    Type: %s\n", field.Type)
		fmt.Printf("    Value: %v\n", value.Interface())
		
		// タグ情報を表示
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			fmt.Printf("    JSON Tag: %s\n", jsonTag)
		}
		if requiredTag := field.Tag.Get("required"); requiredTag != "" {
			fmt.Printf("    Required: %s\n", requiredTag)
		}
	}
}

// DeepCopy 関数の実装
func DeepCopy(src interface{}) (interface{}, error) {
	// 1. JSONにシリアライズ
	data, err := json.Marshal(src)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %v", err)
	}
	
	// 2. 元の型を取得
	srcType := reflect.TypeOf(src)
	
	// 3. 新しいインスタンスを作成
	dst := reflect.New(srcType)
	
	// 4. JSONからデシリアライズ
	err = json.Unmarshal(data, dst.Interface())
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}
	
	// 5. 実際の値を返す（ポインタではなく）
	return reflect.Indirect(dst).Interface(), nil
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
	
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	
	// ポインタの場合は実際の値を取得
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	// 構造体でない場合はエラー
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("value is not a struct")
	}
	
	var errors []string
	
	// 各フィールドをチェック
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)
		
		// "required" タグをチェック
		if requiredTag := field.Tag.Get("required"); requiredTag == "true" {
			// フィールドが空かどうかチェック
			if isZeroValue(fieldValue) {
				errors = append(errors, fmt.Sprintf("field '%s' is required but empty", field.Name))
			}
		}
	}
	
	// エラーがある場合は結合して返す
	if len(errors) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(errors, ", "))
	}
	
	return nil
}

// isZeroValue はreflect.Valueがゼロ値かどうかをチェックする
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return v.IsZero()
	}
}

// CustomString カスタム型の定義
type CustomString string

// String メソッドの実装
func (cs CustomString) String() string {
	return string(cs)
}

// Upper メソッドの実装
func (cs CustomString) Upper() CustomString {
	return CustomString(strings.ToUpper(string(cs)))
}

// Reverse メソッドの実装
func (cs CustomString) Reverse() CustomString {
	// 1. []rune に変換
	runes := []rune(cs)
	
	// 2. 前後を入れ替え
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	// 3. string に戻す
	return CustomString(string(runes))
}

// Length メソッドの実装
func (cs CustomString) Length() int {
	return len(string(cs))
}