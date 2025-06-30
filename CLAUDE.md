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
go run ./exercises/011/exercise.go
go run ./exercises/012/exercise.go
go run ./exercises/013/exercise.go
go run ./exercises/014/exercise.go
go run ./exercises/015/exercise.go
go run ./exercises/016/exercise.go
go run ./exercises/017/exercise.go
go run ./exercises/018/exercise.go
go run ./exercises/019/exercise.go
go run ./exercises/020/exercise.go
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
go test ./exercises/011 -v
go test ./exercises/012 -v
go test ./exercises/013 -v
go test ./exercises/014 -v
go test ./exercises/015 -v
go test ./exercises/016 -v
go test ./exercises/017 -v
go test ./exercises/018 -v
go test ./exercises/019 -v
go test ./exercises/020 -v

# Test all exercises (excluding 010 due to dependencies)
go test ./exercises/001 ./exercises/002 ./exercises/003 ./exercises/004 ./exercises/005 ./exercises/006 ./exercises/007 ./exercises/008 ./exercises/009 ./exercises/011 ./exercises/012 ./exercises/013 ./exercises/014 ./exercises/015 ./exercises/016 ./exercises/017 ./exercises/018 ./exercises/019 ./exercises/020 -v

# Run tests with coverage
go test ./exercises/001 ./exercises/002 ./exercises/003 ./exercises/004 ./exercises/005 ./exercises/006 ./exercises/007 ./exercises/008 ./exercises/009 ./exercises/011 ./exercises/012 ./exercises/013 ./exercises/014 ./exercises/015 ./exercises/016 ./exercises/017 ./exercises/018 ./exercises/019 ./exercises/020 -v -cover
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
- Exercise 010: Database operations and CRUD with SQLite (Complete)
- Exercise 011: JSON processing and RESTful API (Complete)
- Exercise 012: File I/O and CSV processing (Complete)
- Exercise 013: Regular expressions and text processing (Incomplete - needs implementation)
- Exercise 014: Concurrency with worker pools (Incomplete - needs implementation)
- Exercise 015: Template engine and HTML generation (Incomplete - needs implementation)
- Exercise 016: Command-line tool with flags (Incomplete - needs implementation)
- Exercise 017: Middleware and HTTP handlers (Incomplete - needs implementation)
- Exercise 018: Testing and benchmarking (Incomplete - needs implementation)
- Exercise 019: Encryption and security (Incomplete - needs implementation)
- Exercise 020: WebSocket real-time communication (Incomplete - needs implementation)

## Key Development Notes

1. Most exercises use only Go's standard library (no external dependencies)
2. Exercise 010 uses SQLite driver (github.com/mattn/go-sqlite3) - has its own go.mod
3. Uses Go's standard testing framework - no test runners or additional tools needed
4. Each exercise is independent - modifications to one exercise don't affect others
5. When implementing exercises, ensure all tests pass before considering it complete

## Implementation Guidelines

### Code Style and Comments
When implementing exercises, focus on clean, well-documented code:

1. **Remove TODO comments**: Delete original TODO comments and hints after implementation
2. **Add inline implementation comments**: Place step-by-step comments within the implementation to explain the logic
3. **Format for clarity**: Structure comments to clearly explain what each section of code does

Example pattern:
```go
func ExampleFunction() {
    // 1. Step one description
    actualImplementationCode()
    
    // 2. Step two description  
    moreImplementationCode()
}
```

### Testing and Validation
- Always run tests after implementation: `go test ./exercises/XXX -v`
- Ensure all tests pass before marking an exercise as complete
- Fix any linting issues that arise during development
- Remove unused imports from test files if needed

### Error Handling
- Implement proper error handling for all operations
- Return appropriate HTTP status codes for web APIs
- Use descriptive error messages for debugging

### JSON and HTTP APIs
- Set proper Content-Type headers for JSON responses
- Validate input data before processing
- Follow RESTful conventions for endpoint design
- Use standard HTTP status codes (200, 201, 404, 400, etc.)