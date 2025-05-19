package main

/*
Exercise:
1の階乗から指定された数字の階乗までの数字を計算して、カンマ区切りで1行に表示するプログラムを実装してください。
*/

import (
	"errors" // エラーを扱うための標準ライブラリ
	"fmt"    // 入出力（表示など）を扱うパッケージ
	"log"    // エラーログ出力用
)

func main() {
	fmt.Println("Exercise 002") // ラベル表示（見た目の識別用）

	// ユーザー入力用の変数を用意
	var input int

	// ユーザーに入力を促すメッセージ
	fmt.Print("Please enter a number : ")

	// 標準入力から数値を読み込む（&input に格納）
	_, err := fmt.Scanln(&input)

	// 入力が数値でないなどのエラーをチェック
	if err != nil {
		log.Fatal("Please enter a number") // エラーがあれば即終了
	}

	// 入力値を使って Exercise002 を実行
	result, err := Exercise002(input)

	// 入力値が 0 以下などでエラーが返ってきたとき
	if err != nil {
		log.Fatalf("Error for input %v: %v", input, err) // エラー内容を表示して終了
	}

	// 計算結果を出力
	fmt.Printf("Factorial of %d = %d", input, result)
}

// Exercise002 は、指定された数の階乗を返す関数
func Exercise002(input int) (int, error) {
	if input < 1 {
		// 負の数や 0 のときはエラーを返す
		return 0, errors.New("input must be a positive integer")
	}

	// 階乗を計算するための変数（初期値 1）
	result := 1

	// 1 から input まで掛け算していく
	for i := 1; i <= input; i++ {
		result *= i
	}

	// 結果を返す（エラーなし）
	return result, nil
}
