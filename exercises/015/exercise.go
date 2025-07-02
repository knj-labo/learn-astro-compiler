package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
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
type TemplateEngine struct {
	funcMap template.FuncMap
}

// NewTemplateEngine関数の実装
func NewTemplateEngine() *TemplateEngine {
	// 1. TemplateEngine構造体を初期化
	te := &TemplateEngine{}
	
	// 2. カスタム関数を登録
	te.funcMap = te.GetTemplateFunction()
	
	// 3. 基本テンプレートを準備
	return te
}

// RenderUserProfile メソッドの実装
func (te *TemplateEngine) RenderUserProfile(user User) (string, error) {
	// 1. ユーザープロファイル用のHTMLテンプレートを作成
	templateStr := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Name | title}}のプロファイル</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; background-color: #f5f5f5; }
        .profile-card { background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); max-width: 600px; margin: 0 auto; }
        .profile-header { text-align: center; border-bottom: 2px solid #eee; padding-bottom: 20px; margin-bottom: 20px; }
        .profile-info { display: grid; grid-template-columns: 1fr 1fr; gap: 15px; }
        .info-item { padding: 10px; background: #f8f9fa; border-radius: 5px; }
        .status { padding: 5px 10px; border-radius: 15px; font-size: 12px; font-weight: bold; }
        .active { background: #d4edda; color: #155724; }
        .inactive { background: #f8d7da; color: #721c24; }
        .role-badge { display: inline-block; padding: 3px 8px; background: #007bff; color: white; border-radius: 12px; font-size: 11px; }
    </style>
</head>
<body>
    <div class="profile-card">
        <div class="profile-header">
            <h1>{{.Name | title}}</h1>
            <p>ユーザーID: {{.ID}}</p>
            <span class="role-badge">{{.Role}}</span>
        </div>
        <div class="profile-info">
            <div class="info-item">
                <strong>メールアドレス:</strong><br>
                {{.Email}}
            </div>
            <div class="info-item">
                <strong>ステータス:</strong><br>
                {{if .Active}}
                    <span class="status active">アクティブ</span>
                {{else}}
                    <span class="status inactive">非アクティブ</span>
                {{end}}
            </div>
            <div class="info-item">
                <strong>最終ログイン:</strong><br>
                {{formatDate .LastLogin}}
            </div>
            <div class="info-item">
                <strong>権限レベル:</strong><br>
                {{.Role}}
            </div>
        </div>
    </div>
</body>
</html>`

	// 2. html/templateパッケージを使用してテンプレートを作成
	// 3. XSS防止を考慮
	tmpl, err := template.New("userProfile").Funcs(te.funcMap).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("テンプレートの解析に失敗: %v", err)
	}

	// 4. ユーザーデータを埋め込み
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, user); err != nil {
		return "", fmt.Errorf("テンプレートの実行に失敗: %v", err)
	}

	return buf.String(), nil
}

// RenderBlogList メソッドの実装
func (te *TemplateEngine) RenderBlogList(data BlogData) (string, error) {
	// 1. ブログ一覧用のHTMLテンプレートを作成
	templateStr := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; background-color: #f8f9fa; }
        .container { max-width: 1200px; margin: 0 auto; }
        .header { text-align: center; margin-bottom: 40px; padding: 30px; background: white; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .blog-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(350px, 1fr)); gap: 20px; }
        .blog-card { background: white; border-radius: 10px; padding: 25px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); transition: transform 0.2s; }
        .blog-card:hover { transform: translateY(-2px); }
        .blog-title { color: #333; margin-bottom: 10px; font-size: 1.4em; }
        .blog-meta { color: #666; font-size: 0.9em; margin-bottom: 15px; }
        .blog-content { color: #555; line-height: 1.6; margin-bottom: 15px; }
        .tags { margin-top: 15px; }
        .tag { display: inline-block; background: #e9ecef; color: #495057; padding: 3px 8px; border-radius: 12px; font-size: 0.8em; margin-right: 5px; margin-bottom: 5px; }
        .published { border-left: 4px solid #28a745; }
        .unpublished { border-left: 4px solid #6c757d; opacity: 0.7; }
        .user-info { text-align: center; margin-top: 20px; padding: 15px; background: #f8f9fa; border-radius: 8px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>{{.Title}}</h1>
            <div class="user-info">
                ようこそ、{{.User.Name}}さん ({{.User.Role}})
            </div>
        </div>
        
        <div class="blog-grid">
            {{range .Posts}}
                {{if .Published}}
                <div class="blog-card published">
                    <h2 class="blog-title">{{.Title}}</h2>
                    <div class="blog-meta">
                        作成者: {{.Author}} | 公開日: {{formatDate .PublishedAt}}
                    </div>
                    <div class="blog-content">
                        {{truncate 150 .Content}}
                    </div>
                    {{if .Tags}}
                    <div class="tags">
                        {{range .Tags}}
                            <span class="tag">{{.}}</span>
                        {{end}}
                    </div>
                    {{end}}
                </div>
                {{end}}
            {{end}}
        </div>
        
        {{$unpublishedCount := 0}}
        {{range .Posts}}
            {{if not .Published}}
                {{$unpublishedCount = add $unpublishedCount 1}}
            {{end}}
        {{end}}
        
        {{if gt $unpublishedCount 0}}
        <div style="margin-top: 40px; padding: 20px; background: #fff3cd; border-radius: 10px; text-align: center;">
            <strong>下書き: {{$unpublishedCount}}件の投稿が非公開です</strong>
        </div>
        {{end}}
    </div>
</body>
</html>`

	// 2. html/templateパッケージを使用してテンプレートを作成
	// 3. 公開済みの投稿のみ表示（テンプレート内で条件分岐）
	// 4. タグ表示機能を実装
	funcMap := te.funcMap
	funcMap["add"] = func(a, b int) int { return a + b }
	
	tmpl, err := template.New("blogList").Funcs(funcMap).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("テンプレートの解析に失敗: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("テンプレートの実行に失敗: %v", err)
	}

	return buf.String(), nil
}

// RenderBlogPost メソッドの実装
func (te *TemplateEngine) RenderBlogPost(post BlogPost) (string, error) {
	// 1. 個別記事用のHTMLテンプレートを作成
	templateStr := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; background-color: #f8f9fa; line-height: 1.6; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 40px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .post-header { border-bottom: 2px solid #eee; padding-bottom: 20px; margin-bottom: 30px; }
        .post-title { color: #333; margin-bottom: 10px; font-size: 2.5em; }
        .post-meta { color: #666; font-size: 0.9em; margin-bottom: 20px; }
        .post-content { color: #333; font-size: 1.1em; line-height: 1.8; margin-bottom: 30px; }
        .tags { margin-top: 30px; padding-top: 20px; border-top: 1px solid #eee; }
        .tag { display: inline-block; background: #007bff; color: white; padding: 5px 12px; border-radius: 15px; font-size: 0.8em; margin-right: 8px; margin-bottom: 5px; }
        .status-badge { padding: 5px 10px; border-radius: 15px; font-size: 0.8em; font-weight: bold; }
        .published { background: #d4edda; color: #155724; }
        .draft { background: #f8d7da; color: #721c24; }
    </style>
</head>
<body>
    <div class="container">
        <div class="post-header">
            <h1 class="post-title">{{.Title}}</h1>
            <div class="post-meta">
                作成者: <strong>{{.Author}}</strong> | 
                {{if .Published}}
                    公開日: {{formatDate .PublishedAt}}
                    <span class="status-badge published">公開済み</span>
                {{else}}
                    <span class="status-badge draft">下書き</span>
                {{end}}
            </div>
        </div>
        
        <div class="post-content">
            {{.Content}}
        </div>
        
        {{if .Tags}}
        <div class="tags">
            <strong>タグ:</strong><br>
            {{range .Tags}}
                <span class="tag">{{.}}</span>
            {{end}}
        </div>
        {{end}}
    </div>
</body>
</html>`

	// 2. 記事内容の安全な表示
	// 3. 日時のフォーマット
	// 4. タグの表示
	tmpl, err := template.New("blogPost").Funcs(te.funcMap).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("テンプレートの解析に失敗: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, post); err != nil {
		return "", fmt.Errorf("テンプレートの実行に失敗: %v", err)
	}

	return buf.String(), nil
}

// RenderDashboard メソッドの実装
func (te *TemplateEngine) RenderDashboard(dashboard Dashboard) (string, error) {
	// 1. ダッシュボード用のHTMLテンプレートを作成
	templateStr := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ダッシュボード - {{.User.Name}}</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; background-color: #f8f9fa; }
        .container { max-width: 1200px; margin: 0 auto; }
        .header { background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); margin-bottom: 20px; text-align: center; }
        .stats-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 20px; margin-bottom: 30px; }
        .stat-card { background: white; padding: 20px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); text-align: center; }
        .stat-number { font-size: 2.5em; font-weight: bold; color: #007bff; margin-bottom: 5px; }
        .stat-label { color: #666; font-size: 0.9em; }
        .recent-section { background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); margin-bottom: 20px; }
        .recent-item { padding: 15px; border-bottom: 1px solid #eee; display: flex; justify-content: space-between; align-items: center; }
        .recent-item:last-child { border-bottom: none; }
        .recent-views { background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .view-item { display: flex; justify-content: space-between; align-items: center; padding: 10px 0; border-bottom: 1px solid #f0f0f0; }
        .view-item:last-child { border-bottom: none; }
        .page-name { font-weight: bold; color: #333; }
        .view-count { background: #e9ecef; padding: 3px 8px; border-radius: 12px; font-size: 0.8em; }
        .view-date { color: #666; font-size: 0.8em; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>ダッシュボード</h1>
            <p>ようこそ、{{.User.Name}}さん ({{.User.Role}})</p>
        </div>
        
        <div class="stats-grid">
            <div class="stat-card">
                <div class="stat-number">{{.TotalViews}}</div>
                <div class="stat-label">総ビュー数</div>
            </div>
            <div class="stat-card">
                <div class="stat-number">{{.TotalUsers}}</div>
                <div class="stat-label">ユーザー数</div>
            </div>
            <div class="stat-card">
                <div class="stat-number">{{.TotalPosts}}</div>
                <div class="stat-label">投稿数</div>
            </div>
            <div class="stat-card">
                <div class="stat-number">{{len .Posts}}</div>
                <div class="stat-label">現在の記事</div>
            </div>
        </div>
        
        <div class="recent-section">
            <h2>最近の投稿</h2>
            {{if .Posts}}
                {{range .Posts}}
                    {{if .Published}}
                    <div class="recent-item">
                        <div>
                            <strong>{{.Title}}</strong><br>
                            <small>作成者: {{.Author}} | {{formatDate .PublishedAt}}</small>
                        </div>
                        <div>
                            {{if .Tags}}
                                {{range .Tags}}
                                    <span style="background: #e9ecef; padding: 2px 6px; border-radius: 8px; font-size: 0.7em; margin-left: 3px;">{{.}}</span>
                                {{end}}
                            {{end}}
                        </div>
                    </div>
                    {{end}}
                {{end}}
            {{else}}
                <p>投稿がありません。</p>
            {{end}}
        </div>
        
        <div class="recent-views">
            <h2>最近のページビュー</h2>
            {{if .RecentViews}}
                {{range .RecentViews}}
                <div class="view-item">
                    <div class="page-name">{{.Page}}</div>
                    <div>
                        <span class="view-count">{{.Views}} views</span>
                        <span class="view-date">{{formatDate .Date}}</span>
                    </div>
                </div>
                {{end}}
            {{else}}
                <p>ビューデータがありません。</p>
            {{end}}
        </div>
    </div>
</body>
</html>`

	// 2. 統計情報の表示
	// 3. 最近の投稿一覧
	// 4. チャートやグラフの基本HTML
	tmpl, err := template.New("dashboard").Funcs(te.funcMap).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("テンプレートの解析に失敗: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, dashboard); err != nil {
		return "", fmt.Errorf("テンプレートの実行に失敗: %v", err)
	}

	return buf.String(), nil
}

// RenderErrorPage メソッドの実装
func (te *TemplateEngine) RenderErrorPage(code int, message string) (string, error) {
	// 1. エラーページ用のHTMLテンプレートを作成
	templateStr := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>エラー {{.Code}} - {{.Message}}</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 0; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); min-height: 100vh; display: flex; align-items: center; justify-content: center; }
        .error-container { background: white; padding: 60px 40px; border-radius: 20px; box-shadow: 0 20px 60px rgba(0,0,0,0.2); text-align: center; max-width: 500px; width: 90%; }
        .error-code { font-size: 6em; color: #667eea; font-weight: bold; margin-bottom: 20px; text-shadow: 2px 2px 4px rgba(0,0,0,0.1); }
        .error-message { font-size: 1.5em; color: #333; margin-bottom: 30px; }
        .error-description { color: #666; margin-bottom: 40px; line-height: 1.6; }
        .home-button { display: inline-block; background: linear-gradient(45deg, #667eea, #764ba2); color: white; padding: 15px 30px; text-decoration: none; border-radius: 25px; font-weight: bold; transition: transform 0.2s; }
        .home-button:hover { transform: translateY(-2px); }
        .error-icon { font-size: 3em; margin-bottom: 20px; }
        @media (max-width: 600px) {
            .error-container { padding: 40px 20px; }
            .error-code { font-size: 4em; }
            .error-message { font-size: 1.2em; }
        }
    </style>
</head>
<body>
    <div class="error-container">
        <div class="error-icon">
            {{if eq .Code 404}}📜{{else if eq .Code 500}}⚠️{{else if eq .Code 403}}🚫{{else}}❌{{end}}
        </div>
        <div class="error-code">{{.Code}}</div>
        <div class="error-message">{{.Message}}</div>
        <div class="error-description">
            {{if eq .Code 404}}
                お探しのページが見つかりませんでした。URLをご確認ください。
            {{else if eq .Code 500}}
                サーバー内部でエラーが発生しました。しばらくしてから再度お試しください。
            {{else if eq .Code 403}}
                このページへのアクセス権限がありません。
            {{else}}
                予期しないエラーが発生しました。
            {{end}}
        </div>
        <a href="/" class="home-button">ホームに戻る</a>
    </div>
</body>
</html>`

	// 2. エラーコードとメッセージを表示
	// 3. ホームページへのリンクを含む
	// 4. レスポンシブデザインを考慮
	data := struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: message,
	}

	tmpl, err := template.New("errorPage").Funcs(te.funcMap).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("テンプレートの解析に失敗: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("テンプレートの実行に失敗: %v", err)
	}

	return buf.String(), nil
}

// RenderEmailTemplate メソッドの実装
func (te *TemplateEngine) RenderEmailTemplate(data EmailData) (string, error) {
	// 1. メール用のHTMLテンプレートを作成
	// 2. インラインCSSを使用
	templateStr := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Subject}}</title>
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%" style="background-color: #f4f4f4;">
        <tr>
            <td align="center" style="padding: 20px 0;">
                <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="600" style="background-color: #ffffff; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1);">
                    <!-- Header -->
                    <tr>
                        <td style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 40px 30px; text-align: center; border-radius: 8px 8px 0 0;">
                            <h1 style="color: #ffffff; margin: 0; font-size: 28px; font-weight: bold;">あなたのアカウント</h1>
                        </td>
                    </tr>
                    
                    <!-- Content -->
                    <tr>
                        <td style="padding: 40px 30px;">
                            <h2 style="color: #333333; margin: 0 0 20px 0; font-size: 24px;">こんにちは、{{.RecipientName}}さん！</h2>
                            
                            <p style="color: #666666; font-size: 16px; line-height: 1.6; margin: 0 0 20px 0;">
                                {{.Message}}
                            </p>
                            
                            <p style="color: #666666; font-size: 16px; line-height: 1.6; margin: 0 0 30px 0;">
                                下記のボタンをクリックして、アカウントの設定を完了してください。
                            </p>
                            
                            <!-- CTA Button -->
                            <table role="presentation" cellspacing="0" cellpadding="0" border="0" style="margin: 30px 0;">
                                <tr>
                                    <td style="text-align: center;">
                                        <a href="{{.ActionURL}}" style="background: linear-gradient(45deg, #667eea, #764ba2); color: #ffffff; padding: 15px 30px; text-decoration: none; border-radius: 25px; font-weight: bold; font-size: 16px; display: inline-block; box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);">
                                            {{.ActionText}}
                                        </a>
                                    </td>
                                </tr>
                            </table>
                            
                            <p style="color: #999999; font-size: 14px; line-height: 1.5; margin: 30px 0 0 0;">
                                ボタンが機能しない場合は、以下のURLをコピーしてブラウザに貼り付けてください：<br>
                                <a href="{{.ActionURL}}" style="color: #667eea; word-break: break-all;">{{.ActionURL}}</a>
                            </p>
                        </td>
                    </tr>
                    
                    <!-- Footer -->
                    <tr>
                        <td style="background-color: #f8f9fa; padding: 20px 30px; text-align: center; border-radius: 0 0 8px 8px; border-top: 1px solid #e9ecef;">
                            <p style="color: #999999; font-size: 12px; margin: 0; line-height: 1.4;">
                                このメールに心当たりがない場合は、無視してください。<br>
                                &copy; 2023 あなたの会社. All rights reserved.
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`

	// 3. 呼び出しアクション（CTA）ボタンを含む
	// 4. メールクライアント対応を考慮
	tmpl, err := template.New("emailTemplate").Funcs(te.funcMap).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("テンプレートの解析に失敗: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("テンプレートの実行に失敗: %v", err)
	}

	return buf.String(), nil
}

// RenderFromJSON メソッドの実装
func (te *TemplateEngine) RenderFromJSON(templateName, jsonData string) (string, error) {
	// 1. JSONデータをパース
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return "", fmt.Errorf("JSONのパースに失敗: %v", err)
	}

	// 2. 動的テンプレートを作成
	templateStr := `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>動的テンプレート - {{.name}}</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; background-color: #f5f5f5; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .header { text-align: center; margin-bottom: 30px; padding-bottom: 20px; border-bottom: 2px solid #eee; }
        .data-section { margin-bottom: 30px; }
        .data-item { background: #f8f9fa; padding: 15px; border-radius: 8px; margin-bottom: 10px; }
        .data-label { font-weight: bold; color: #333; }
        .data-value { color: #666; margin-top: 5px; }
        .list-container { background: #e9ecef; padding: 20px; border-radius: 8px; }
        .list-item { background: white; padding: 10px; margin: 5px 0; border-radius: 5px; border-left: 4px solid #007bff; }
        .count-badge { display: inline-block; background: #007bff; color: white; padding: 5px 10px; border-radius: 15px; font-size: 0.8em; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>動的テンプレート: {{.name}}</h1>
            <p>テンプレート名: {{.templateName}}</p>
        </div>
        
        <div class="data-section">
            {{range $key, $value := .}}
                {{if ne $key "templateName"}}
                    {{if eq (printf "%T" $value) "[]interface {}"}}
                        <div class="data-item">
                            <div class="data-label">{{$key}} <span class="count-badge">{{len $value}} items</span></div>
                            <div class="list-container">
                                {{range $value}}
                                    <div class="list-item">{{.}}</div>
                                {{end}}
                            </div>
                        </div>
                    {{else}}
                        <div class="data-item">
                            <div class="data-label">{{$key}}</div>
                            <div class="data-value">{{$value}}</div>
                        </div>
                    {{end}}
                {{end}}
            {{end}}
        </div>
        
        <div style="text-align: center; margin-top: 30px; padding-top: 20px; border-top: 1px solid #eee; color: #666; font-size: 0.9em;">
            Generated by Dynamic Template Engine
        </div>
    </div>
</body>
</html>`

	// 3. データを埋め込んでレンダリング
	// テンプレート名をデータに追加
	data["templateName"] = templateName

	// 4. エラーハンドリングを実装
	tmpl, err := template.New("dynamicTemplate").Funcs(te.funcMap).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("テンプレートの解析に失敗: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("テンプレートの実行に失敗: %v", err)
	}

	return buf.String(), nil
}

// SaveHTMLToFile メソッドの実装
func (te *TemplateEngine) SaveHTMLToFile(filename, html string) error {
	// 1. ファイルを作成または上書き
	// 2. HTMLコンテンツを書き込み
	// 3. 適切なファイルパーミッションを設定 (0644)
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ファイルの作成に失敗: %v", err)
	}
	defer file.Close()

	// 4. エラーハンドリングを実装
	if _, err := file.WriteString(html); err != nil {
		return fmt.Errorf("HTMLの書き込みに失敗: %v", err)
	}

	return nil
}

// GetTemplateFunction メソッドの実装
func (te *TemplateEngine) GetTemplateFunction() map[string]interface{} {
	// 1. カスタム関数のマップを作成
	return template.FuncMap{
		// 2. 日時フォーマット関数
		"formatDate": formatDate,
		
		// 3. 文字列操作関数
		"truncate":   func(length int, text string) string { return truncateText(text, length) },
		"join":       joinStrings,
		"upper":      strings.ToUpper,
		"lower":      strings.ToLower,
		"title":      strings.Title,
		
		// 4. 数値フォーマット関数
		"formatNumber": func(n int) string {
			return fmt.Sprintf("%d", n)
		},
		
		// 5. 条件判定関数
		"isEven": isEven,
		"isOdd": func(n int) bool {
			return !isEven(n)
		},
		"isEmpty": func(s string) bool {
			return strings.TrimSpace(s) == ""
		},
		"isNotEmpty": func(s string) bool {
			return strings.TrimSpace(s) != ""
		},
	}
}

// formatDate ヘルパー関数の実装
func formatDate(t time.Time) string {
	// 1. 日時を読みやすい形式にフォーマット
	// 2. "2006-01-02 15:04:05" 形式を使用
	if t.IsZero() {
		return "未設定"
	}
	return t.Format("2006-01-02 15:04:05")
}

// truncateText ヘルパー関数の実装
func truncateText(text string, length int) string {
	// 1. 指定された長さでテキストを切り詰め
	if len(text) <= length {
		return text
	}
	
	// 2. 必要に応じて"..."を追加
	// 3. 単語の境界を考慮
	truncated := text[:length]
	lastSpace := strings.LastIndex(truncated, " ")
	if lastSpace > 0 {
		truncated = truncated[:lastSpace]
	}
	return truncated + "..."
}

// joinStrings ヘルパー関数の実装
func joinStrings(strs []string, sep string) string {
	// 1. strings.Join()を使用
	// 2. 空の配列をハンドリング
	if len(strs) == 0 {
		return ""
	}
	return strings.Join(strs, sep)
}

// isEven ヘルパー関数の実装
func isEven(n int) bool {
	// 1. 偶数判定を実装
	// 2. modulo演算子を使用
	return n%2 == 0
}