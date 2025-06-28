package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Exercise 017: ミドルウェアとHTTPハンドラー

このエクササイズでは、HTTPミドルウェアパターンとカスタムハンドラーの実装を学びます：

1. ミドルウェアの実装
   - リクエスト/レスポンスの前後処理
   - 認証・認可機能
   - ログとメトリクス

2. ハンドラーチェーン
   - 複数のミドルウェアの組み合わせ
   - エラーハンドリング
   - レスポンスの変換

3. 実用的な機能
   - CORS対応
   - レート制限
   - キャッシュ制御

期待される動作:
- モジュラーなWebアプリケーション
- 再利用可能なミドルウェア
- 高性能なHTTPサーバー
*/

func main() {
	fmt.Println("Exercise 017: Middleware and HTTP Handlers")
	
	// HTTPサーバーを初期化
	server := NewHTTPServer()
	
	// ルートとミドルウェアを設定
	server.SetupRoutes()
	
	// サーバーを開始
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", server.Router()))
}

// HTTPServer構造体
type HTTPServer struct {
	router *Router
}

// Router構造体
type Router struct {
	middlewares []Middleware
	routes      map[string]http.HandlerFunc
}

// Middleware インターフェース
type Middleware interface {
	Handle(next http.HandlerFunc) http.HandlerFunc
}

// NewHTTPServer関数の実装
func NewHTTPServer() *HTTPServer {
	// TODO: 実装する
	return nil
}

// SetupRoutes メソッドの実装
func (hs *HTTPServer) SetupRoutes() {
	// TODO: 実装する
	// ヒント:
	// 1. 各種ミドルウェアを追加
	// 2. APIエンドポイントを設定
	// 3. 静的ファイルのハンドリング
}

// Router メソッドの実装
func (hs *HTTPServer) Router() http.Handler {
	// TODO: 実装する
	return nil
}

// LoggingMiddleware構造体とメソッドの実装
type LoggingMiddleware struct{}

func (lm *LoggingMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	// TODO: 実装する
	return nil
}

// AuthMiddleware構造体とメソッドの実装
type AuthMiddleware struct{}

func (am *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	// TODO: 実装する
	return nil
}

// CORSMiddleware構造体とメソッドの実装
type CORSMiddleware struct{}

func (cm *CORSMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	// TODO: 実装する
	return nil
}