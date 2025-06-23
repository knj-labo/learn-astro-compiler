package main

import (
	"fmt"
	"time"
)

func main() {
	// 例1: ゴルーチンの基本
	fmt.Println("=== ゴルーチンの基本 ===")
	
	// 通常の関数呼び出し（順番に実行される）
	fmt.Println("順番に実行:")
	printSlowly("こんにちは")
	printSlowly("世界")
	
	fmt.Println("\n並行に実行:")
	// ゴルーチンで実行（並行して実行される）
	go printSlowly("こんにちは")
	go printSlowly("世界")
	
	// メインが終了しないように少し待つ
	time.Sleep(2 * time.Second)
	
	// 例2: チャネルの基本
	fmt.Println("\n\n=== チャネルの基本 ===")
	
	// チャネルを作成
	messages := make(chan string)
	
	// ゴルーチンでメッセージを送信
	go func() {
		time.Sleep(1 * time.Second)
		messages <- "チャネルから送られたメッセージ"
	}()
	
	// メッセージを受信（受信するまでブロックされる）
	msg := <-messages
	fmt.Println("受信:", msg)
	
	// 例3: 複数のゴルーチンとチャネル
	fmt.Println("\n=== 複数のゴルーチンとチャネル ===")
	
	results := make(chan int)
	
	// 3つの計算を並行で実行
	go calculate(2, 3, results)
	go calculate(5, 5, results)
	go calculate(10, 10, results)
	
	// 3つの結果を受信
	for i := 0; i < 3; i++ {
		result := <-results
		fmt.Printf("計算結果 %d: %d\n", i+1, result)
	}
	
	// 例4: バッファ付きチャネル
	fmt.Println("\n=== バッファ付きチャネル ===")
	
	// バッファサイズ2のチャネル
	buffered := make(chan string, 2)
	
	// ブロックされずに2つまで送信できる
	buffered <- "最初"
	buffered <- "次"
	
	// 受信
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
}

func printSlowly(text string) {
	for _, char := range text {
		fmt.Printf("%c ", char)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()
}

func calculate(a, b int, results chan int) {
	// 計算に時間がかかるふり
	time.Sleep(500 * time.Millisecond)
	results <- a * b
}