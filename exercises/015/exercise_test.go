package main

import (
	"strings"
	"testing"
	"time"
)

func TestTemplateEngine(t *testing.T) {
	engine := NewTemplateEngine()
	if engine == nil {
		t.Fatal("NewTemplateEngine returned nil")
	}
}

func TestRenderUserProfile(t *testing.T) {
	engine := NewTemplateEngine()
	
	user := User{
		ID:       1,
		Name:     "Test User",
		Email:    "test@example.com",
		Role:     "Admin",
		Active:   true,
		LastLogin: time.Now(),
	}
	
	html, err := engine.RenderUserProfile(user)
	if err != nil {
		t.Fatalf("RenderUserProfile failed: %v", err)
	}
	
	if html == "" {
		t.Error("Expected non-empty HTML output")
	}
	
	// Check if user data is included
	if !strings.Contains(html, user.Name) {
		t.Error("HTML should contain user name")
	}
}

func TestRenderBlogList(t *testing.T) {
	engine := NewTemplateEngine()
	
	posts := []BlogPost{
		{
			ID:        1,
			Title:     "Test Post",
			Content:   "Test content",
			Author:    "Test Author",
			Published: true,
			Tags:      []string{"test", "blog"},
		},
	}
	
	blogData := BlogData{
		Title: "Test Blog",
		Posts: posts,
		User:  User{Name: "Test User"},
	}
	
	html, err := engine.RenderBlogList(blogData)
	if err != nil {
		t.Fatalf("RenderBlogList failed: %v", err)
	}
	
	if html == "" {
		t.Error("Expected non-empty HTML output")
	}
}

func TestRenderBlogPost(t *testing.T) {
	engine := NewTemplateEngine()
	
	post := BlogPost{
		ID:          1,
		Title:       "Test Post",
		Content:     "This is test content",
		Author:      "Test Author",
		PublishedAt: time.Now(),
		Tags:        []string{"test", "example"},
		Published:   true,
	}
	
	html, err := engine.RenderBlogPost(post)
	if err != nil {
		t.Fatalf("RenderBlogPost failed: %v", err)
	}
	
	if html == "" {
		t.Error("Expected non-empty HTML output")
	}
	
	if !strings.Contains(html, post.Title) {
		t.Error("HTML should contain post title")
	}
}

func TestRenderDashboard(t *testing.T) {
	engine := NewTemplateEngine()
	
	dashboard := Dashboard{
		User:       User{Name: "Admin User"},
		Posts:      []BlogPost{},
		TotalViews: 1000,
		TotalUsers: 50,
		TotalPosts: 10,
		RecentViews: []PageView{
			{Page: "/test", Views: 100, Date: time.Now()},
		},
	}
	
	html, err := engine.RenderDashboard(dashboard)
	if err != nil {
		t.Fatalf("RenderDashboard failed: %v", err)
	}
	
	if html == "" {
		t.Error("Expected non-empty HTML output")
	}
}

func TestRenderErrorPage(t *testing.T) {
	engine := NewTemplateEngine()
	
	html, err := engine.RenderErrorPage(404, "Page Not Found")
	if err != nil {
		t.Fatalf("RenderErrorPage failed: %v", err)
	}
	
	if html == "" {
		t.Error("Expected non-empty HTML output")
	}
	
	if !strings.Contains(html, "404") {
		t.Error("HTML should contain error code")
	}
}

func TestRenderEmailTemplate(t *testing.T) {
	engine := NewTemplateEngine()
	
	emailData := EmailData{
		RecipientName: "Test User",
		Subject:       "Test Email",
		Message:       "This is a test email",
		ActionURL:     "https://example.com/action",
		ActionText:    "Click Here",
	}
	
	html, err := engine.RenderEmailTemplate(emailData)
	if err != nil {
		t.Fatalf("RenderEmailTemplate failed: %v", err)
	}
	
	if html == "" {
		t.Error("Expected non-empty HTML output")
	}
	
	if !strings.Contains(html, emailData.RecipientName) {
		t.Error("HTML should contain recipient name")
	}
}

func TestRenderFromJSON(t *testing.T) {
	engine := NewTemplateEngine()
	
	jsonData := `{"name": "Test", "count": 42}`
	
	html, err := engine.RenderFromJSON("test_template", jsonData)
	if err != nil {
		t.Fatalf("RenderFromJSON failed: %v", err)
	}
	
	if html == "" {
		t.Error("Expected non-empty HTML output")
	}
}

func TestSaveHTMLToFile(t *testing.T) {
	engine := NewTemplateEngine()
	
	html := "<html><body>Test</body></html>"
	filename := "test.html"
	
	err := engine.SaveHTMLToFile(filename, html)
	if err != nil {
		t.Fatalf("SaveHTMLToFile failed: %v", err)
	}
	
	// Clean up
	// Note: In a real implementation, you might want to check if file exists and remove it
}

func TestGetTemplateFunction(t *testing.T) {
	engine := NewTemplateEngine()
	
	functions := engine.GetTemplateFunction()
	if functions == nil {
		t.Error("Expected non-nil function map")
	}
}

func TestHelperFunctions(t *testing.T) {
	// Test formatDate
	now := time.Now()
	formatted := formatDate(now)
	if formatted == "" {
		t.Error("formatDate should return non-empty string")
	}
	
	// Test truncateText
	text := "This is a long text that should be truncated"
	truncated := truncateText(text, 10)
	if len(truncated) > 13 { // 10 + "..." = 13
		t.Error("Text should be truncated")
	}
	
	// Test joinStrings
	strs := []string{"a", "b", "c"}
	joined := joinStrings(strs, ",")
	expected := "a,b,c"
	if joined != expected {
		t.Errorf("Expected %s, got %s", expected, joined)
	}
	
	// Test isEven
	if !isEven(4) {
		t.Error("4 should be even")
	}
	if isEven(5) {
		t.Error("5 should not be even")
	}
}