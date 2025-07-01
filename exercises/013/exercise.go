package main

import (
	"fmt"
	"regexp"
	"strings"
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
	// 1. TextProcessor構造体を初期化して返す
	return &TextProcessor{}
}

// ExtractEmails メソッドの実装
func (tp *TextProcessor) ExtractEmails(text string) []string {
	// 1. regexp.MustCompile()でメールアドレスの正規表現を作成
	// 3. パターン: [a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	
	// 2. FindAllString()でマッチするすべての文字列を取得
	return emailPattern.FindAllString(text, -1)
}

// ExtractPhoneNumbers メソッドの実装
func (tp *TextProcessor) ExtractPhoneNumbers(text string) []string {
	// 1. 複数の電話番号フォーマットをサポート
	// 2. \+?\d{1,3}[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4} のようなパターン
	phonePattern := regexp.MustCompile(`\+?\d{1,3}[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)
	
	// 3. FindAllString()を使用
	return phonePattern.FindAllString(text, -1)
}

// ExtractURLs メソッドの実装
func (tp *TextProcessor) ExtractURLs(text string) []string {
	// 1. httpとhttpsの両方をサポート
	// 2. パターン: https?://[^\s]+
	urlPattern := regexp.MustCompile(`https?://[^\s]+`)
	
	// 3. FindAllString()を使用
	return urlPattern.FindAllString(text, -1)
}

// ExtractHashtags メソッドの実装
func (tp *TextProcessor) ExtractHashtags(text string) []string {
	// 1. #で始まる文字列を抽出
	// 2. パターン: #[a-zA-Z0-9_]+
	hashtagPattern := regexp.MustCompile(`#[a-zA-Z0-9_]+`)
	
	// 3. FindAllString()を使用
	return hashtagPattern.FindAllString(text, -1)
}

// ExtractMentions メソッドの実装
func (tp *TextProcessor) ExtractMentions(text string) []string {
	// 1. @で始まる文字列を抽出
	// 2. パターン: @[a-zA-Z0-9_]+
	mentionPattern := regexp.MustCompile(`@[a-zA-Z0-9_]+`)
	
	// 3. FindAllString()を使用
	return mentionPattern.FindAllString(text, -1)
}

// ParseLogEntry メソッドの実装
func (tp *TextProcessor) ParseLogEntry(logLine string) *LogEntry {
	// 1. 日時、ログレベル、メッセージを抽出
	// 3. 名前付きキャプチャグループを使用
	logPattern := regexp.MustCompile(`(?P<timestamp>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(?P<level>\w+)\] (?P<message>.+)`)
	
	matches := logPattern.FindStringSubmatch(logLine)
	if len(matches) == 0 {
		return nil
	}
	
	// 4. LogEntry構造体に格納して返す
	entry := &LogEntry{
		Timestamp: matches[1],
		Level:     matches[2],
		Message:   matches[3],
	}
	
	// 2. 可能であればIP、メール、期間、ステータスコードも抽出
	// IPアドレスを抽出
	ipPattern := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	if ipMatch := ipPattern.FindString(logLine); ipMatch != "" {
		entry.IP = ipMatch
	}
	
	// メールアドレスを抽出
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	if emailMatch := emailPattern.FindString(logLine); emailMatch != "" {
		entry.Email = emailMatch
	}
	
	// 期間を抽出
	durationPattern := regexp.MustCompile(`\d+[ms]?\d*[sm]`)
	if durationMatch := durationPattern.FindString(logLine); durationMatch != "" {
		entry.Duration = durationMatch
	}
	
	// ステータスコードを抽出
	statusPattern := regexp.MustCompile(`\b[1-5]\d{2}\b`)
	if statusMatch := statusPattern.FindString(logLine); statusMatch != "" {
		entry.StatusCode = statusMatch
	}
	
	return entry
}

// CleanText メソッドの実装
func (tp *TextProcessor) CleanText(text string) string {
	// 1. 先頭と末尾の空白を削除
	// 4. strings.TrimSpaceとregexpを使用
	cleaned := strings.TrimSpace(text)
	
	// 2. 連続する空白を単一の空白に置換
	spacePattern := regexp.MustCompile(`\s+`)
	cleaned = spacePattern.ReplaceAllString(cleaned, " ")
	
	// 3. 連続する句読点を削除
	punctPattern := regexp.MustCompile(`[!.]{2,}`)
	cleaned = punctPattern.ReplaceAllString(cleaned, ".")
	
	return cleaned
}

// IsValidEmail メソッドの実装
func (tp *TextProcessor) IsValidEmail(email string) bool {
	// 1. メールアドレスの正規表現パターンを作成
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	
	// 2. MatchString()を使用して検証
	return emailPattern.MatchString(email)
}

// IsValidURL メソッドの実装
func (tp *TextProcessor) IsValidURL(url string) bool {
	// 1. URLの正規表現パターンを作成
	// 2. httpまたはhttpsで始まることを確認
	urlPattern := regexp.MustCompile(`^https?://[^\s]+$`)
	
	// 3. MatchString()を使用
	return urlPattern.MatchString(url)
}

// IsValidPhoneNumber メソッドの実装
func (tp *TextProcessor) IsValidPhoneNumber(phone string) bool {
	// 1. 電話番号の正規表現パターンを作成
	// 2. 複数のフォーマットをサポート
	patterns := []string{
		`^\+?\d{1,3}[-.\s]?\d{3}[-.\s]?\d{3}[-.\s]?\d{4}$`,           // +1-555-123-4567
		`^\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`,                     // (555) 987-6543
		`^\d{3}\.?\d{3}\.?\d{4}$`,                                   // 555.123.4567
	}
	
	// 3. MatchString()を使用
	for _, pattern := range patterns {
		if regexp.MustCompile(pattern).MatchString(phone) {
			return true
		}
	}
	
	return false
}

// IsStrongPassword メソッドの実装
func (tp *TextProcessor) IsStrongPassword(password string) bool {
	// 1. 最低8文字の長さをチェック
	if len(password) < 8 {
		return false
	}
	
	// 2. 大文字、小文字、数字、特殊文字を含むかチェック
	// 3. 複数の正規表現を使用
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
	
	return hasLower && hasUpper && hasDigit && hasSpecial
}

// ToSnakeCase メソッドの実装
func (tp *TextProcessor) ToSnakeCase(text string) string {
	// 1. 大文字を小文字に変換
	result := strings.ToLower(text)
	
	// 2. 空白と句読点をアンダースコアに置換
	nonAlphaPattern := regexp.MustCompile(`[^a-z0-9]+`)
	result = nonAlphaPattern.ReplaceAllString(result, "_")
	
	// 3. 連続するアンダースコアを単一に正規化
	underscorePattern := regexp.MustCompile(`_+`)
	result = underscorePattern.ReplaceAllString(result, "_")
	
	// 先頭と末尾のアンダースコアを削除
	result = strings.Trim(result, "_")
	
	return result
}

// ToCamelCase メソッドの実装
func (tp *TextProcessor) ToCamelCase(text string) string {
	// 3. 句読点と空白を削除
	cleanPattern := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	cleaned := cleanPattern.ReplaceAllString(text, " ")
	
	// 1. 単語を分割
	words := strings.Fields(cleaned)
	
	if len(words) == 0 {
		return ""
	}
	
	var result strings.Builder
	
	// 2. 最初の単語は小文字、以降は最初の文字を大文字に
	for i, word := range words {
		if i == 0 {
			result.WriteString(strings.ToLower(word))
		} else {
			if len(word) > 0 {
				result.WriteString(strings.ToUpper(string(word[0])) + strings.ToLower(word[1:]))
			}
		}
	}
	
	return result.String()
}

// ToKebabCase メソッドの実装
func (tp *TextProcessor) ToKebabCase(text string) string {
	// 1. 大文字を小文字に変換
	result := strings.ToLower(text)
	
	// 2. 空白と句読点をハイフンに置換
	nonAlphaPattern := regexp.MustCompile(`[^a-z0-9]+`)
	result = nonAlphaPattern.ReplaceAllString(result, "-")
	
	// 3. 連続するハイフンを単一に正規化
	hyphenPattern := regexp.MustCompile(`-+`)
	result = hyphenPattern.ReplaceAllString(result, "-")
	
	// 先頭と末尾のハイフンを削除
	result = strings.Trim(result, "-")
	
	return result
}

// AnalyzeText メソッドの実装
func (tp *TextProcessor) AnalyzeText(text string) TextStats {
	// 1. 文字数、単語数、行数、文数をカウント
	charCount := len(text)
	
	lines := strings.Split(text, "\n")
	lineCount := len(lines)
	
	words := strings.Fields(text)
	wordCount := len(words)
	
	sentences := regexp.MustCompile(`[.!?]+`).Split(text, -1)
	sentenceCount := len(sentences) - 1
	if sentenceCount < 0 {
		sentenceCount = 0
	}
	
	// 2. 平均単語長を計算
	var totalWordLength int
	for _, word := range words {
		// 句読点を除いた文字数をカウント
		cleanWord := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(word, "")
		totalWordLength += len(cleanWord)
	}
	
	var avgWordLength float64
	if wordCount > 0 {
		avgWordLength = float64(totalWordLength) / float64(wordCount)
	}
	
	// 3. メール、URL、電話番号の数をカウント
	emailCount := len(tp.ExtractEmails(text))
	urlCount := len(tp.ExtractURLs(text))
	phoneCount := len(tp.ExtractPhoneNumbers(text))
	
	// 4. TextStats構造体を作成して返す
	return TextStats{
		CharCount:     charCount,
		WordCount:     wordCount,
		LineCount:     lineCount,
		SentenceCount: sentenceCount,
		AvgWordLength: avgWordLength,
		EmailCount:    emailCount,
		URLCount:      urlCount,
		PhoneCount:    phoneCount,
	}
}

// RemovePattern ヘルパー関数の実装
func RemovePattern(text, pattern string) string {
	// 1. regexp.MustCompile()でパターンをコンパイル
	re := regexp.MustCompile(pattern)
	
	// 2. ReplaceAllString()で空文字に置換
	return re.ReplaceAllString(text, "")
}

// ReplacePattern ヘルパー関数の実装
func ReplacePattern(text, pattern, replacement string) string {
	// 1. regexp.MustCompile()でパターンをコンパイル
	re := regexp.MustCompile(pattern)
	
	// 2. ReplaceAllString()で置換
	return re.ReplaceAllString(text, replacement)
}

// CountMatches ヘルパー関数の実装
func CountMatches(text, pattern string) int {
	// 1. regexp.MustCompile()でパターンをコンパイル
	re := regexp.MustCompile(pattern)
	
	// 2. FindAllString()でマッチを取得
	matches := re.FindAllString(text, -1)
	
	// 3. マッチ数を返す
	return len(matches)
}