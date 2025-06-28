package main

import (
	"strings"
	"testing"
)

func TestCLIApp(t *testing.T) {
	app := NewCLIApp()
	if app == nil {
		t.Fatal("NewCLIApp returned nil")
	}
}

func TestRegisterCommand(t *testing.T) {
	app := NewCLIApp()
	
	// Create a test command
	testCmd := &FileProcessCommand{}
	
	app.RegisterCommand(testCmd)
	
	// Verify command was registered
	if len(app.commands) == 0 {
		t.Error("Command should be registered")
	}
}

func TestRunWithHelp(t *testing.T) {
	app := NewCLIApp()
	
	// Test help command
	args := []string{"program", "help"}
	err := app.Run(args)
	
	if err != nil {
		t.Errorf("Help command should not return error: %v", err)
	}
}

func TestRunWithInvalidCommand(t *testing.T) {
	app := NewCLIApp()
	
	// Test invalid command
	args := []string{"program", "invalid"}
	err := app.Run(args)
	
	if err == nil {
		t.Error("Invalid command should return error")
	}
}

func TestFileProcessCommand(t *testing.T) {
	cmd := &FileProcessCommand{}
	
	// Test command properties
	name := cmd.Name()
	if name == "" {
		t.Error("Command name should not be empty")
	}
	
	description := cmd.Description()
	if description == "" {
		t.Error("Command description should not be empty")
	}
	
	// Test command execution with minimal args
	args := []string{"fileprocess", "-help"}
	err := cmd.Run(args)
	
	// Should not panic and might return error for missing flags
	_ = err
}

func TestServerCommand(t *testing.T) {
	cmd := &ServerCommand{}
	
	// Test command properties
	name := cmd.Name()
	if name == "" {
		t.Error("Command name should not be empty")
	}
	
	description := cmd.Description()
	if description == "" {
		t.Error("Command description should not be empty")
	}
	
	// Test command execution with help flag
	args := []string{"server", "-help"}
	err := cmd.Run(args)
	
	// Should not panic
	_ = err
}

func TestConfigCommand(t *testing.T) {
	cmd := &ConfigCommand{}
	
	// Test command properties
	name := cmd.Name()
	if name == "" {
		t.Error("Command name should not be empty")
	}
	
	description := cmd.Description()
	if description == "" {
		t.Error("Command description should not be empty")
	}
	
	// Test command execution with help flag
	args := []string{"config", "-help"}
	err := cmd.Run(args)
	
	// Should not panic
	_ = err
}

func TestShowHelp(t *testing.T) {
	app := NewCLIApp()
	
	// Register some commands
	app.RegisterCommand(&FileProcessCommand{})
	app.RegisterCommand(&ServerCommand{})
	app.RegisterCommand(&ConfigCommand{})
	
	// This should not panic
	app.ShowHelp()
}

func TestCommandInterface(t *testing.T) {
	commands := []Command{
		&FileProcessCommand{},
		&ServerCommand{},
		&ConfigCommand{},
	}
	
	for _, cmd := range commands {
		// Verify interface implementation
		name := cmd.Name()
		description := cmd.Description()
		
		if name == "" {
			t.Errorf("Command %T should have a name", cmd)
		}
		
		if description == "" {
			t.Errorf("Command %T should have a description", cmd)
		}
		
		// Test Run method exists (will test with help flag)
		err := cmd.Run([]string{name, "-help"})
		_ = err // Don't require specific error handling
	}
}

func TestAppWithMultipleCommands(t *testing.T) {
	app := NewCLIApp()
	
	// Register multiple commands
	commands := []Command{
		&FileProcessCommand{},
		&ServerCommand{},
		&ConfigCommand{},
	}
	
	for _, cmd := range commands {
		app.RegisterCommand(cmd)
	}
	
	// Test that all commands are registered
	if len(app.commands) != len(commands) {
		t.Errorf("Expected %d commands, got %d", len(commands), len(app.commands))
	}
}

func TestEmptyArgs(t *testing.T) {
	app := NewCLIApp()
	
	// Test with no arguments (should show help or handle gracefully)
	args := []string{"program"}
	err := app.Run(args)
	
	// Should handle empty args gracefully
	_ = err
}

func TestCommandNameUniqueness(t *testing.T) {
	app := NewCLIApp()
	
	cmd1 := &FileProcessCommand{}
	cmd2 := &FileProcessCommand{}
	
	app.RegisterCommand(cmd1)
	app.RegisterCommand(cmd2)
	
	// Should handle duplicate command names appropriately
	// (either reject or overwrite)
	_ = app.commands
}