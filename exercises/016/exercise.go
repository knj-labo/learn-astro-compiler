package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

/*
Exercise 016: コマンドラインツールとフラグ処理

このエクササイズでは、flagパッケージを使用してコマンドライン引数を処理し、
実用的なCLIツールを作成することを学びます：

1. フラグの定義と解析
   - 文字列、整数、ブール値フラグ
   - デフォルト値の設定
   - ヘルプメッセージの生成

2. サブコマンドの実装
   - 複数のコマンドモード
   - コマンド固有のフラグ
   - 階層的なコマンド構造

3. 実用的な機能
   - ファイル処理
   - 設定ファイルの読み込み
   - 出力フォーマットの制御

期待される動作:
- ユーザーフレンドリーなCLI
- 豊富なヘルプ情報
- エラーハンドリング
*/

func main() {
	fmt.Println("Exercise 016: Command-line Tool and Flag Processing")
	
	// CLIアプリケーションを初期化
	app := NewCLIApp()
	
	// アプリケーションを実行
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

// CLIApp構造体
type CLIApp struct {
	commands map[string]Command
}

// Command インターフェース
type Command interface {
	Name() string
	Description() string
	Run(args []string) error
}

// FileProcessCommand構造体
type FileProcessCommand struct{}

// ServerCommand構造体
type ServerCommand struct{}

// ConfigCommand構造体
type ConfigCommand struct{}

// NewCLIApp関数の実装
func NewCLIApp() *CLIApp {
	// 1. CLIApp構造体を初期化
	app := &CLIApp{
		commands: make(map[string]Command),
	}
	
	// 2. 利用可能なコマンドを登録
	app.RegisterCommand(&FileProcessCommand{})
	app.RegisterCommand(&ServerCommand{})
	app.RegisterCommand(&ConfigCommand{})
	
	// 3. ヘルプコマンドを追加（実際にはRunメソッドで処理）
	return app
}

// Run メソッドの実装
func (app *CLIApp) Run(args []string) error {
	// 1. 引数を解析してコマンドを特定
	if len(args) < 2 {
		// 引数が不足している場合はヘルプを表示
		app.ShowHelp()
		return nil
	}
	
	commandName := args[1]
	
	// 2. 適切なコマンドを実行
	// 4. ヘルプメッセージの表示
	if commandName == "help" || commandName == "-help" || commandName == "--help" {
		app.ShowHelp()
		return nil
	}
	
	// 3. エラーハンドリングを実装
	command, exists := app.commands[commandName]
	if !exists {
		fmt.Printf("Error: Unknown command '%s'\n\n", commandName)
		app.ShowHelp()
		return fmt.Errorf("unknown command: %s", commandName)
	}
	
	// コマンドを実行（残りの引数を渡す）
	return command.Run(args[1:])
}

// RegisterCommand メソッドの実装
func (app *CLIApp) RegisterCommand(cmd Command) {
	// 1. コマンドをマップに登録
	// 2. 名前の重複チェック（上書きを許可）
	if cmd != nil {
		app.commands[cmd.Name()] = cmd
	}
}

// ShowHelp メソッドの実装
func (app *CLIApp) ShowHelp() {
	// 1. 使用方法を表示
	fmt.Println("CLI Tool - Command Line Interface")
	fmt.Println("")
	fmt.Println("Usage: program <command> [options]")
	fmt.Println("")
	
	// 2. 利用可能なコマンド一覧
	fmt.Println("Available Commands:")
	for name, cmd := range app.commands {
		fmt.Printf("  %-12s %s\n", name, cmd.Description())
	}
	fmt.Println("  help         Show this help message")
	fmt.Println("")
	
	// 3. 例の表示
	fmt.Println("Examples:")
	fmt.Println("  program fileprocess -input file.txt -output result.txt")
	fmt.Println("  program server -port 8080 -host localhost")
	fmt.Println("  program config show")
	fmt.Println("  program help")
	fmt.Println("")
}

// FileProcessCommand の実装

// Name メソッドの実装
func (fpc *FileProcessCommand) Name() string {
	return "fileprocess"
}

// Description メソッドの実装
func (fpc *FileProcessCommand) Description() string {
	return "Process files with various operations (copy, transform, analyze)"
}

// Run メソッドの実装
func (fpc *FileProcessCommand) Run(args []string) error {
	// 1. flagパッケージでフラグを定義
	fs := flag.NewFlagSet("fileprocess", flag.ContinueOnError)
	
	// 2. 入力ファイル、出力ファイル、処理モードのフラグ
	inputFile := fs.String("input", "", "Input file path (required)")
	outputFile := fs.String("output", "", "Output file path (required)")
	mode := fs.String("mode", "copy", "Processing mode: copy, uppercase, lowercase, wordcount")
	verbose := fs.Bool("verbose", false, "Enable verbose output")
	
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}
	
	// 必須フラグのチェック
	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Error: Both -input and -output flags are required")
		fs.Usage()
		return fmt.Errorf("missing required flags")
	}
	
	if *verbose {
		fmt.Printf("Processing file: %s -> %s (mode: %s)\n", *inputFile, *outputFile, *mode)
	}
	
	// 3. ファイル処理ロジックを実装
	// 4. 進捗表示機能
	return fpc.processFile(*inputFile, *outputFile, *mode, *verbose)
}

// processFile ファイル処理のヘルパーメソッド
func (fpc *FileProcessCommand) processFile(inputPath, outputPath, mode string, verbose bool) error {
	// 入力ファイルの存在確認
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", inputPath)
	}
	
	// 出力ディレクトリの作成
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}
	
	if verbose {
		fmt.Printf("Reading file: %s\n", inputPath)
	}
	
	// ファイル読み込み
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer inputFile.Close()
	
	// 出力ファイル作成
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()
	
	// モードに応じた処理
	switch mode {
	case "copy":
		return fpc.copyFile(inputFile, outputFile, verbose)
	case "uppercase":
		return fpc.transformFile(inputFile, outputFile, strings.ToUpper, verbose)
	case "lowercase":
		return fpc.transformFile(inputFile, outputFile, strings.ToLower, verbose)
	case "wordcount":
		return fpc.wordCountFile(inputFile, outputFile, verbose)
	default:
		return fmt.Errorf("unknown processing mode: %s", mode)
	}
}

// copyFile ファイルコピー
func (fpc *FileProcessCommand) copyFile(src io.Reader, dst io.Writer, verbose bool) error {
	if verbose {
		fmt.Println("Copying file...")
	}
	_, err := io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}
	if verbose {
		fmt.Println("File copied successfully")
	}
	return nil
}

// transformFile テキスト変換
func (fpc *FileProcessCommand) transformFile(src io.Reader, dst io.Writer, transform func(string) string, verbose bool) error {
	if verbose {
		fmt.Println("Transforming file...")
	}
	
	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dst)
	defer writer.Flush()
	
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		transformedLine := transform(line)
		if _, err := writer.WriteString(transformedLine + "\n"); err != nil {
			return fmt.Errorf("failed to write transformed line: %v", err)
		}
		lineCount++
		if verbose && lineCount%100 == 0 {
			fmt.Printf("Processed %d lines\n", lineCount)
		}
	}
	
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	
	if verbose {
		fmt.Printf("File transformed successfully (%d lines)\n", lineCount)
	}
	return nil
}

// wordCountFile 単語カウント
func (fpc *FileProcessCommand) wordCountFile(src io.Reader, dst io.Writer, verbose bool) error {
	if verbose {
		fmt.Println("Counting words...")
	}
	
	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dst)
	defer writer.Flush()
	
	lineCount := 0
	wordCount := 0
	charCount := 0
	
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		lineCount++
		wordCount += len(words)
		charCount += len(line)
	}
	
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	
	// 結果を出力ファイルに書き込み
	result := fmt.Sprintf("File Statistics:\nLines: %d\nWords: %d\nCharacters: %d\n", lineCount, wordCount, charCount)
	if _, err := writer.WriteString(result); err != nil {
		return fmt.Errorf("failed to write statistics: %v", err)
	}
	
	if verbose {
		fmt.Printf("Word count completed: %d lines, %d words, %d characters\n", lineCount, wordCount, charCount)
	}
	return nil
}

// ServerCommand の実装

// Name メソッドの実装
func (sc *ServerCommand) Name() string {
	return "server"
}

// Description メソッドの実装
func (sc *ServerCommand) Description() string {
	return "Start HTTP server with configurable host and port"
}

// Run メソッドの実装
func (sc *ServerCommand) Run(args []string) error {
	// 1. ポート番号、ホスト、設定ファイルのフラグ
	fs := flag.NewFlagSet("server", flag.ContinueOnError)
	
	port := fs.String("port", "8080", "Server port")
	host := fs.String("host", "localhost", "Server host")
	configFile := fs.String("config", "", "Configuration file path")
	verbose := fs.Bool("verbose", false, "Enable verbose logging")
	
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}
	
	if *verbose {
		fmt.Printf("Starting server on %s:%s\n", *host, *port)
		if *configFile != "" {
			fmt.Printf("Using config file: %s\n", *configFile)
		}
	}
	
	// 2. HTTPサーバーの起動
	return sc.startServer(*host, *port, *configFile, *verbose)
}

// startServer HTTPサーバーの実装
func (sc *ServerCommand) startServer(host, port, configFile string, verbose bool) error {
	// HTTPハンドラーを設定
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if verbose {
			fmt.Printf("Request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		}
		
		if r.URL.Path == "/" {
			fmt.Fprintf(w, "Hello from CLI Server!\nTime: %s\n", time.Now().Format(time.RFC3339))
		} else if r.URL.Path == "/health" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"status":"ok","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
		} else {
			http.NotFound(w, r)
		}
	})
	
	// 設定ファイルから追加設定を読み込み（あれば）
	if configFile != "" {
		if err := sc.loadConfig(configFile, verbose); err != nil {
			fmt.Printf("Warning: Failed to load config file: %v\n", err)
		}
	}
	
	// サーバー設定
	server := &http.Server{
		Addr:         host + ":" + port,
		Handler:      nil, // デフォルトのServeMuxを使用
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// 3. グレースフルシャットダウン
	// シグナルハンドリング設定
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// サーバーを別のgoroutineで開始
	go func() {
		fmt.Printf("Server starting on http://%s:%s\n", host, port)
		fmt.Println("Press Ctrl+C to stop the server")
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v\n", err)
		}
	}()
	
	// シグナルを待機
	<-sigChan
	fmt.Println("\nShutting down server...")
	
	// グレースフルシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}
	
	fmt.Println("Server stopped gracefully")
	return nil
}

// loadConfig 設定ファイルの読み込み
func (sc *ServerCommand) loadConfig(configFile string, verbose bool) error {
	if verbose {
		fmt.Printf("Loading configuration from: %s\n", configFile)
	}
	
	// 設定ファイルの存在確認
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist: %s", configFile)
	}
	
	// 簡単な設定ファイル読み込みの例
	// 実際の実装では、JSONやYAMLパーサーを使用する
	file, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()
	
	// ここでは設定ファイルの存在確認のみ実施
	if verbose {
		fmt.Println("Configuration loaded successfully")
	}
	return nil
}

// ConfigCommand の実装

// Name メソッドの実装
func (cc *ConfigCommand) Name() string {
	return "config"
}

// Description メソッドの実装
func (cc *ConfigCommand) Description() string {
	return "Manage configuration files (show, set, get commands)"
}

// Run メソッドの実装
func (cc *ConfigCommand) Run(args []string) error {
	// 1. 設定の表示、設定、取得のサブコマンド
	if len(args) < 2 {
		fmt.Println("Usage: config <subcommand> [options]")
		fmt.Println("Subcommands:")
		fmt.Println("  show    - Show all configuration")
		fmt.Println("  get     - Get specific configuration value")
		fmt.Println("  set     - Set configuration value")
		fmt.Println("  init    - Initialize configuration file")
		return fmt.Errorf("missing subcommand")
	}
	
	subcommand := args[1]
	
	switch subcommand {
	case "show":
		return cc.showConfig(args[2:])
	case "get":
		return cc.getConfig(args[2:])
	case "set":
		return cc.setConfig(args[2:])
	case "init":
		return cc.initConfig(args[2:])
	default:
		return fmt.Errorf("unknown subcommand: %s", subcommand)
	}
}

// 設定管理用の構造体
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Debug    bool   `json:"debug"`
	LogLevel string `json:"log_level"`
}

// getConfigFilePath 設定ファイルのパスを取得
func (cc *ConfigCommand) getConfigFilePath() string {
	return "cli-config.json"
}

// showConfig 設定の表示
func (cc *ConfigCommand) showConfig(args []string) error {
	fs := flag.NewFlagSet("config show", flag.ContinueOnError)
	configFile := fs.String("file", cc.getConfigFilePath(), "Configuration file path")
	
	if err := fs.Parse(args); err != nil {
		return err
	}
	
	// 2. 設定ファイルの管理
	config, err := cc.loadConfigFile(*configFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}
	
	fmt.Printf("Configuration from %s:\n", *configFile)
	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Port: %s\n", config.Port)
	fmt.Printf("Debug: %t\n", config.Debug)
	fmt.Printf("Log Level: %s\n", config.LogLevel)
	
	return nil
}

// getConfig 特定の設定値を取得
func (cc *ConfigCommand) getConfig(args []string) error {
	fs := flag.NewFlagSet("config get", flag.ContinueOnError)
	configFile := fs.String("file", cc.getConfigFilePath(), "Configuration file path")
	key := fs.String("key", "", "Configuration key to get (required)")
	
	if err := fs.Parse(args); err != nil {
		return err
	}
	
	if *key == "" {
		fmt.Println("Error: -key flag is required")
		fs.Usage()
		return fmt.Errorf("missing key")
	}
	
	config, err := cc.loadConfigFile(*configFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}
	
	switch *key {
	case "host":
		fmt.Println(config.Host)
	case "port":
		fmt.Println(config.Port)
	case "debug":
		fmt.Println(config.Debug)
	case "log_level":
		fmt.Println(config.LogLevel)
	default:
		return fmt.Errorf("unknown configuration key: %s", *key)
	}
	
	return nil
}

// setConfig 設定値を設定
func (cc *ConfigCommand) setConfig(args []string) error {
	fs := flag.NewFlagSet("config set", flag.ContinueOnError)
	configFile := fs.String("file", cc.getConfigFilePath(), "Configuration file path")
	key := fs.String("key", "", "Configuration key to set (required)")
	value := fs.String("value", "", "Configuration value to set (required)")
	
	if err := fs.Parse(args); err != nil {
		return err
	}
	
	if *key == "" || *value == "" {
		fmt.Println("Error: Both -key and -value flags are required")
		fs.Usage()
		return fmt.Errorf("missing key or value")
	}
	
	config, err := cc.loadConfigFile(*configFile)
	if err != nil {
		// ファイルが存在しない場合は新規作成
		config = &Config{
			Host:     "localhost",
			Port:     "8080",
			Debug:    false,
			LogLevel: "info",
		}
	}
	
	// 4. 設定の検証機能
	switch *key {
	case "host":
		config.Host = *value
	case "port":
		config.Port = *value
	case "debug":
		if *value == "true" || *value == "1" {
			config.Debug = true
		} else if *value == "false" || *value == "0" {
			config.Debug = false
		} else {
			return fmt.Errorf("invalid boolean value for debug: %s", *value)
		}
	case "log_level":
		validLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
		if !validLevels[*value] {
			return fmt.Errorf("invalid log level: %s (valid: debug, info, warn, error)", *value)
		}
		config.LogLevel = *value
	default:
		return fmt.Errorf("unknown configuration key: %s", *key)
	}
	
	if err := cc.saveConfigFile(*configFile, config); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}
	
	fmt.Printf("Configuration updated: %s = %s\n", *key, *value)
	return nil
}

// initConfig 設定ファイルの初期化
func (cc *ConfigCommand) initConfig(args []string) error {
	fs := flag.NewFlagSet("config init", flag.ContinueOnError)
	configFile := fs.String("file", cc.getConfigFilePath(), "Configuration file path")
	force := fs.Bool("force", false, "Force overwrite existing config file")
	
	if err := fs.Parse(args); err != nil {
		return err
	}
	
	// 既存ファイルの確認
	if _, err := os.Stat(*configFile); err == nil && !*force {
		return fmt.Errorf("config file already exists: %s (use -force to overwrite)", *configFile)
	}
	
	// デフォルト設定を作成
	defaultConfig := &Config{
		Host:     "localhost",
		Port:     "8080",
		Debug:    false,
		LogLevel: "info",
	}
	
	if err := cc.saveConfigFile(*configFile, defaultConfig); err != nil {
		return fmt.Errorf("failed to create config file: %v", err)
	}
	
	fmt.Printf("Configuration file initialized: %s\n", *configFile)
	return nil
}

// loadConfigFile 設定ファイルの読み込み
func (cc *ConfigCommand) loadConfigFile(filename string) (*Config, error) {
	// 3. JSON/YAML形式のサポート（JSONのみ実装）
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to parse JSON config: %v", err)
	}
	
	return &config, nil
}

// saveConfigFile 設定ファイルの保存
func (cc *ConfigCommand) saveConfigFile(filename string, config *Config) error {
	// ディレクトリの作成
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode JSON config: %v", err)
	}
	
	return nil
}