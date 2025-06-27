# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go learning project containing programming exercises. Each exercise is self-contained in its own directory under `exercises/` with corresponding test files.

## Common Commands

### Running Exercises
```bash
# Run a specific exercise
go run ./exercises/001/exercise.go
go run ./exercises/002/exercise.go
go run ./exercises/003/exercise.go
go run ./exercises/004/exercise.go
go run ./exercises/005/exercise.go
go run ./exercises/006/exercise.go
go run ./exercises/007/exercise.go
go run ./exercises/008/exercise.go
go run ./exercises/009/exercise.go
cd exercises/010 && go run exercise.go  # Special case: has dependencies
```

### Testing
```bash
# Test a specific exercise
go test ./exercises/001 -v
go test ./exercises/002 -v
go test ./exercises/003 -v
go test ./exercises/004 -v
go test ./exercises/005 -v
go test ./exercises/006 -v
go test ./exercises/007 -v
go test ./exercises/008 -v
go test ./exercises/009 -v
cd exercises/010 && go test -v  # Special case: has dependencies

# Test all exercises (excluding 010 due to dependencies)
go test ./exercises/001 ./exercises/002 ./exercises/003 ./exercises/004 ./exercises/005 ./exercises/006 ./exercises/007 ./exercises/008 ./exercises/009 -v

# Run tests with coverage
go test ./exercises/001 ./exercises/002 ./exercises/003 ./exercises/004 ./exercises/005 ./exercises/006 ./exercises/007 ./exercises/008 ./exercises/009 -v -cover
```

### Development
```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...
```

## Project Structure

The codebase follows a simple exercise-based structure:
- Each exercise lives in `exercises/XXX/` where XXX is the exercise number
- Each exercise has an `exercise.go` file containing the implementation
- Each exercise has an `exercise_test.go` file containing tests
- Exercises are designed to be completed sequentially

## Exercise Status

- Exercise 001: Find numbers divisible by 7 but not 5 (Complete)
- Exercise 002: Calculate factorial with error handling (Complete)  
- Exercise 003: Generate integer-to-square map (Complete)
- Exercise 004: Go concurrency with goroutines, channels, and interfaces (Complete)
- Exercise 005: HTTP API server with JSON processing (Complete)
- Exercise 006: File operations and error handling (Incomplete - needs implementation)
- Exercise 007: Slice operations and sorting algorithms (Incomplete - needs implementation)
- Exercise 008: Advanced concurrency patterns with context (Incomplete - needs implementation)
- Exercise 009: Reflection and custom types (Incomplete - needs implementation)
- Exercise 010: Database operations and CRUD with SQLite (Incomplete - needs implementation)

## Key Development Notes

1. Most exercises use only Go's standard library (no external dependencies)
2. Exercise 010 uses SQLite driver (github.com/mattn/go-sqlite3) - has its own go.mod
3. Uses Go's standard testing framework - no test runners or additional tools needed
4. Each exercise is independent - modifications to one exercise don't affect others
5. When implementing exercises, ensure all tests pass before considering it complete