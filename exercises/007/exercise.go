package main

import (
	"fmt"
	"sort"
)

/*
Exercise 007: スライス操作とソート

このエクササイズでは、Goでのスライス操作とソートアルゴリズムを学びます：

1. FilterEven 関数を実装する
   - 整数スライスから偶数のみを抽出する
   - 新しいスライスを返す

2. MapSquare 関数を実装する
   - 整数スライスの各要素を二乗する
   - 新しいスライスを返す

3. SortByLength 関数を実装する
   - 文字列スライスを文字列の長さでソートする
   - sort.Slice を使用する

4. RemoveDuplicates 関数を実装する
   - 整数スライスから重複を削除する
   - mapを使って重複をチェック

期待される動作:
- FilterEven([1,2,3,4,5,6]) → [2,4,6]
- MapSquare([1,2,3,4]) → [1,4,9,16]
- SortByLength(["apple","go","banana","hi"]) → ["go","hi","apple","banana"]
- RemoveDuplicates([1,2,2,3,3,3,4]) → [1,2,3,4]
*/

func main() {
	fmt.Println("Exercise 007: Slice Operations and Sorting")
	
	// テスト用データ
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	words := []string{"apple", "go", "banana", "hi", "programming"}
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 5, 4, 1}
	
	// 偶数フィルター
	evens := FilterEven(numbers)
	fmt.Printf("Even numbers: %v\n", evens)
	
	// 二乗マップ
	squares := MapSquare([]int{1, 2, 3, 4, 5})
	fmt.Printf("Squares: %v\n", squares)
	
	// 長さでソート
	sortedWords := SortByLength(words)
	fmt.Printf("Sorted by length: %v\n", sortedWords)
	
	// 重複削除
	unique := RemoveDuplicates(duplicates)
	fmt.Printf("Unique numbers: %v\n", unique)
}

// FilterEven 関数の実装
func FilterEven(numbers []int) []int {
	// 1. 結果用の空のスライスを作成
	result := []int{}

	// 2. range でスライスをループ
	for _, n := range numbers {

	// 3. 偶数の場合（n % 2 == 0）は結果に追加
	    if n%2 == 0 {
	                result = append(result, n)
        }
    }
	return result
}

// MapSquare 関数の実装
func MapSquare(numbers []int) []int {
	// 1. 同じ長さの結果用スライスを作成：make([]int, len(numbers))
	result := make([]int, len(numbers))

	// 2. range でスライスをループ
	for i, n := range numbers {

	// 3. 各要素を二乗して結果に格納
	    result[i] = n * n
    }
	return result
}

// SortByLength 関数の実装
func SortByLength(words []string) []string {
	// 1. 元のスライスをコピー：make([]string, len(words)) + copy()
	result := make([]string, len(words))
	copy(result, words)

	// 2. sort.Slice() を使用
	sort.Slice(result, func(i, j int) bool {

	// 3. 比較関数で len(words[i]) < len(words[j]) を使用
	    return len(result[i]) < len(result[j])
    })
	return result
}

// RemoveDuplicates 関数の実装
func RemoveDuplicates(numbers []int) []int {
	// 1. map[int]bool を作成して重複チェック用
	seen := make(map[int]bool)

	// 2. 結果用の空のスライスを作成
	result := []int{}

	// 3. range でスライスをループ
	for _, n := range numbers {

		// 4. mapに存在しない場合のみ結果に追加し、mapにマーク
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}
	return result
}