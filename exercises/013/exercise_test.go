package main

import (
	"testing"
)

func TestTextProcessor(t *testing.T) {
	processor := NewTextProcessor()
	if processor == nil {
		t.Fatal("NewTextProcessor returned nil")
	}
}

func TestExtractEmails(t *testing.T) {
	processor := NewTextProcessor()
	
	text := "Contact us at support@example.com or sales@company.org for assistance."
	emails := processor.ExtractEmails(text)
	
	if len(emails) != 2 {
		t.Errorf("Expected 2 emails, got %d", len(emails))
	}
	
	expectedEmails := []string{"support@example.com", "sales@company.org"}
	for i, expected := range expectedEmails {
		if i < len(emails) && emails[i] != expected {
			t.Errorf("Expected email %s, got %s", expected, emails[i])
		}
	}
}

func TestExtractPhoneNumbers(t *testing.T) {
	processor := NewTextProcessor()
	
	text := "Call us at +1-555-123-4567 or (555) 987-6543"
	phones := processor.ExtractPhoneNumbers(text)
	
	if len(phones) == 0 {
		t.Error("Expected to find phone numbers")
	}
}

func TestExtractURLs(t *testing.T) {
	processor := NewTextProcessor()
	
	text := "Visit https://www.example.com and http://blog.company.org"
	urls := processor.ExtractURLs(text)
	
	if len(urls) != 2 {
		t.Errorf("Expected 2 URLs, got %d", len(urls))
	}
}

func TestExtractHashtags(t *testing.T) {
	processor := NewTextProcessor()
	
	text := "Learning #golang and #programming is fun! #coding"
	hashtags := processor.ExtractHashtags(text)
	
	if len(hashtags) != 3 {
		t.Errorf("Expected 3 hashtags, got %d", len(hashtags))
	}
}

func TestExtractMentions(t *testing.T) {
	processor := NewTextProcessor()
	
	text := "Thanks @john_doe and @jane_smith for the help!"
	mentions := processor.ExtractMentions(text)
	
	if len(mentions) != 2 {
		t.Errorf("Expected 2 mentions, got %d", len(mentions))
	}
}

func TestParseLogEntry(t *testing.T) {
	processor := NewTextProcessor()
	
	logLine := "2024-01-15 10:30:45 [INFO] User login successful: user@example.com from 192.168.1.100"
	logEntry := processor.ParseLogEntry(logLine)
	
	if logEntry == nil {
		t.Fatal("ParseLogEntry returned nil")
	}
	
	if logEntry.Level != "INFO" {
		t.Errorf("Expected level INFO, got %s", logEntry.Level)
	}
}

func TestCleanText(t *testing.T) {
	processor := NewTextProcessor()
	
	dirtyText := "  Hello,   World!!!   This is    a   test...   "
	cleaned := processor.CleanText(dirtyText)
	
	if cleaned == dirtyText {
		t.Error("Text should be cleaned")
	}
	
	if len(cleaned) >= len(dirtyText) {
		t.Error("Cleaned text should be shorter")
	}
}

func TestEmailValidation(t *testing.T) {
	processor := NewTextProcessor()
	
	validEmails := []string{
		"user@example.com",
		"test.email@domain.org",
		"user+tag@example.co.uk",
	}
	
	invalidEmails := []string{
		"invalid-email",
		"@example.com",
		"user@",
		"user.example.com",
	}
	
	for _, email := range validEmails {
		if !processor.IsValidEmail(email) {
			t.Errorf("Expected %s to be valid email", email)
		}
	}
	
	for _, email := range invalidEmails {
		if processor.IsValidEmail(email) {
			t.Errorf("Expected %s to be invalid email", email)
		}
	}
}

func TestURLValidation(t *testing.T) {
	processor := NewTextProcessor()
	
	validURLs := []string{
		"https://www.example.com",
		"http://example.org",
		"https://subdomain.example.com/path",
	}
	
	invalidURLs := []string{
		"not-a-url",
		"ftp://example.com",
		"www.example.com",
	}
	
	for _, url := range validURLs {
		if !processor.IsValidURL(url) {
			t.Errorf("Expected %s to be valid URL", url)
		}
	}
	
	for _, url := range invalidURLs {
		if processor.IsValidURL(url) {
			t.Errorf("Expected %s to be invalid URL", url)
		}
	}
}

func TestPhoneValidation(t *testing.T) {
	processor := NewTextProcessor()
	
	validPhones := []string{
		"+1-555-123-4567",
		"(555) 987-6543",
		"555.123.4567",
	}
	
	invalidPhones := []string{
		"invalid-phone",
		"123",
		"abc-def-ghij",
	}
	
	for _, phone := range validPhones {
		if !processor.IsValidPhoneNumber(phone) {
			t.Errorf("Expected %s to be valid phone", phone)
		}
	}
	
	for _, phone := range invalidPhones {
		if processor.IsValidPhoneNumber(phone) {
			t.Errorf("Expected %s to be invalid phone", phone)
		}
	}
}

func TestPasswordValidation(t *testing.T) {
	processor := NewTextProcessor()
	
	strongPasswords := []string{
		"Strong_Password123!",
		"MySecure@Pass2024",
		"C0mpl3x&Password",
	}
	
	weakPasswords := []string{
		"weak",
		"password",
		"12345678",
		"Password",
	}
	
	for _, password := range strongPasswords {
		if !processor.IsStrongPassword(password) {
			t.Errorf("Expected %s to be strong password", password)
		}
	}
	
	for _, password := range weakPasswords {
		if processor.IsStrongPassword(password) {
			t.Errorf("Expected %s to be weak password", password)
		}
	}
}

func TestCaseConversions(t *testing.T) {
	processor := NewTextProcessor()
	
	original := "Hello World Test"
	
	snakeCase := processor.ToSnakeCase(original)
	if snakeCase == original {
		t.Error("Snake case conversion failed")
	}
	
	camelCase := processor.ToCamelCase(original)
	if camelCase == original {
		t.Error("Camel case conversion failed")
	}
	
	kebabCase := processor.ToKebabCase(original)
	if kebabCase == original {
		t.Error("Kebab case conversion failed")
	}
}

func TestAnalyzeText(t *testing.T) {
	processor := NewTextProcessor()
	
	text := "Hello world! This is a test. Contact us at test@example.com."
	stats := processor.AnalyzeText(text)
	
	if stats.CharCount == 0 {
		t.Error("Character count should be greater than 0")
	}
	
	if stats.WordCount == 0 {
		t.Error("Word count should be greater than 0")
	}
	
	if stats.EmailCount == 0 {
		t.Error("Should detect email in text")
	}
}

func TestHelperFunctions(t *testing.T) {
	text := "Hello 123 World 456"
	
	// Test RemovePattern
	result := RemovePattern(text, `\d+`)
	if result == text {
		t.Error("RemovePattern should remove digits")
	}
	
	// Test ReplacePattern
	result = ReplacePattern(text, `\d+`, "X")
	if result == text {
		t.Error("ReplacePattern should replace digits")
	}
	
	// Test CountMatches
	count := CountMatches(text, `\d+`)
	if count != 2 {
		t.Errorf("Expected 2 digit matches, got %d", count)
	}
}