package main

import (
	"fmt"
	"log"
)

/*
Exercise 013: 正規表現とテキスト処理

このエクササイズでは、正規表現を使ったテキスト処理と文字列操作を学びます：

1. 正規表現の基本操作
   - パターンマッチング
   - 文字列の検索と置換
   - グループキャプチャ

2. テキスト解析
   - ログファイルの解析
   - データ抽出と検証
   - テキストのクリーニング

3. 文字列処理
   - テキストの分割と結合
   - フォーマット変換
   - 文字列の正規化

期待される動作:
- 様々なパターンのテキストマッチング
- ログデータからの情報抽出
- テキストの変換と整形
*/

func main() {
	fmt.Println("Exercise 013: Regular Expressions and Text Processing")
	
	// テキストプロセッサーを初期化
	processor := NewTextProcessor()
	
	// サンプルテキストデータ
	fmt.Println("=== Sample Text Processing ===")
	sampleText := `
	Contact us at: support@example.com or sales@company.org
	Phone numbers: +1-555-123-4567, (555) 987-6543
	Visit our website: https://www.example.com and http://blog.company.org
	Follow us on social media @company_official #golang
	`
	
	fmt.Printf("Original text:\n%s\n", sampleText)
	
	// メールアドレスの抽出
	fmt.Println("=== Email Address Extraction ===")
	emails := processor.ExtractEmails(sampleText)
	fmt.Printf("Found %d email addresses:\n", len(emails))
	for _, email := range emails {
		fmt.Printf("  - %s\n", email)
	}
	
	// 電話番号の抽出
	fmt.Println("\n=== Phone Number Extraction ===")
	phones := processor.ExtractPhoneNumbers(sampleText)
	fmt.Printf("Found %d phone numbers:\n", len(phones))
	for _, phone := range phones {
		fmt.Printf("  - %s\n", phone)
	}
	
	// URLの抽出
	fmt.Println("\n=== URL Extraction ===")
	urls := processor.ExtractURLs(sampleText)
	fmt.Printf("Found %d URLs:\n", len(urls))
	for _, url := range urls {
		fmt.Printf("  - %s\n", url)
	}
	
	// ハッシュタグとメンションの抽出
	fmt.Println("\n=== Social Media Tags ===")
	hashtags := processor.ExtractHashtags(sampleText)
	mentions := processor.ExtractMentions(sampleText)
	fmt.Printf("Hashtags: %v\n", hashtags)
	fmt.Printf("Mentions: %v\n", mentions)
	
	// ログファイル解析のデモ
	fmt.Println("\n=== Log File Analysis ===")
	logEntries := []string{
		"2024-01-15 10:30:45 [INFO] User login successful: user@example.com from 192.168.1.100",
		"2024-01-15 10:31:02 [ERROR] Database connection failed: timeout after 30s",
		"2024-01-15 10:31:15 [WARN] High memory usage detected: 85% used",
		"2024-01-15 10:32:00 [INFO] User logout: user@example.com session duration 15m30s",
		"2024-01-15 10:33:45 [ERROR] API request failed: 404 Not Found /api/users/999",
	}
	
	for _, entry := range logEntries {
		logData := processor.ParseLogEntry(entry)
		if logData != nil {
			fmt.Printf("Parsed: %+v\n", *logData)
		}
	}
	
	// テキストクリーニング
	fmt.Println("\n=== Text Cleaning ===")
	dirtyText := "  Hello,   World!!!   This is    a   test...   "
	cleaned := processor.CleanText(dirtyText)
	fmt.Printf("Original: '%s'\n", dirtyText)
	fmt.Printf("Cleaned:  '%s'\n", cleaned)
	
	// データ検証
	fmt.Println("\n=== Data Validation ===")
	testData := []string{
		"user@example.com",
		"invalid-email",
		"https://www.example.com",
		"not-a-url",
		"+1-555-123-4567",
		"invalid-phone",
		"Strong_Password123!",
		"weak",
	}
	
	for _, data := range testData {
		fmt.Printf("'%s':\n", data)
		fmt.Printf("  Email: %t\n", processor.IsValidEmail(data))
		fmt.Printf("  URL: %t\n", processor.IsValidURL(data))
		fmt.Printf("  Phone: %t\n", processor.IsValidPhoneNumber(data))
		fmt.Printf("  Strong Password: %t\n", processor.IsStrongPassword(data))
		fmt.Println()
	}
	
	// テキスト変換
	fmt.Println("=== Text Transformations ===")
	originalText := "Hello, World! This is a Test."
	fmt.Printf("Original: %s\n", originalText)
	fmt.Printf("Snake Case: %s\n", processor.ToSnakeCase(originalText))
	fmt.Printf("Camel Case: %s\n", processor.ToCamelCase(originalText))
	fmt.Printf("Kebab Case: %s\n", processor.ToKebabCase(originalText))
	
	// 統計情報
	fmt.Println("\n=== Text Statistics ===")
	stats := processor.AnalyzeText(sampleText)
	fmt.Printf("Text analysis: %+v\n", stats)
}

// LogEntry構造体
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	IP        string `json:"ip,omitempty"`
	Email     string `json:"email,omitempty"`
	Duration  string `json:"duration,omitempty"`
	StatusCode string `json:"status_code,omitempty"`
}

// TextStats構造体
type TextStats struct {
	CharCount     int     `json:"char_count"`
	WordCount     int     `json:"word_count"`
	LineCount     int     `json:"line_count"`
	SentenceCount int     `json:"sentence_count"`
	AvgWordLength float64 `json:"avg_word_length"`
	EmailCount    int     `json:"email_count"`
	URLCount      int     `json:"url_count"`
	PhoneCount    int     `json:"phone_count"`
}

// TextProcessor構造体
type TextProcessor struct{}

// NewTextProcessor関数の実装
func NewTextProcessor() *TextProcessor {
	// TODO: 実装する
	// ヒント:
	// 1. TextProcessor構造体を初期化して返す
	return nil
}

// ExtractEmails メソッドの実装
func (tp *TextProcessor) ExtractEmails(text string) []string {
	// TODO: 実装する
	// ヒント:
	// 1. regexp.MustCompile()でメールアドレスの正規表現を作成
	// 2. FindAllString()でマッチするすべての文字列を取得
	// 3. パターン: [a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}
	return nil
}

// ExtractPhoneNumbers メソッドの実装
func (tp *TextProcessor) ExtractPhoneNumbers(text string) []string {
	// TODO: 実装する
	// ヒント:
	// 1. 複数の電話番号フォーマットをサポート
	// 2. \+?\d{1,3}[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4} のようなパターン
	// 3. FindAllString()を使用
	return nil
}

// ExtractURLs メソッドの実装
func (tp *TextProcessor) ExtractURLs(text string) []string {
	// TODO: 実装する
	// ヒント:
	// 1. httpとhttpsの両方をサポート
	// 2. パターン: https?://[^\s]+
	// 3. FindAllString()を使用
	return nil
}

// ExtractHashtags メソッドの実装
func (tp *TextProcessor) ExtractHashtags(text string) []string {
	// TODO: 実装する
	// ヒント:
	// 1. #で始まる文字列を抽出
	// 2. パターン: #[a-zA-Z0-9_]+
	// 3. FindAllString()を使用
	return nil
}

// ExtractMentions メソッドの実装
func (tp *TextProcessor) ExtractMentions(text string) []string {
	// TODO: 実装する
	// ヒント:
	// 1. @で始まる文字列を抽出
	// 2. パターン: @[a-zA-Z0-9_]+
	// 3. FindAllString()を使用
	return nil
}

// ParseLogEntry メソッドの実装
func (tp *TextProcessor) ParseLogEntry(logLine string) *LogEntry {
	// TODO: 実装する
	// ヒント:
	// 1. 日時、ログレベル、メッセージを抽出
	// 2. 可能であればIP、メール、期間、ステータスコードも抽出
	// 3. 名前付きキャプチャグループを使用
	// 4. LogEntry構造体に格納して返す
	return nil
}

// CleanText メソッドの実装
func (tp *TextProcessor) CleanText(text string) string {
	// TODO: 実装する
	// ヒント:
	// 1. 先頭と末尾の空白を削除
	// 2. 連続する空白を単一の空白に置換
	// 3. 連続する句読点を削除
	// 4. strings.TrimSpaceとregexpを使用
	return ""
}

// IsValidEmail メソッドの実装
func (tp *TextProcessor) IsValidEmail(email string) bool {
	// TODO: 実装する
	// ヒント:
	// 1. メールアドレスの正規表現パターンを作成
	// 2. MatchString()を使用して検証
	return false
}

// IsValidURL メソッドの実装
func (tp *TextProcessor) IsValidURL(url string) bool {
	// TODO: 実装する
	// ヒント:
	// 1. URLの正規表現パターンを作成
	// 2. httpまたはhttpsで始まることを確認
	// 3. MatchString()を使用
	return false
}

// IsValidPhoneNumber メソッドの実装
func (tp *TextProcessor) IsValidPhoneNumber(phone string) bool {
	// TODO: 実装する
	// ヒント:
	// 1. 電話番号の正規表現パターンを作成
	// 2. 複数のフォーマットをサポート
	// 3. MatchString()を使用
	return false
}

// IsStrongPassword メソッドの実装
func (tp *TextProcessor) IsStrongPassword(password string) bool {
	// TODO: 実装する
	// ヒント:
	// 1. 最低8文字の長さをチェック
	// 2. 大文字、小文字、数字、特殊文字を含むかチェック
	// 3. 複数の正規表現を使用
	return false
}

// ToSnakeCase メソッドの実装
func (tp *TextProcessor) ToSnakeCase(text string) string {
	// TODO: 実装する
	// ヒント:
	// 1. 大文字を小文字に変換
	// 2. 空白と句読点をアンダースコアに置換
	// 3. 連続するアンダースコアを単一に正規化
	return ""
}

// ToCamelCase メソッドの実装
func (tp *TextProcessor) ToCamelCase(text string) string {
	// TODO: 実装する
	// ヒント:
	// 1. 単語を分割
	// 2. 最初の単語は小文字、以降は最初の文字を大文字に
	// 3. 句読点と空白を削除
	return ""
}

// ToKebabCase メソッドの実装
func (tp *TextProcessor) ToKebabCase(text string) string {
	// TODO: 実装する
	// ヒント:
	// 1. 大文字を小文字に変換
	// 2. 空白と句読点をハイフンに置換
	// 3. 連続するハイフンを単一に正規化
	return ""
}

// AnalyzeText メソッドの実装
func (tp *TextProcessor) AnalyzeText(text string) TextStats {
	// TODO: 実装する
	// ヒント:
	// 1. 文字数、単語数、行数、文数をカウント
	// 2. 平均単語長を計算
	// 3. メール、URL、電話番号の数をカウント
	// 4. TextStats構造体を作成して返す
	return TextStats{}
}

// RemovePattern ヘルパー関数の実装
func RemovePattern(text, pattern string) string {
	// TODO: 実装する
	// ヒント:
	// 1. regexp.MustCompile()でパターンをコンパイル
	// 2. ReplaceAllString()で空文字に置換
	return ""
}

// ReplacePattern ヘルパー関数の実装
func ReplacePattern(text, pattern, replacement string) string {
	// TODO: 実装する
	// ヒント:
	// 1. regexp.MustCompile()でパターンをコンパイル
	// 2. ReplaceAllString()で置換
	return ""
}

// CountMatches ヘルパー関数の実装
func CountMatches(text, pattern string) int {
	// TODO: 実装する
	// ヒント:
	// 1. regexp.MustCompile()でパターンをコンパイル
	// 2. FindAllString()でマッチを取得
	// 3. マッチ数を返す
	return 0
}