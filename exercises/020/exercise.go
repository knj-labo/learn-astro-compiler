package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Exercise 020: WebSocketリアルタイム通信

このエクササイズでは、WebSocketを使用したリアルタイム通信システムの構築を学びます：

1. WebSocketサーバーの実装
   - 接続の管理
   - メッセージの送受信
   - クライアントの管理

2. リアルタイム機能
   - チャットシステム
   - ライブ通知
   - リアルタイムデータ更新

3. 高度な機能
   - ルーム/チャンネル管理
   - 認証とセッション
   - メッセージの永続化

期待される動作:
- 高性能なリアルタイム通信
- 複数クライアントの同時接続
- 信頼性のあるメッセージ配信
*/

func main() {
	fmt.Println("Exercise 020: WebSocket Real-time Communication")
	
	// WebSocketサーバーを初期化
	server := NewWebSocketServer()
	
	// ルートを設定
	server.SetupRoutes()
	
	fmt.Println("WebSocket server starting on :8080")
	fmt.Println("Open http://localhost:8080 in your browser")
	
	// サーバーを開始
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// WebSocketServer構造体
type WebSocketServer struct {
	clients    map[*Client]bool
	rooms      map[string]*Room
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

// Client構造体
type Client struct {
	ID       string
	Username string
	Room     *Room
	conn     interface{} // WebSocket connection (gorilla/websocket.Conn)
	send     chan []byte
}

// Room構造体
type Room struct {
	ID      string
	Name    string
	clients map[*Client]bool
}

// Message構造体
type Message struct {
	Type      string      `json:"type"`
	Content   string      `json:"content"`
	Username  string      `json:"username"`
	Room      string      `json:"room"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
}

// NewWebSocketServer関数の実装
func NewWebSocketServer() *WebSocketServer {
	// TODO: 実装する
	// ヒント:
	// 1. WebSocketServer構造体を初期化
	// 2. チャネルとマップを初期化
	// 3. バックグラウンドでハブを開始
	return nil
}

// SetupRoutes メソッドの実装
func (wss *WebSocketServer) SetupRoutes() {
	// TODO: 実装する
	// ヒント:
	// 1. WebSocketエンドポイントを設定
	// 2. 静的ファイルの配信
	// 3. REST APIエンドポイント
}

// Run メソッドの実装
func (wss *WebSocketServer) Run() {
	// TODO: 実装する
	// ヒント:
	// 1. メインループを実装
	// 2. クライアントの登録/登録解除を処理
	// 3. メッセージのブロードキャスト
	// 4. ルーム管理
}

// HandleWebSocket メソッドの実装
func (wss *WebSocketServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// TODO: 実装する
	// ヒント:
	// 1. WebSocket接続をアップグレード
	// 2. 新しいクライアントを作成
	// 3. クライアントを登録
	// 4. メッセージ処理のgoroutineを開始
}

// RegisterClient メソッドの実装
func (wss *WebSocketServer) RegisterClient(client *Client) {
	// TODO: 実装する
	// ヒント:
	// 1. クライアントをサーバーに追加
	// 2. ルームへの参加処理
	// 3. 他のクライアントに通知
}

// UnregisterClient メソッドの実装
func (wss *WebSocketServer) UnregisterClient(client *Client) {
	// TODO: 実装する
	// ヒント:
	// 1. クライアントをサーバーから削除
	// 2. ルームからの退出処理
	// 3. 接続を閉じる
}

// BroadcastMessage メソッドの実装
func (wss *WebSocketServer) BroadcastMessage(message Message) {
	// TODO: 実装する
	// ヒント:
	// 1. メッセージをJSONにエンコード
	// 2. 適切なクライアントに送信
	// 3. ルーム別の配信
}

// CreateRoom メソッドの実装
func (wss *WebSocketServer) CreateRoom(roomID, roomName string) *Room {
	// TODO: 実装する
	// ヒント:
	// 1. 新しいRoom構造体を作成
	// 2. サーバーのルームマップに追加
	// 3. 初期化処理
	return nil
}

// GetRoom メソッドの実装
func (wss *WebSocketServer) GetRoom(roomID string) *Room {
	// TODO: 実装する
	// ヒント:
	// 1. ルームマップから検索
	// 2. 存在しない場合はnilを返す
	return nil
}

// NewClient関数の実装
func NewClient(id, username string, conn interface{}) *Client {
	// TODO: 実装する
	// ヒント:
	// 1. Client構造体を初期化
	// 2. sendチャネルを作成
	// 3. 接続を設定
	return nil
}

// ReadMessages メソッドの実装
func (c *Client) ReadMessages(server *WebSocketServer) {
	// TODO: 実装する
	// ヒント:
	// 1. WebSocketからメッセージを読み取り
	// 2. メッセージをパースしてサーバーに送信
	// 3. エラーハンドリング
	// 4. 接続が閉じられた時の処理
}

// WriteMessages メソッドの実装
func (c *Client) WriteMessages() {
	// TODO: 実装する
	// ヒント:
	// 1. sendチャネルからメッセージを受信
	// 2. WebSocketにメッセージを書き込み
	// 3. タイムアウト処理
	// 4. 接続エラーの処理
}

// JoinRoom メソッドの実装
func (c *Client) JoinRoom(room *Room) {
	// TODO: 実装する
	// ヒント:
	// 1. 現在のルームから退出
	// 2. 新しいルームに参加
	// 3. 参加通知を送信
}

// LeaveRoom メソッドの実装
func (c *Client) LeaveRoom() {
	// TODO: 実装する
	// ヒント:
	// 1. 現在のルームから退出
	// 2. 退出通知を送信
	// 3. クライアントのルーム情報をクリア
}

// AddClient メソッド（Room）の実装
func (r *Room) AddClient(client *Client) {
	// TODO: 実装する
	// ヒント:
	// 1. クライアントをルームに追加
	// 2. クライアントのルーム情報を更新
}

// RemoveClient メソッド（Room）の実装
func (r *Room) RemoveClient(client *Client) {
	// TODO: 実装する
	// ヒント:
	// 1. クライアントをルームから削除
	// 2. ルームが空の場合の処理
}

// BroadcastToRoom メソッドの実装
func (r *Room) BroadcastToRoom(message Message) {
	// TODO: 実装する
	// ヒント:
	// 1. ルーム内の全クライアントにメッセージを送信
	// 2. JSONエンコード
	// 3. エラーハンドリング
}

// GetOnlineUsers メソッドの実装
func (r *Room) GetOnlineUsers() []string {
	// TODO: 実装する
	// ヒント:
	// 1. ルーム内のオンラインユーザー一覧を取得
	// 2. ユーザー名のスライスを返す
	return nil
}

// HandleChatMessage ヘルパー関数の実装
func HandleChatMessage(server *WebSocketServer, client *Client, message Message) {
	// TODO: 実装する
	// ヒント:
	// 1. チャットメッセージの処理
	// 2. メッセージの検証
	// 3. ルームへのブロードキャスト
}

// HandleJoinRoom ヘルパー関数の実装
func HandleJoinRoom(server *WebSocketServer, client *Client, roomID string) {
	// TODO: 実装する
	// ヒント:
	// 1. ルーム参加の処理
	// 2. ルームが存在しない場合の作成
	// 3. 参加通知の送信
}

// HandleLeaveRoom ヘルパー関数の実装
func HandleLeaveRoom(server *WebSocketServer, client *Client) {
	// TODO: 実装する
	// ヒント:
	// 1. ルーム退出の処理
	// 2. 退出通知の送信
	// 3. クライアント情報の更新
}