package main

import (
	"fmt"
	"sync"
)

/*
Exercise 004: Go言語の並行処理とインターフェースの理解

このエクササイズでは、Goの重要な機能を理解するために以下の実装を行ってください：

1. Worker インターフェースを実装する
   - Process(id int) string メソッドを持つ

2. SimpleWorker 構造体を実装する
   - Worker インターフェースを満たす
   - Process メソッドは "Processed: {id}" という文字列を返す

3. Exercise004 関数を実装する
   - 与えられた数のゴルーチンを起動
   - 各ゴルーチンは Worker を使って処理を実行
   - チャネルを使って結果を収集
   - すべての結果をスライスで返す（順序は問わない）

期待される動作:
- Exercise004(3) を呼ぶと、3つのゴルーチンが並行して実行される
- 各ゴルーチンは異なるIDで Process を呼ぶ
- 結果として ["Processed: 0", "Processed: 1", "Processed: 2"] が返る（順序は異なる可能性あり）
*/

func main() {
	fmt.Println("Exercise 004")
	worker := &SimpleWorker{}
	results := Exercise004(5, worker)
	fmt.Println("Results:", results)
}

// Worker インターフェースの定義
type Worker interface {
	Process(id int) string
}

// SimpleWorker 構造体の定義
type SimpleWorker struct{}

// Process メソッドの実装
func (w *SimpleWorker) Process(id int) string {
	// TODO: "Processed: {id}" という形式の文字列を返す
	return fmt.Sprintf("Processed: %d", id)
}

// Exercise004 関数の実装
func Exercise004(numWorkers int, worker Worker) []string {
	if numWorkers <= 0 {
		return []string{}
	}

	// 1. 結果を収集するためのチャネルを作成
	results := make(chan string, numWorkers)
	
	// 結果を格納するスライス
	var finalResults []string

	// 2. WaitGroupを使って全ゴルーチンの完了を待つ
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// 3. 各ゴルーチンでworker.Process(id)を呼ぶ
	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			defer wg.Done()
			// 4. 結果をチャネルに送信
			result := worker.Process(id)
			results <- result
		}(i)
	}

	// 5. 別のゴルーチンで全ての完了を待ってチャネルを閉じる
	go func() {
		wg.Wait()
		close(results)
	}()

	// 6. チャネルから結果を収集
	for result := range results {
		finalResults = append(finalResults, result)
	}

	return finalResults
}