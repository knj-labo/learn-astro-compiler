# ゴルーチンとチャネルの理解

## 1. ゴルーチン（Goroutine）

### 通常の実行 vs ゴルーチン

```
通常の実行:
タスク1 ----完了----> タスク2 ----完了----> タスク3
         (待つ)              (待つ)

ゴルーチンの実行:
タスク1 ----実行中---->
タスク2 ----実行中---->     （全部同時に実行！）
タスク3 ----実行中---->
```

### 使い方
```go
// 普通の関数呼び出し
doWork()  // この関数が終わるまで待つ

// ゴルーチンで実行
go doWork()  // すぐに次の行に進む（doWorkは裏で実行される）
```

## 2. チャネル（Channel）

チャネルは、ゴルーチン間でデータを安全にやり取りするためのパイプです。

### チャネルの動作イメージ

```
ゴルーチン1                     ゴルーチン2
    |                              |
    |  データ"hello"               |
    |  ----------->  チャネル      |
    |                    |         |
    |                    v         |
    |               [hello]        |
    |                    |         |
    |                    ----------> データ"hello"を受信
```

### 基本的な使い方

```go
// 1. チャネルを作る
ch := make(chan string)

// 2. ゴルーチンAでデータを送信
go func() {
    ch <- "こんにちは"  // チャネルにデータを送る
}()

// 3. メインゴルーチンで受信
message := <-ch  // チャネルからデータを受け取る
fmt.Println(message)  // "こんにちは"
```

## 3. Exercise 004 で必要な実装

### SimpleWorker の Process メソッド
```go
func (w *SimpleWorker) Process(id int) string {
    // "Processed: {id}" という形式の文字列を返す
    return fmt.Sprintf("Processed: %d", id)
}
```

### Exercise004 関数の実装パターン

```go
func Exercise004(numWorkers int, worker Worker) []string {
    // 1. 結果を受け取るチャネルを作成
    results := make(chan string, numWorkers)
    
    // 2. 結果を格納するスライス
    var finalResults []string
    
    // 3. WaitGroupで全ゴルーチンの完了を待つ
    var wg sync.WaitGroup
    wg.Add(numWorkers)
    
    // 4. 指定された数のゴルーチンを起動
    for i := 0; i < numWorkers; i++ {
        go func(id int) {
            defer wg.Done()  // このゴルーチンが終わったことを通知
            
            // ワーカーで処理を実行
            result := worker.Process(id)
            
            // 結果をチャネルに送信
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
```

## ポイント

1. **`go` キーワード**: 関数の前に付けると、その関数は新しいゴルーチンで実行される
2. **チャネル**: `make(chan 型)` で作成し、`<-` で送受信
3. **sync.WaitGroup**: 複数のゴルーチンの完了を待つために使用
4. **defer**: 関数が終了する時に実行される（クリーンアップに便利）

## なぜ並行処理を使うのか？

- 複数のタスクを同時に実行できる
- I/O待ち（ネットワーク、ファイル読み書き）の間に他の処理ができる
- マルチコアCPUを有効活用できる