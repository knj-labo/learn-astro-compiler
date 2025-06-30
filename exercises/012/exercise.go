package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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
	// 1. CSVProcessor構造体を初期化して返す
	return &CSVProcessor{}
}

// ReadCustomersFromCSV メソッドの実装
func (cp *CSVProcessor) ReadCustomersFromCSV(filename string) ([]Customer, error) {
	// 1. ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 2. CSVリーダーを作成
	reader := csv.NewReader(file)
	
	// 3. 全レコードを読み取り
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return []Customer{}, nil
	}

	// 4. ヘッダー行をスキップして顧客データを解析
	var customers []Customer
	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) != 5 {
			continue
		}

		// 5. 各行をCustomer構造体に変換
		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		age, err := strconv.Atoi(record[3])
		if err != nil {
			continue
		}

		customer := Customer{
			ID:    id,
			Name:  record[1],
			Email: record[2],
			Age:   age,
			City:  record[4],
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// WriteCustomersToCSV メソッドの実装
func (cp *CSVProcessor) WriteCustomersToCSV(filename string, customers []Customer) error {
	// 1. ファイルを作成
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 2. CSVライターを作成
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 3. ヘッダー行を書き込み
	header := []string{"id", "name", "email", "age", "city"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// 4. 各Customer構造体をCSV行に変換して書き込み
	for _, customer := range customers {
		record := []string{
			strconv.Itoa(customer.ID),
			customer.Name,
			customer.Email,
			strconv.Itoa(customer.Age),
			customer.City,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// FilterByAge メソッドの実装
func (cp *CSVProcessor) FilterByAge(customers []Customer, minAge int) []Customer {
	// 1. 新しいスライスを作成
	var filtered []Customer
	
	// 2. 各顧客の年齢をチェック
	for _, customer := range customers {
		// 3. 条件を満たす顧客のみを新しいスライスに追加
		if customer.Age >= minAge {
			filtered = append(filtered, customer)
		}
	}
	
	return filtered
}

// GroupByCity メソッドの実装
func (cp *CSVProcessor) GroupByCity(customers []Customer) map[string][]Customer {
	// 1. マップを作成（キー: 都市名、値: Customer スライス）
	cityGroups := make(map[string][]Customer)
	
	// 2. 各顧客を対応する都市のスライスに追加
	for _, customer := range customers {
		cityGroups[customer.City] = append(cityGroups[customer.City], customer)
	}
	
	return cityGroups
}

// CalculateStats メソッドの実装
func (cp *CSVProcessor) CalculateStats(customers []Customer) CustomerStats {
	// 1. 顧客数をカウント
	totalCustomers := len(customers)
	
	if totalCustomers == 0 {
		return CustomerStats{}
	}
	
	// 2. 年齢の合計、最小値、最大値を計算
	totalAge := 0
	minAge := customers[0].Age
	maxAge := customers[0].Age
	
	// 4. ユニークな都市数をカウント
	cities := make(map[string]bool)
	
	for _, customer := range customers {
		totalAge += customer.Age
		
		if customer.Age < minAge {
			minAge = customer.Age
		}
		if customer.Age > maxAge {
			maxAge = customer.Age
		}
		
		cities[customer.City] = true
	}
	
	// 3. 平均年齢を計算
	averageAge := float64(totalAge) / float64(totalCustomers)
	
	// 5. CustomerStats構造体を作成して返す
	return CustomerStats{
		TotalCustomers: totalCustomers,
		AverageAge:     averageAge,
		MinAge:         minAge,
		MaxAge:         maxAge,
		CitiesCount:    len(cities),
	}
}

// GenerateReport メソッドの実装
func (cp *CSVProcessor) GenerateReport(filename string, customers []Customer, stats CustomerStats) error {
	// 1. ファイルを作成
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 2. レポートのヘッダーを書き込み
	_, err = file.WriteString("=== Customer Report ===\n\n")
	if err != nil {
		return err
	}

	// 3. 統計情報を書き込み
	_, err = file.WriteString(fmt.Sprintf("Statistics:\n"))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("  Total Customers: %d\n", stats.TotalCustomers))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("  Average Age: %.2f\n", stats.AverageAge))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("  Min Age: %d\n", stats.MinAge))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("  Max Age: %d\n", stats.MaxAge))
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("  Cities Count: %d\n\n", stats.CitiesCount))
	if err != nil {
		return err
	}

	// 4. 顧客一覧を書き込み
	_, err = file.WriteString("Customer List:\n")
	if err != nil {
		return err
	}
	
	for _, customer := range customers {
		_, err = file.WriteString(fmt.Sprintf("  ID: %d, Name: %s, Email: %s, Age: %d, City: %s\n",
			customer.ID, customer.Name, customer.Email, customer.Age, customer.City))
		if err != nil {
			return err
		}
	}

	return nil
}

// FileExists メソッドの実装
func (cp *CSVProcessor) FileExists(filename string) bool {
	// 1. os.Stat()を使用してファイル情報を取得
	_, err := os.Stat(filename)
	// 2. エラーをチェックしてファイルの存在を判定
	return err == nil
}

// GetFileSize メソッドの実装
func (cp *CSVProcessor) GetFileSize(filename string) (int64, error) {
	// 1. os.Stat()を使用してファイル情報を取得
	fileInfo, err := os.Stat(filename)
	if err != nil {
		// 3. エラーハンドリングを実装
		return 0, err
	}
	
	// 2. FileInfo.Size()でファイルサイズを取得
	return fileInfo.Size(), nil
}

// ReadTextFile ヘルパー関数の実装
func ReadTextFile(filename string) (string, error) {
	// 1. os.ReadFile()を使用してファイル全体を読み取り
	data, err := os.ReadFile(filename)
	if err != nil {
		// 3. エラーハンドリングを実装
		return "", err
	}
	
	// 2. バイトスライスを文字列に変換
	return string(data), nil
}

// WriteTextFile ヘルパー関数の実装
func WriteTextFile(filename, content string) error {
	// 1. os.WriteFile()を使用してファイルに書き込み
	// 2. 適切なファイルパーミッション（0644）を設定
	// 3. エラーハンドリングを実装
	return os.WriteFile(filename, []byte(content), 0644)
}

// AppendToFile ヘルパー関数の実装
func AppendToFile(filename, content string) error {
	// 1. os.OpenFile()をO_APPEND|O_CREATE|O_WRITE-ONLYフラグで使用
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// 4. エラーハンドリングを実装
		return err
	}
	// 3. ファイルを適切にクローズ
	defer file.Close()

	// 2. ファイルに内容を書き込み
	_, err = file.WriteString(content)
	return err
}