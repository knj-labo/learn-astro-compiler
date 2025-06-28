package main

import (
	"fmt"
	"log"
	"time"
)

/*
Exercise 015: テンプレートエンジンとHTML生成

このエクササイズでは、Goのtext/templateとhtml/templateパッケージを使用して
動的なHTMLコンテンツの生成を学びます：

1. テンプレートの基本操作
   - テンプレートの作成と解析
   - データの埋め込み
   - 条件分岐とループ

2. HTML生成
   - 安全なHTML出力
   - XSS攻撃の防止
   - レスポンシブデザイン

3. 高度なテンプレート機能
   - テンプレートの継承
   - カスタム関数
   - パーシャルテンプレート

期待される動作:
- 動的なWebページの生成
- データドリブンなHTML出力
- 安全で保守可能なテンプレート
*/

func main() {
	fmt.Println("Exercise 015: Template Engine and HTML Generation")
	
	// テンプレートエンジンを初期化
	engine := NewTemplateEngine()
	
	// サンプルデータの準備
	fmt.Println("=== Preparing Sample Data ===")
	user := User{
		ID:       1,
		Name:     "Alice Johnson",
		Email:    "alice@example.com",
		Role:     "Admin",
		Active:   true,
		LastLogin: time.Now().Add(-2 * time.Hour),
	}
	
	posts := []BlogPost{
		{
			ID:          1,
			Title:       "Getting Started with Go",
			Content:     "Go is a powerful programming language...",
			Author:      "Alice Johnson",
			PublishedAt: time.Now().Add(-24 * time.Hour),
			Tags:        []string{"go", "programming", "tutorial"},
			Published:   true,
		},
		{
			ID:          2,
			Title:       "Advanced Go Concurrency",
			Content:     "Learn about channels and goroutines...",
			Author:      "Bob Smith",
			PublishedAt: time.Now().Add(-12 * time.Hour),
			Tags:        []string{"go", "concurrency", "advanced"},
			Published:   true,
		},
		{
			ID:          3,
			Title:       "Draft Post",
			Content:     "This is a draft post...",
			Author:      "Carol Brown",
			PublishedAt: time.Time{},
			Tags:        []string{"draft"},
			Published:   false,
		},
	}
	
	// ユーザープロファイルページの生成
	fmt.Println("\n=== Generating User Profile ===")
	profileHTML, err := engine.RenderUserProfile(user)
	if err != nil {
		log.Printf("Error rendering user profile: %v", err)
	} else {
		fmt.Printf("Generated profile HTML (%d bytes)\n", len(profileHTML))
		// HTMLをファイルに保存
		engine.SaveHTMLToFile("user_profile.html", profileHTML)
	}
	
	// ブログ一覧ページの生成
	fmt.Println("\n=== Generating Blog List ===")
	blogData := BlogData{
		Title: "My Tech Blog",
		Posts: posts,
		User:  user,
	}
	
	blogHTML, err := engine.RenderBlogList(blogData)
	if err != nil {
		log.Printf("Error rendering blog list: %v", err)
	} else {
		fmt.Printf("Generated blog HTML (%d bytes)\n", len(blogHTML))
		engine.SaveHTMLToFile("blog_list.html", blogHTML)
	}
	
	// 個別記事ページの生成
	fmt.Println("\n=== Generating Individual Post ===")
	if len(posts) > 0 {
		postHTML, err := engine.RenderBlogPost(posts[0])
		if err != nil {
			log.Printf("Error rendering blog post: %v", err)
		} else {
			fmt.Printf("Generated post HTML (%d bytes)\n", len(postHTML))
			engine.SaveHTMLToFile("blog_post.html", postHTML)
		}
	}
	
	// ダッシュボードページの生成
	fmt.Println("\n=== Generating Dashboard ===")
	dashboard := Dashboard{
		User:        user,
		Posts:       posts,
		TotalViews:  12543,
		TotalUsers:  456,
		TotalPosts:  len(posts),
		RecentViews: []PageView{
			{Page: "/blog/getting-started", Views: 234, Date: time.Now().Add(-1 * time.Hour)},
			{Page: "/blog/advanced-concurrency", Views: 189, Date: time.Now().Add(-2 * time.Hour)},
			{Page: "/about", Views: 87, Date: time.Now().Add(-3 * time.Hour)},
		},
	}
	
	dashboardHTML, err := engine.RenderDashboard(dashboard)
	if err != nil {
		log.Printf("Error rendering dashboard: %v", err)
	} else {
		fmt.Printf("Generated dashboard HTML (%d bytes)\n", len(dashboardHTML))
		engine.SaveHTMLToFile("dashboard.html", dashboardHTML)
	}
	
	// エラーページの生成
	fmt.Println("\n=== Generating Error Pages ===")
	errorPages := []struct {
		code    int
		message string
		file    string
	}{
		{404, "Page Not Found", "404.html"},
		{500, "Internal Server Error", "500.html"},
		{403, "Access Forbidden", "403.html"},
	}
	
	for _, errPage := range errorPages {
		errorHTML, err := engine.RenderErrorPage(errPage.code, errPage.message)
		if err != nil {
			log.Printf("Error rendering %d page: %v", errPage.code, err)
		} else {
			fmt.Printf("Generated %d error page\n", errPage.code)
			engine.SaveHTMLToFile(errPage.file, errorHTML)
		}
	}
	
	// メール テンプレートの生成
	fmt.Println("\n=== Generating Email Templates ===")
	emailData := EmailData{
		RecipientName: "Alice Johnson",
		Subject:       "Welcome to Our Platform",
		Message:       "Thank you for joining our platform. We're excited to have you!",
		ActionURL:     "https://example.com/activate",
		ActionText:    "Activate Account",
	}
	
	emailHTML, err := engine.RenderEmailTemplate(emailData)
	if err != nil {
		log.Printf("Error rendering email template: %v", err)
	} else {
		fmt.Printf("Generated email HTML (%d bytes)\n", len(emailHTML))
		engine.SaveHTMLToFile("email_template.html", emailHTML)
	}
	
	// JSONからテンプレートへの変換
	fmt.Println("\n=== JSON to Template ===")
	jsonData := `{
		"name": "Dynamic User",
		"items": ["Item 1", "Item 2", "Item 3"],
		"count": 42
	}`
	
	dynamicHTML, err := engine.RenderFromJSON("dynamic_template", jsonData)
	if err != nil {
		log.Printf("Error rendering from JSON: %v", err)
	} else {
		fmt.Printf("Generated dynamic HTML (%d bytes)\n", len(dynamicHTML))
		engine.SaveHTMLToFile("dynamic.html", dynamicHTML)
	}
	
	fmt.Println("\n=== Template Generation Complete ===")
	fmt.Println("Generated HTML files:")
	files := []string{
		"user_profile.html", "blog_list.html", "blog_post.html",
		"dashboard.html", "404.html", "500.html", "403.html",
		"email_template.html", "dynamic.html",
	}
	for _, file := range files {
		fmt.Printf("  - %s\n", file)
	}
}

// User構造体
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	LastLogin time.Time `json:"last_login"`
}

// BlogPost構造体
type BlogPost struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Tags        []string  `json:"tags"`
	Published   bool      `json:"published"`
}

// BlogData構造体
type BlogData struct {
	Title string     `json:"title"`
	Posts []BlogPost `json:"posts"`
	User  User       `json:"user"`
}

// Dashboard構造体
type Dashboard struct {
	User        User       `json:"user"`
	Posts       []BlogPost `json:"posts"`
	TotalViews  int        `json:"total_views"`
	TotalUsers  int        `json:"total_users"`
	TotalPosts  int        `json:"total_posts"`
	RecentViews []PageView `json:"recent_views"`
}

// PageView構造体
type PageView struct {
	Page  string    `json:"page"`
	Views int       `json:"views"`
	Date  time.Time `json:"date"`
}

// EmailData構造体
type EmailData struct {
	RecipientName string `json:"recipient_name"`
	Subject       string `json:"subject"`
	Message       string `json:"message"`
	ActionURL     string `json:"action_url"`
	ActionText    string `json:"action_text"`
}

// TemplateEngine構造体
type TemplateEngine struct {}

// NewTemplateEngine関数の実装
func NewTemplateEngine() *TemplateEngine {
	// TODO: 実装する
	// ヒント:
	// 1. TemplateEngine構造体を初期化
	// 2. カスタム関数を登録
	// 3. 基本テンプレートを準備
	return nil
}

// RenderUserProfile メソッドの実装
func (te *TemplateEngine) RenderUserProfile(user User) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. ユーザープロファイル用のHTMLテンプレートを作成
	// 2. ユーザーデータを埋め込み
	// 3. html/templateパッケージを使用
	// 4. XSS防止を考慮
	return "", nil
}

// RenderBlogList メソッドの実装
func (te *TemplateEngine) RenderBlogList(data BlogData) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. ブログ一覧用のHTMLテンプレートを作成
	// 2. 投稿リストをループで表示
	// 3. 公開済みの投稿のみ表示
	// 4. タグ表示機能を実装
	return "", nil
}

// RenderBlogPost メソッドの実装
func (te *TemplateEngine) RenderBlogPost(post BlogPost) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. 個別記事用のHTMLテンプレートを作成
	// 2. 記事内容の安全な表示
	// 3. 日時のフォーマット
	// 4. タグの表示
	return "", nil
}

// RenderDashboard メソッドの実装
func (te *TemplateEngine) RenderDashboard(dashboard Dashboard) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. ダッシュボード用のHTMLテンプレートを作成
	// 2. 統計情報の表示
	// 3. 最近の投稿一覧
	// 4. チャートやグラフの基本HTML
	return "", nil
}

// RenderErrorPage メソッドの実装
func (te *TemplateEngine) RenderErrorPage(code int, message string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. エラーページ用のHTMLテンプレートを作成
	// 2. エラーコードとメッセージを表示
	// 3. ホームページへのリンクを含む
	// 4. レスポンシブデザインを考慮
	return "", nil
}

// RenderEmailTemplate メソッドの実装
func (te *TemplateEngine) RenderEmailTemplate(data EmailData) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. メール用のHTMLテンプレートを作成
	// 2. インラインCSSを使用
	// 3. 呼び出しアクション（CTA）ボタンを含む
	// 4. メールクライアント対応を考慮
	return "", nil
}

// RenderFromJSON メソッドの実装
func (te *TemplateEngine) RenderFromJSON(templateName, jsonData string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. JSONデータをパース
	// 2. 動的テンプレートを作成
	// 3. データを埋め込んでレンダリング
	// 4. エラーハンドリングを実装
	return "", nil
}

// SaveHTMLToFile メソッドの実装
func (te *TemplateEngine) SaveHTMLToFile(filename, html string) error {
	// TODO: 実装する
	// ヒント:
	// 1. ファイルを作成または上書き
	// 2. HTMLコンテンツを書き込み
	// 3. 適切なファイルパーミッションを設定
	// 4. エラーハンドリングを実装
	return nil
}

// GetTemplateFunction メソッドの実装
func (te *TemplateEngine) GetTemplateFunction() map[string]interface{} {
	// TODO: 実装する
	// ヒント:
	// 1. カスタム関数のマップを作成
	// 2. 日時フォーマット関数
	// 3. 文字列操作関数
	// 4. 数値フォーマット関数
	// 5. 条件判定関数
	return nil
}

// formatDate ヘルパー関数の実装
func formatDate(t time.Time) string {
	// TODO: 実装する
	// ヒント:
	// 1. 日時を読みやすい形式にフォーマット
	// 2. "2006-01-02 15:04:05" 形式を使用
	return ""
}

// truncateText ヘルパー関数の実装
func truncateText(text string, length int) string {
	// TODO: 実装する
	// ヒント:
	// 1. 指定された長さでテキストを切り詰め
	// 2. 必要に応じて"..."を追加
	// 3. 単語の境界を考慮
	return ""
}

// joinStrings ヘルパー関数の実装
func joinStrings(strs []string, sep string) string {
	// TODO: 実装する
	// ヒント:
	// 1. strings.Join()を使用
	// 2. 空の配列をハンドリング
	return ""
}

// isEven ヘルパー関数の実装
func isEven(n int) bool {
	// TODO: 実装する
	// ヒント:
	// 1. 偶数判定を実装
	// 2. modulo演算子を使用
	return false
}