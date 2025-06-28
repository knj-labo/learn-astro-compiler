package main

import (
	"fmt"
	"log"
	"os"
)

/*
Exercise 012: ファイルI/OとCSV処理

このエクササイズでは、ファイル操作とCSVデータの処理を学びます：

1. ファイル読み書き操作
   - テキストファイルの読み取り
   - ファイルへの書き込み
   - ファイル存在チェック

2. CSV処理
   - CSVファイルの読み取り
   - データの解析と構造化
   - CSVファイルの書き込み

3. エラーハンドリング
   - ファイル操作のエラー処理
   - データ検証

期待される動作:
- CSVファイルからの顧客データ読み取り
- データの処理と変換
- 結果の新しいCSVファイルへの保存
*/

func main() {
	fmt.Println("Exercise 012: File I/O and CSV Processing")
	
	// CSVプロセッサーを初期化
	processor := NewCSVProcessor()
	
	// サンプルデータを作成
	fmt.Println("=== Creating Sample Data ===")
	sampleData := []Customer{
		{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", Age: 30, City: "Tokyo"},
		{ID: 2, Name: "Bob Smith", Email: "bob@example.com", Age: 25, City: "Osaka"},
		{ID: 3, Name: "Carol Brown", Email: "carol@example.com", Age: 35, City: "Kyoto"},
		{ID: 4, Name: "David Wilson", Email: "david@example.com", Age: 28, City: "Tokyo"},
	}
	
	// サンプルCSVファイルを作成
	inputFile := "customers.csv"
	err := processor.WriteCustomersToCSV(inputFile, sampleData)
	if err != nil {
		log.Fatalf("Error writing sample data: %v", err)
	}
	fmt.Printf("Sample data written to %s\n", inputFile)
	
	// CSVファイルから顧客データを読み取り
	fmt.Println("\n=== Reading Customer Data ===")
	customers, err := processor.ReadCustomersFromCSV(inputFile)
	if err != nil {
		log.Fatalf("Error reading customers: %v", err)
	}
	
	fmt.Printf("Read %d customers:\n", len(customers))
	for _, customer := range customers {
		fmt.Printf("  %+v\n", customer)
	}
	
	// データ処理
	fmt.Println("\n=== Processing Data ===")
	
	// 年齢でフィルタリング（30歳以上）
	adults := processor.FilterByAge(customers, 30)
	fmt.Printf("Customers aged 30 or above: %d\n", len(adults))
	
	// 都市でグループ化
	cityGroups := processor.GroupByCity(customers)
	fmt.Println("Customers grouped by city:")
	for city, cityCustomers := range cityGroups {
		fmt.Printf("  %s: %d customers\n", city, len(cityCustomers))
	}
	
	// 統計情報
	stats := processor.CalculateStats(customers)
	fmt.Printf("Statistics: %+v\n", stats)
	
	// 処理結果を新しいCSVファイルに保存
	fmt.Println("\n=== Saving Processed Data ===")
	outputFile := "adults.csv"
	err = processor.WriteCustomersToCSV(outputFile, adults)
	if err != nil {
		log.Printf("Error writing processed data: %v", err)
	} else {
		fmt.Printf("Processed data written to %s\n", outputFile)
	}
	
	// レポートファイルを生成
	reportFile := "customer_report.txt"
	err = processor.GenerateReport(reportFile, customers, stats)
	if err != nil {
		log.Printf("Error generating report: %v", err)
	} else {
		fmt.Printf("Report generated: %s\n", reportFile)
	}
	
	// ファイルサイズ情報
	fmt.Println("\n=== File Information ===")
	files := []string{inputFile, outputFile, reportFile}
	for _, file := range files {
		size, err := processor.GetFileSize(file)
		if err != nil {
			fmt.Printf("%s: Error - %v\n", file, err)
		} else {
			fmt.Printf("%s: %d bytes\n", file, size)
		}
	}
	
	// クリーンアップ
	fmt.Println("\n=== Cleanup ===")
	for _, file := range files {
		if processor.FileExists(file) {
			os.Remove(file)
			fmt.Printf("Removed %s\n", file)
		}
	}
}

// Customer構造体
type Customer struct {
	ID    int    `csv:"id"`
	Name  string `csv:"name"`
	Email string `csv:"email"`
	Age   int    `csv:"age"`
	City  string `csv:"city"`
}

// CustomerStats構造体
type CustomerStats struct {
	TotalCustomers int     `json:"total_customers"`
	AverageAge     float64 `json:"average_age"`
	MinAge         int     `json:"min_age"`
	MaxAge         int     `json:"max_age"`
	CitiesCount    int     `json:"cities_count"`
}

// CSVProcessor構造体
type CSVProcessor struct{}

// NewCSVProcessor関数の実装
func NewCSVProcessor() *CSVProcessor {
	// TODO: 実装する
	// ヒント:
	// 1. CSVProcessor構造体を初期化して返す
	return nil
}

// ReadCustomersFromCSV メソッドの実装
func (cp *CSVProcessor) ReadCustomersFromCSV(filename string) ([]Customer, error) {
	// TODO: 実装する
	// ヒント:
	// 1. encoding/csvパッケージを使用
	// 2. ファイルを開いてCSVリーダーを作成
	// 3. ヘッダー行を読み取り
	// 4. 各行をCustomer構造体に変換
	// 5. エラーハンドリングを適切に実装
	return nil, nil
}

// WriteCustomersToCSV メソッドの実装
func (cp *CSVProcessor) WriteCustomersToCSV(filename string, customers []Customer) error {
	// TODO: 実装する
	// ヒント:
	// 1. ファイルを作成
	// 2. CSVライターを作成
	// 3. ヘッダー行を書き込み
	// 4. 各Customer構造体をCSV行に変換して書き込み
	// 5. ファイルを適切にクローズ
	return nil
}

// FilterByAge メソッドの実装
func (cp *CSVProcessor) FilterByAge(customers []Customer, minAge int) []Customer {
	// TODO: 実装する
	// ヒント:
	// 1. 新しいスライスを作成
	// 2. 各顧客の年齢をチェック
	// 3. 条件を満たす顧客のみを新しいスライスに追加
	return nil
}

// GroupByCity メソッドの実装
func (cp *CSVProcessor) GroupByCity(customers []Customer) map[string][]Customer {
	// TODO: 実装する
	// ヒント:
	// 1. マップを作成（キー: 都市名、値: Customer スライス）
	// 2. 各顧客を対応する都市のスライスに追加
	return nil
}

// CalculateStats メソッドの実装
func (cp *CSVProcessor) CalculateStats(customers []Customer) CustomerStats {
	// TODO: 実装する
	// ヒント:
	// 1. 顧客数をカウント
	// 2. 年齢の合計、最小値、最大値を計算
	// 3. 平均年齢を計算
	// 4. ユニークな都市数をカウント
	// 5. CustomerStats構造体を作成して返す
	return CustomerStats{}
}

// GenerateReport メソッドの実装
func (cp *CSVProcessor) GenerateReport(filename string, customers []Customer, stats CustomerStats) error {
	// TODO: 実装する
	// ヒント:
	// 1. ファイルを作成
	// 2. レポートのヘッダーを書き込み
	// 3. 統計情報を書き込み
	// 4. 顧客一覧を書き込み
	// 5. ファイルを適切にクローズ
	return nil
}

// FileExists メソッドの実装
func (cp *CSVProcessor) FileExists(filename string) bool {
	// TODO: 実装する
	// ヒント:
	// 1. os.Stat()を使用してファイル情報を取得
	// 2. エラーをチェックしてファイルの存在を判定
	return false
}

// GetFileSize メソッドの実装
func (cp *CSVProcessor) GetFileSize(filename string) (int64, error) {
	// TODO: 実装する
	// ヒント:
	// 1. os.Stat()を使用してファイル情報を取得
	// 2. FileInfo.Size()でファイルサイズを取得
	// 3. エラーハンドリングを実装
	return 0, nil
}

// ReadTextFile ヘルパー関数の実装
func ReadTextFile(filename string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. os.ReadFile()を使用してファイル全体を読み取り
	// 2. バイトスライスを文字列に変換
	// 3. エラーハンドリングを実装
	return "", nil
}

// WriteTextFile ヘルパー関数の実装
func WriteTextFile(filename, content string) error {
	// TODO: 実装する
	// ヒント:
	// 1. os.WriteFile()を使用してファイルに書き込み
	// 2. 適切なファイルパーミッション（0644）を設定
	// 3. エラーハンドリングを実装
	return nil
}

// AppendToFile ヘルパー関数の実装
func AppendToFile(filename, content string) error {
	// TODO: 実装する
	// ヒント:
	// 1. os.OpenFile()をO_APPEND|O_CREATE|O_WRITEONLYフラグで使用
	// 2. ファイルに内容を書き込み
	// 3. ファイルを適切にクローズ
	// 4. エラーハンドリングを実装
	return nil
}