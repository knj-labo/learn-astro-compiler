package main

/*
Exercise:
2000から3200までの数字の中で、
- 5では割り切れない
- 7で割り切れる
という条件に当てはまる数字を探し出して、カンマ区切りで1行に表示するプログラムを実装してください。
*/

import (
	// 文字を表示する
	"fmt"
	// 文字列を数値に変換する
	"strconv"
)

func main() {
	// "Exercise 001"を表示する
	fmt.Println("Exercise 001")
	// Ex001を呼び出す
	res := Ex001(2000, 3200)
	// 結果を表示する
	fmt.Println(res)
}

func Ex001(lowest_number, highest_number int) string {
	// 結果を保持する
	var result string
	for i := lowest_number; i <= highest_number; i++ {
		// 5では割り切れない, 7で割り切れる
		if i%7 == 0 && i%5 != 0 {
			// 結果が空でないとき、カンマを追加する
			if result != "" {
				result += ","
			}
			// 結果をカンマ区切りで連結する
			result += strconv.Itoa(i)
		}
	}
	// 結果を返す
	return result
}
