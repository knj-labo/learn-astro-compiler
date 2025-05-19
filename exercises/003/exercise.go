package main

import "fmt"

/*
Exercise 003
指定された整数 n に対して、次の条件を満たすマップ（map）を生成するプログラムを実装してください：

(i, i*i) の形で、1 から n（両端を含む）までの整数 i に対して、
キーが i、値が i の二乗となるマップを生成し、
そのマップの内容（値の表現を含む）を出力してください。

たとえば、次の入力が与えられたとします：8
その場合、出力は次のようになります：
map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]
*/

func main() {
	fmt.Println("Exercise 003")
	fmt.Println(Exercise003(8))
}

func Exercise003(n int) map[int]int {
	return nil
}
