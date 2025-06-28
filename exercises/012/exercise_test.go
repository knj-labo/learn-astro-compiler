package main

import (
	"os"
	"testing"
)

func TestCSVProcessor(t *testing.T) {
	processor := NewCSVProcessor()
	if processor == nil {
		t.Fatal("NewCSVProcessor returned nil")
	}
}

func TestReadWriteCustomersCSV(t *testing.T) {
	processor := NewCSVProcessor()
	
	// Test data
	customers := []Customer{
		{ID: 1, Name: "Alice", Email: "alice@test.com", Age: 30, City: "Tokyo"},
		{ID: 2, Name: "Bob", Email: "bob@test.com", Age: 25, City: "Osaka"},
	}
	
	filename := "test_customers.csv"
	defer os.Remove(filename)
	
	// Test write
	err := processor.WriteCustomersToCSV(filename, customers)
	if err != nil {
		t.Fatalf("WriteCustomersToCSV failed: %v", err)
	}
	
	// Test read
	readCustomers, err := processor.ReadCustomersFromCSV(filename)
	if err != nil {
		t.Fatalf("ReadCustomersFromCSV failed: %v", err)
	}
	
	if len(readCustomers) != len(customers) {
		t.Errorf("Expected %d customers, got %d", len(customers), len(readCustomers))
	}
	
	// Verify first customer
	if len(readCustomers) > 0 {
		if readCustomers[0].Name != customers[0].Name {
			t.Errorf("Expected name %s, got %s", customers[0].Name, readCustomers[0].Name)
		}
	}
}

func TestFilterByAge(t *testing.T) {
	processor := NewCSVProcessor()
	
	customers := []Customer{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Carol", Age: 35},
	}
	
	filtered := processor.FilterByAge(customers, 30)
	
	if len(filtered) != 2 {
		t.Errorf("Expected 2 customers aged 30+, got %d", len(filtered))
	}
}

func TestGroupByCity(t *testing.T) {
	processor := NewCSVProcessor()
	
	customers := []Customer{
		{ID: 1, Name: "Alice", City: "Tokyo"},
		{ID: 2, Name: "Bob", City: "Osaka"},
		{ID: 3, Name: "Carol", City: "Tokyo"},
	}
	
	groups := processor.GroupByCity(customers)
	
	if len(groups) != 2 {
		t.Errorf("Expected 2 city groups, got %d", len(groups))
	}
	
	if len(groups["Tokyo"]) != 2 {
		t.Errorf("Expected 2 customers in Tokyo, got %d", len(groups["Tokyo"]))
	}
	
	if len(groups["Osaka"]) != 1 {
		t.Errorf("Expected 1 customer in Osaka, got %d", len(groups["Osaka"]))
	}
}

func TestCalculateStats(t *testing.T) {
	processor := NewCSVProcessor()
	
	customers := []Customer{
		{ID: 1, Name: "Alice", Age: 30, City: "Tokyo"},
		{ID: 2, Name: "Bob", Age: 20, City: "Osaka"},
		{ID: 3, Name: "Carol", Age: 40, City: "Tokyo"},
	}
	
	stats := processor.CalculateStats(customers)
	
	if stats.TotalCustomers != 3 {
		t.Errorf("Expected 3 total customers, got %d", stats.TotalCustomers)
	}
	
	if stats.AverageAge != 30.0 {
		t.Errorf("Expected average age 30.0, got %f", stats.AverageAge)
	}
	
	if stats.MinAge != 20 {
		t.Errorf("Expected min age 20, got %d", stats.MinAge)
	}
	
	if stats.MaxAge != 40 {
		t.Errorf("Expected max age 40, got %d", stats.MaxAge)
	}
	
	if stats.CitiesCount != 2 {
		t.Errorf("Expected 2 cities, got %d", stats.CitiesCount)
	}
}

func TestFileOperations(t *testing.T) {
	processor := NewCSVProcessor()
	
	filename := "test_file.txt"
	defer os.Remove(filename)
	
	// Test file doesn't exist initially
	if processor.FileExists(filename) {
		t.Error("File should not exist initially")
	}
	
	// Create file
	customers := []Customer{{ID: 1, Name: "Test"}}
	err := processor.WriteCustomersToCSV(filename, customers)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	
	// Test file exists
	if !processor.FileExists(filename) {
		t.Error("File should exist after creation")
	}
	
	// Test file size
	size, err := processor.GetFileSize(filename)
	if err != nil {
		t.Fatalf("GetFileSize failed: %v", err)
	}
	
	if size <= 0 {
		t.Error("File size should be greater than 0")
	}
}

func TestGenerateReport(t *testing.T) {
	processor := NewCSVProcessor()
	
	customers := []Customer{
		{ID: 1, Name: "Alice", Age: 30, City: "Tokyo"},
	}
	
	stats := CustomerStats{
		TotalCustomers: 1,
		AverageAge:     30.0,
		MinAge:         30,
		MaxAge:         30,
		CitiesCount:    1,
	}
	
	filename := "test_report.txt"
	defer os.Remove(filename)
	
	err := processor.GenerateReport(filename, customers, stats)
	if err != nil {
		t.Fatalf("GenerateReport failed: %v", err)
	}
	
	if !processor.FileExists(filename) {
		t.Error("Report file should exist")
	}
}

func TestTextFileOperations(t *testing.T) {
	filename := "test_text.txt"
	defer os.Remove(filename)
	
	content := "Hello, World!"
	
	// Test write
	err := WriteTextFile(filename, content)
	if err != nil {
		t.Fatalf("WriteTextFile failed: %v", err)
	}
	
	// Test read
	readContent, err := ReadTextFile(filename)
	if err != nil {
		t.Fatalf("ReadTextFile failed: %v", err)
	}
	
	if readContent != content {
		t.Errorf("Expected content %s, got %s", content, readContent)
	}
	
	// Test append
	appendContent := "\nGoodbye!"
	err = AppendToFile(filename, appendContent)
	if err != nil {
		t.Fatalf("AppendToFile failed: %v", err)
	}
	
	// Verify append
	finalContent, err := ReadTextFile(filename)
	if err != nil {
		t.Fatalf("ReadTextFile after append failed: %v", err)
	}
	
	expected := content + appendContent
	if finalContent != expected {
		t.Errorf("Expected final content %s, got %s", expected, finalContent)
	}
}

func TestInvalidFile(t *testing.T) {
	processor := NewCSVProcessor()
	
	// Test reading non-existent file
	_, err := processor.ReadCustomersFromCSV("non_existent.csv")
	if err == nil {
		t.Error("Expected error when reading non-existent file")
	}
	
	// Test getting size of non-existent file
	_, err = processor.GetFileSize("non_existent.txt")
	if err == nil {
		t.Error("Expected error when getting size of non-existent file")
	}
}