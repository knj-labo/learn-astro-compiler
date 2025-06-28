package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
Exercise 014: 並行性とワーカープール

このエクササイズでは、Goの並行性機能を使ったワーカープールパターンを学びます：

1. ワーカープールの実装
   - 複数のgoroutineでタスクを並列処理
   - チャネルを使った通信
   - ワーカーの制御と同期

2. タスクスケジューリング
   - ジョブキューの管理
   - 優先度付きタスク処理
   - 負荷分散

3. 結果の集約
   - ワーカーからの結果収集
   - エラーハンドリング
   - 統計情報の計算

期待される動作:
- 効率的な並列処理
- リソースの適切な管理
- 結果の正確な集約
*/

func main() {
	fmt.Println("Exercise 014: Concurrency and Worker Pools")
	
	// ワーカープールマネージャーを初期化
	manager := NewWorkerPoolManager(5, 100) // 5ワーカー、キューサイズ100
	
	// ワーカープールを開始
	fmt.Println("=== Starting Worker Pool ===")
	manager.Start()
	defer manager.Stop()
	
	// 数値処理タスクのデモ
	fmt.Println("\n=== Number Processing Tasks ===")
	numbers := []int{10, 15, 20, 25, 30, 35, 40, 45, 50}
	
	// 平方計算タスクを送信
	for _, num := range numbers {
		task := Task{
			ID:       fmt.Sprintf("square-%d", num),
			Type:     "square",
			Data:     num,
			Priority: 1,
		}
		manager.SubmitTask(task)
	}
	
	// 素数チェックタスクを送信
	primeNumbers := []int{17, 29, 31, 37, 41, 43, 47, 53, 59, 61}
	for _, num := range primeNumbers {
		task := Task{
			ID:       fmt.Sprintf("prime-%d", num),
			Type:     "prime",
			Data:     num,
			Priority: 2, // 高優先度
		}
		manager.SubmitTask(task)
	}
	
	// ファイル処理タスクのデモ
	fmt.Println("\n=== File Processing Tasks ===")
	fileSizes := []int{1024, 2048, 4096, 8192, 16384}
	for i, size := range fileSizes {
		task := Task{
			ID:       fmt.Sprintf("file-%d", i),
			Type:     "file",
			Data:     size,
			Priority: 1,
		}
		manager.SubmitTask(task)
	}
	
	// ネットワークタスクのデモ
	fmt.Println("\n=== Network Tasks ===")
	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/3",
	}
	for i, url := range urls {
		task := Task{
			ID:       fmt.Sprintf("network-%d", i),
			Type:     "network",
			Data:     url,
			Priority: 3, // 最高優先度
		}
		manager.SubmitTask(task)
	}
	
	// 処理完了を待つ
	fmt.Println("\n=== Waiting for completion ===")
	time.Sleep(10 * time.Second)
	
	// 統計情報を表示
	fmt.Println("\n=== Statistics ===")
	stats := manager.GetStatistics()
	fmt.Printf("Total tasks submitted: %d\n", stats.TotalSubmitted)
	fmt.Printf("Total tasks completed: %d\n", stats.TotalCompleted)
	fmt.Printf("Total tasks failed: %d\n", stats.TotalFailed)
	fmt.Printf("Average processing time: %.2fms\n", stats.AvgProcessingTime)
	fmt.Printf("Active workers: %d\n", stats.ActiveWorkers)
	fmt.Printf("Queue size: %d\n", stats.QueueSize)
	
	// バッチ処理のデモ
	fmt.Println("\n=== Batch Processing ===")
	batchProcessor := NewBatchProcessor(3, 50) // 3ワーカー、バッチサイズ50
	
	// 大量のタスクを生成
	var batchTasks []BatchItem
	for i := 0; i < 200; i++ {
		item := BatchItem{
			ID:   i,
			Data: rand.Intn(1000),
		}
		batchTasks = append(batchTasks, item)
	}
	
	// バッチ処理を実行
	results := batchProcessor.ProcessBatch(batchTasks)
	fmt.Printf("Processed %d items in batch\n", len(results))
	
	// 結果の統計
	var sum int
	for _, result := range results {
		sum += result.Result
	}
	fmt.Printf("Sum of results: %d\n", sum)
	fmt.Printf("Average result: %.2f\n", float64(sum)/float64(len(results)))
}

// Task構造体
type Task struct {
	ID       string      `json:"id"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
	Priority int         `json:"priority"`
	SubmitTime time.Time `json:"submit_time"`
}

// TaskResult構造体
type TaskResult struct {
	TaskID      string        `json:"task_id"`
	Success     bool          `json:"success"`
	Result      interface{}   `json:"result"`
	Error       string        `json:"error,omitempty"`
	ProcessTime time.Duration `json:"process_time"`
}

// WorkerStatistics構造体
type WorkerStatistics struct {
	TotalSubmitted     int     `json:"total_submitted"`
	TotalCompleted     int     `json:"total_completed"`
	TotalFailed        int     `json:"total_failed"`
	AvgProcessingTime  float64 `json:"avg_processing_time_ms"`
	ActiveWorkers      int     `json:"active_workers"`
	QueueSize          int     `json:"queue_size"`
}

// WorkerPoolManager構造体
type WorkerPoolManager struct {
	workerCount int
	taskQueue   chan Task
	resultQueue chan TaskResult
	workers     []*Worker
	statistics  *WorkerStatistics
	mutex       sync.RWMutex
	stopChan    chan bool
	wg          sync.WaitGroup
}

// Worker構造体
type Worker struct {
	ID          int
	taskQueue   chan Task
	resultQueue chan TaskResult
	stopChan    chan bool
	active      bool
	mutex       sync.RWMutex
}

// BatchItem構造体
type BatchItem struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
}

// BatchResult構造体
type BatchResult struct {
	ID     int         `json:"id"`
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

// BatchProcessor構造体
type BatchProcessor struct {
	workerCount int
	batchSize   int
}

// NewWorkerPoolManager関数の実装
func NewWorkerPoolManager(workerCount, queueSize int) *WorkerPoolManager {
	// TODO: 実装する
	// ヒント:
	// 1. WorkerPoolManager構造体を初期化
	// 2. タスクキューと結果キューを作成
	// 3. 統計情報を初期化
	// 4. ワーカーを作成（まだ開始しない）
	return nil
}

// Start メソッドの実装
func (wpm *WorkerPoolManager) Start() {
	// TODO: 実装する
	// ヒント:
	// 1. 各ワーカーを開始
	// 2. 結果処理のgoroutineを開始
	// 3. ワーカーグループに追加
}

// Stop メソッドの実装
func (wpm *WorkerPoolManager) Stop() {
	// TODO: 実装する
	// ヒント:
	// 1. stopChanを使ってワーカーに停止信号を送信
	// 2. 全ワーカーの完了を待つ
	// 3. チャネルを閉じる
}

// SubmitTask メソッドの実装
func (wpm *WorkerPoolManager) SubmitTask(task Task) {
	// TODO: 実装する
	// ヒント:
	// 1. タスクに送信時刻を設定
	// 2. 優先度に基づいてタスクをキューに追加
	// 3. 統計情報を更新
}

// GetStatistics メソッドの実装
func (wpm *WorkerPoolManager) GetStatistics() WorkerStatistics {
	// TODO: 実装する
	// ヒント:
	// 1. 読み取りロックを取得
	// 2. 現在の統計情報をコピーして返す
	return WorkerStatistics{}
}

// NewWorker関数の実装
func NewWorker(id int, taskQueue chan Task, resultQueue chan TaskResult, stopChan chan bool) *Worker {
	// TODO: 実装する
	// ヒント:
	// 1. Worker構造体を初期化
	// 2. チャネルを設定
	return nil
}

// Start メソッド（Worker）の実装
func (w *Worker) Start() {
	// TODO: 実装する
	// ヒント:
	// 1. goroutineでワーカーのメインループを開始
	// 2. タスクキューからタスクを受信
	// 3. タスクを処理して結果を送信
	// 4. 停止信号をチェック
}

// processTask メソッドの実装
func (w *Worker) processTask(task Task) TaskResult {
	// TODO: 実装する
	// ヒント:
	// 1. タスクタイプに基づいて適切な処理を実行
	// 2. 処理時間を測定
	// 3. エラーハンドリングを実装
	// 4. TaskResult構造体を返す
	return TaskResult{}
}

// NewBatchProcessor関数の実装
func NewBatchProcessor(workerCount, batchSize int) *BatchProcessor {
	// TODO: 実装する
	// ヒント:
	// 1. BatchProcessor構造体を初期化
	return nil
}

// ProcessBatch メソッドの実装
func (bp *BatchProcessor) ProcessBatch(items []BatchItem) []BatchResult {
	// TODO: 実装する
	// ヒント:
	// 1. アイテムをバッチに分割
	// 2. 各バッチを並列処理
	// 3. 結果を収集して返す
	// 4. sync.WaitGroupを使用して同期
	return nil
}

// processBatchItems ヘルパー関数の実装
func processBatchItems(items []BatchItem, results chan<- BatchResult, wg *sync.WaitGroup) {
	// TODO: 実装する
	// ヒント:
	// 1. defer wg.Done()を設定
	// 2. 各アイテムを処理
	// 3. 結果をチャネルに送信
}

// calculateSquare ヘルパー関数の実装
func calculateSquare(n int) int {
	// TODO: 実装する
	// ヒント:
	// 1. 平方を計算
	// 2. 処理時間をシミュレート
	return 0
}

// isPrime ヘルパー関数の実装
func isPrime(n int) bool {
	// TODO: 実装する
	// ヒント:
	// 1. 素数判定のアルゴリズムを実装
	// 2. 2からsqrt(n)まで除算チェック
	return false
}

// simulateFileProcessing ヘルパー関数の実装
func simulateFileProcessing(size int) string {
	// TODO: 実装する
	// ヒント:
	// 1. ファイル処理をシミュレート
	// 2. ランダムな処理時間を追加
	// 3. 処理結果を文字列で返す
	return ""
}

// simulateNetworkRequest ヘルパー関数の実装
func simulateNetworkRequest(url string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. ネットワークリクエストをシミュレート
	// 2. time.Sleep()で遅延を追加
	// 3. ランダムにエラーを発生させる
	return "", nil
}