package main

import (
	"testing"
	"time"
)

func TestWebSocketServer(t *testing.T) {
	server := NewWebSocketServer()
	if server == nil {
		t.Fatal("NewWebSocketServer returned nil")
	}
}

func TestSetupRoutes(t *testing.T) {
	server := NewWebSocketServer()
	
	// Should not panic
	server.SetupRoutes()
}

func TestCreateRoom(t *testing.T) {
	server := NewWebSocketServer()
	
	roomID := "test-room"
	roomName := "Test Room"
	
	room := server.CreateRoom(roomID, roomName)
	if room == nil {
		t.Error("CreateRoom should return a room")
	}
	
	if room.ID != roomID {
		t.Errorf("Expected room ID %s, got %s", roomID, room.ID)
	}
	
	if room.Name != roomName {
		t.Errorf("Expected room name %s, got %s", roomName, room.Name)
	}
}

func TestGetRoom(t *testing.T) {
	server := NewWebSocketServer()
	
	// Test getting non-existent room
	room := server.GetRoom("non-existent")
	if room != nil {
		t.Error("GetRoom should return nil for non-existent room")
	}
	
	// Create a room and test getting it
	roomID := "test-room"
	createdRoom := server.CreateRoom(roomID, "Test Room")
	if createdRoom == nil {
		t.Fatal("Failed to create room")
	}
	
	retrievedRoom := server.GetRoom(roomID)
	if retrievedRoom == nil {
		t.Error("GetRoom should return existing room")
	}
	
	if retrievedRoom.ID != roomID {
		t.Errorf("Expected room ID %s, got %s", roomID, retrievedRoom.ID)
	}
}

func TestNewClient(t *testing.T) {
	clientID := "test-client"
	username := "testuser"
	
	client := NewClient(clientID, username, nil)
	if client == nil {
		t.Fatal("NewClient returned nil")
	}
	
	if client.ID != clientID {
		t.Errorf("Expected client ID %s, got %s", clientID, client.ID)
	}
	
	if client.Username != username {
		t.Errorf("Expected username %s, got %s", username, client.Username)
	}
}

func TestRoomAddClient(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("test-room", "Test Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	// Test adding client to room
	room.AddClient(client)
	
	// Verify client was added
	if client.Room != room {
		t.Error("Client should be assigned to room")
	}
}

func TestRoomRemoveClient(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("test-room", "Test Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	// Add client to room
	room.AddClient(client)
	
	// Remove client from room
	room.RemoveClient(client)
	
	// Verify client was removed
	if client.Room == room {
		t.Error("Client should not be assigned to room after removal")
	}
}

func TestClientJoinRoom(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("test-room", "Test Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	// Test client joining room
	client.JoinRoom(room)
	
	// Verify client joined room
	if client.Room != room {
		t.Error("Client should be in the room after joining")
	}
}

func TestClientLeaveRoom(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("test-room", "Test Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	// Join room first
	client.JoinRoom(room)
	
	// Leave room
	client.LeaveRoom()
	
	// Verify client left room
	if client.Room != nil {
		t.Error("Client should not be in any room after leaving")
	}
}

func TestRoomGetOnlineUsers(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("test-room", "Test Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	// Initially no users
	users := room.GetOnlineUsers()
	if len(users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(users))
	}
	
	// Add some clients
	client1 := NewClient("client1", "user1", nil)
	client2 := NewClient("client2", "user2", nil)
	
	if client1 == nil || client2 == nil {
		t.Fatal("Failed to create clients")
	}
	
	room.AddClient(client1)
	room.AddClient(client2)
	
	users = room.GetOnlineUsers()
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestRegisterClient(t *testing.T) {
	server := NewWebSocketServer()
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	// Test registering client
	server.RegisterClient(client)
	
	// Verification depends on implementation
	// This test mainly ensures no panic occurs
}

func TestUnregisterClient(t *testing.T) {
	server := NewWebSocketServer()
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	// Register first
	server.RegisterClient(client)
	
	// Test unregistering client
	server.UnregisterClient(client)
	
	// Verification depends on implementation
	// This test mainly ensures no panic occurs
}

func TestBroadcastMessage(t *testing.T) {
	server := NewWebSocketServer()
	
	message := Message{
		Type:      "chat",
		Content:   "Hello, World!",
		Username:  "testuser",
		Room:      "test-room",
		Timestamp: time.Now().Unix(),
	}
	
	// Test broadcasting message
	server.BroadcastMessage(message)
	
	// This test mainly ensures no panic occurs
	// Actual message delivery testing would require WebSocket connections
}

func TestRoomBroadcastToRoom(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("test-room", "Test Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	message := Message{
		Type:      "chat",
		Content:   "Room message",
		Username:  "testuser",
		Room:      room.ID,
		Timestamp: time.Now().Unix(),
	}
	
	// Test room broadcast
	room.BroadcastToRoom(message)
	
	// This test mainly ensures no panic occurs
}

func TestHandleChatMessage(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("test-room", "Test Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	client.JoinRoom(room)
	
	message := Message{
		Type:     "chat",
		Content:  "Hello from test",
		Username: client.Username,
		Room:     room.ID,
	}
	
	// Test handling chat message
	HandleChatMessage(server, client, message)
	
	// This test mainly ensures no panic occurs
}

func TestHandleJoinRoom(t *testing.T) {
	server := NewWebSocketServer()
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	roomID := "test-room"
	
	// Test handling join room
	HandleJoinRoom(server, client, roomID)
	
	// This test mainly ensures no panic occurs
	// and that room is created if it doesn't exist
}

func TestHandleLeaveRoom(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("test-room", "Test Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	client := NewClient("test-client", "testuser", nil)
	if client == nil {
		t.Fatal("Failed to create client")
	}
	
	client.JoinRoom(room)
	
	// Test handling leave room
	HandleLeaveRoom(server, client)
	
	// This test mainly ensures no panic occurs
}

func TestMessageStruct(t *testing.T) {
	message := Message{
		Type:      "test",
		Content:   "test content",
		Username:  "testuser",
		Room:      "test-room",
		Timestamp: time.Now().Unix(),
		Data:      map[string]string{"key": "value"},
	}
	
	// Test message fields
	if message.Type != "test" {
		t.Errorf("Expected type 'test', got %s", message.Type)
	}
	
	if message.Content != "test content" {
		t.Errorf("Expected content 'test content', got %s", message.Content)
	}
	
	if message.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got %s", message.Username)
	}
}

func TestMultipleClientsInRoom(t *testing.T) {
	server := NewWebSocketServer()
	room := server.CreateRoom("multi-room", "Multi User Room")
	if room == nil {
		t.Fatal("Failed to create room")
	}
	
	// Create multiple clients
	clients := make([]*Client, 5)
	for i := 0; i < 5; i++ {
		clients[i] = NewClient(
			fmt.Sprintf("client-%d", i),
			fmt.Sprintf("user%d", i),
			nil,
		)
		if clients[i] == nil {
			t.Fatalf("Failed to create client %d", i)
		}
		
		// Add to room
		room.AddClient(clients[i])
	}
	
	// Check online users count
	users := room.GetOnlineUsers()
	if len(users) != 5 {
		t.Errorf("Expected 5 users in room, got %d", len(users))
	}
	
	// Remove one client
	room.RemoveClient(clients[0])
	
	users = room.GetOnlineUsers()
	if len(users) != 4 {
		t.Errorf("Expected 4 users after removal, got %d", len(users))
	}
}