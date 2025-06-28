package main

import (
	"fmt"
	"log"
	"os"
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
	// TODO: 実装する
	// ヒント:
	// 1. CLIApp構造体を初期化
	// 2. 利用可能なコマンドを登録
	// 3. ヘルプコマンドを追加
	return nil
}

// Run メソッドの実装
func (app *CLIApp) Run(args []string) error {
	// TODO: 実装する
	// ヒント:
	// 1. 引数を解析してコマンドを特定
	// 2. 適切なコマンドを実行
	// 3. エラーハンドリングを実装
	// 4. ヘルプメッセージの表示
	return nil
}

// RegisterCommand メソッドの実装
func (app *CLIApp) RegisterCommand(cmd Command) {
	// TODO: 実装する
	// ヒント:
	// 1. コマンドをマップに登録
	// 2. 名前の重複チェック
}

// ShowHelp メソッドの実装
func (app *CLIApp) ShowHelp() {
	// TODO: 実装する
	// ヒント:
	// 1. 使用方法を表示
	// 2. 利用可能なコマンド一覧
	// 3. 例の表示
}

// FileProcessCommand の実装

// Name メソッドの実装
func (fpc *FileProcessCommand) Name() string {
	// TODO: 実装する
	return ""
}

// Description メソッドの実装
func (fpc *FileProcessCommand) Description() string {
	// TODO: 実装する
	return ""
}

// Run メソッドの実装
func (fpc *FileProcessCommand) Run(args []string) error {
	// TODO: 実装する
	// ヒント:
	// 1. flagパッケージでフラグを定義
	// 2. 入力ファイル、出力ファイル、処理モードのフラグ
	// 3. ファイル処理ロジックを実装
	// 4. 進捗表示機能
	return nil
}

// ServerCommand の実装

// Name メソッドの実装
func (sc *ServerCommand) Name() string {
	// TODO: 実装する
	return ""
}

// Description メソッドの実装
func (sc *ServerCommand) Description() string {
	// TODO: 実装する
	return ""
}

// Run メソッドの実装
func (sc *ServerCommand) Run(args []string) error {
	// TODO: 実装する
	// ヒント:
	// 1. ポート番号、ホスト、設定ファイルのフラグ
	// 2. HTTPサーバーの起動
	// 3. グレースフルシャットダウン
	// 4. ログ出力の制御
	return nil
}

// ConfigCommand の実装

// Name メソッドの実装
func (cc *ConfigCommand) Name() string {
	// TODO: 実装する
	return ""
}

// Description メソッドの実装
func (cc *ConfigCommand) Description() string {
	// TODO: 実装する
	return ""
}

// Run メソッドの実装
func (cc *ConfigCommand) Run(args []string) error {
	// TODO: 実装する
	// ヒント:
	// 1. 設定の表示、設定、取得のサブコマンド
	// 2. 設定ファイルの管理
	// 3. JSON/YAML形式のサポート
	// 4. 設定の検証機能
	return nil
}