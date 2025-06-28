package main

import (
	"strings"
	"testing"
)

func TestSecurityManager(t *testing.T) {
	sm := NewSecurityManager()
	if sm == nil {
		t.Fatal("NewSecurityManager returned nil")
	}
}

func TestEncryptDecrypt(t *testing.T) {
	sm := NewSecurityManager()
	
	plaintext := "This is a secret message"
	
	// Test encryption
	encrypted, err := sm.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}
	
	if encrypted == "" {
		t.Error("Encrypted text should not be empty")
	}
	
	if encrypted == plaintext {
		t.Error("Encrypted text should be different from plaintext")
	}
	
	// Test decryption
	decrypted, err := sm.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}
	
	if decrypted != plaintext {
		t.Errorf("Decrypted text %s does not match original %s", decrypted, plaintext)
	}
}

func TestPasswordHashing(t *testing.T) {
	sm := NewSecurityManager()
	
	password := "MySecurePassword123!"
	
	// Test password hashing
	hashedPassword, err := sm.HashPassword(password)
	if err != nil {
		t.Fatalf("Password hashing failed: %v", err)
	}
	
	if hashedPassword == "" {
		t.Error("Hashed password should not be empty")
	}
	
	if hashedPassword == password {
		t.Error("Hashed password should be different from original")
	}
	
	// Test password verification with correct password
	isValid := sm.VerifyPassword(password, hashedPassword)
	if !isValid {
		t.Error("Password verification should succeed with correct password")
	}
	
	// Test password verification with wrong password
	isInvalid := sm.VerifyPassword("WrongPassword", hashedPassword)
	if isInvalid {
		t.Error("Password verification should fail with wrong password")
	}
}

func TestJWTGeneration(t *testing.T) {
	sm := NewSecurityManager()
	
	claims := map[string]interface{}{
		"user_id":  123,
		"username": "testuser",
		"role":     "admin",
	}
	
	// Test JWT generation
	token, err := sm.GenerateJWT(claims)
	if err != nil {
		t.Fatalf("JWT generation failed: %v", err)
	}
	
	if token == "" {
		t.Error("JWT token should not be empty")
	}
	
	// JWT should have 3 parts separated by dots
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Errorf("JWT should have 3 parts, got %d", len(parts))
	}
}

func TestJWTVerification(t *testing.T) {
	sm := NewSecurityManager()
	
	originalClaims := map[string]interface{}{
		"user_id":  123,
		"username": "testuser",
		"role":     "admin",
	}
	
	// Generate token
	token, err := sm.GenerateJWT(originalClaims)
	if err != nil {
		t.Fatalf("JWT generation failed: %v", err)
	}
	
	// Verify token
	decodedClaims, err := sm.VerifyJWT(token)
	if err != nil {
		t.Fatalf("JWT verification failed: %v", err)
	}
	
	if decodedClaims == nil {
		t.Error("Decoded claims should not be nil")
	}
	
	// Check specific claims (if implementation preserves them)
	if len(decodedClaims) == 0 {
		t.Error("Decoded claims should not be empty")
	}
}

func TestInvalidJWTVerification(t *testing.T) {
	sm := NewSecurityManager()
	
	// Test with invalid token
	_, err := sm.VerifyJWT("invalid.token.here")
	if err == nil {
		t.Error("Should return error for invalid JWT")
	}
	
	// Test with empty token
	_, err = sm.VerifyJWT("")
	if err == nil {
		t.Error("Should return error for empty JWT")
	}
}

func TestFileEncryption(t *testing.T) {
	sm := NewSecurityManager()
	
	content := "This is secret file content"
	inputFile := "test_input.txt"
	encryptedFile := "test_encrypted.enc"
	
	// Test file encryption
	err := sm.EncryptFile(inputFile, encryptedFile, content)
	if err != nil {
		t.Fatalf("File encryption failed: %v", err)
	}
	
	// Test file decryption
	decryptedContent, err := sm.DecryptFile(encryptedFile)
	if err != nil {
		t.Fatalf("File decryption failed: %v", err)
	}
	
	if decryptedContent != content {
		t.Errorf("Decrypted content %s does not match original %s", decryptedContent, content)
	}
}

func TestMessageSigning(t *testing.T) {
	sm := NewSecurityManager()
	
	message := "Important message to be signed"
	
	// Test message signing
	signature, err := sm.SignMessage(message)
	if err != nil {
		t.Fatalf("Message signing failed: %v", err)
	}
	
	if signature == "" {
		t.Error("Signature should not be empty")
	}
	
	// Test signature verification
	isValid := sm.VerifySignature(message, signature)
	if !isValid {
		t.Error("Signature verification should succeed for valid signature")
	}
	
	// Test signature verification with modified message
	modifiedMessage := message + " tampered"
	isInvalid := sm.VerifySignature(modifiedMessage, signature)
	if isInvalid {
		t.Error("Signature verification should fail for modified message")
	}
}

func TestSecureHash(t *testing.T) {
	sm := NewSecurityManager()
	
	data := "Data to be hashed"
	
	// Test secure hashing
	hash := sm.SecureHash(data)
	if hash == "" {
		t.Error("Hash should not be empty")
	}
	
	if hash == data {
		t.Error("Hash should be different from original data")
	}
	
	// Test hash consistency
	hash2 := sm.SecureHash(data)
	if hash != hash2 {
		t.Error("Same input should produce same hash")
	}
	
	// Test different data produces different hash
	differentData := data + " modified"
	differentHash := sm.SecureHash(differentData)
	if hash == differentHash {
		t.Error("Different input should produce different hash")
	}
}

func TestHashLength(t *testing.T) {
	sm := NewSecurityManager()
	
	// SHA-256 hash should be 64 characters (32 bytes in hex)
	hash := sm.SecureHash("test")
	if len(hash) != 64 {
		t.Errorf("SHA-256 hash should be 64 characters, got %d", len(hash))
	}
}

func TestGenerateRandomKey(t *testing.T) {
	// Test different key lengths
	lengths := []int{16, 32, 64}
	
	for _, length := range lengths {
		key, err := GenerateRandomKey(length)
		if err != nil {
			t.Fatalf("GenerateRandomKey(%d) failed: %v", length, err)
		}
		
		if len(key) != length {
			t.Errorf("Expected key length %d, got %d", length, len(key))
		}
		
		// Generate another key to test uniqueness
		key2, err := GenerateRandomKey(length)
		if err != nil {
			t.Fatalf("GenerateRandomKey(%d) failed: %v", length, err)
		}
		
		// Keys should be different (very unlikely to be same with good randomness)
		if string(key) == string(key2) {
			t.Error("Generated keys should be different")
		}
	}
}

func TestZeroLengthKey(t *testing.T) {
	_, err := GenerateRandomKey(0)
	if err != nil {
		// It's acceptable to return an error for zero length
		return
	}
	
	// If no error, should return empty slice
	key, _ := GenerateRandomKey(0)
	if len(key) != 0 {
		t.Error("Zero length key should return empty slice")
	}
}

func TestEncryptionWithDifferentKeys(t *testing.T) {
	sm1 := NewSecurityManager()
	sm2 := NewSecurityManager()
	
	plaintext := "Test message"
	
	// Encrypt with first security manager
	encrypted1, err := sm1.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encryption with sm1 failed: %v", err)
	}
	
	// Try to decrypt with second security manager (should fail)
	_, err = sm2.Decrypt(encrypted1)
	if err == nil {
		t.Error("Decryption with different key should fail")
	}
}

func TestLargeDataEncryption(t *testing.T) {
	sm := NewSecurityManager()
	
	// Test with larger data
	largeData := strings.Repeat("This is a large text that we want to encrypt. ", 100)
	
	encrypted, err := sm.Encrypt(largeData)
	if err != nil {
		t.Fatalf("Large data encryption failed: %v", err)
	}
	
	decrypted, err := sm.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Large data decryption failed: %v", err)
	}
	
	if decrypted != largeData {
		t.Error("Large data encryption/decryption failed")
	}
}

func TestEmptyDataEncryption(t *testing.T) {
	sm := NewSecurityManager()
	
	// Test with empty string
	encrypted, err := sm.Encrypt("")
	if err != nil {
		t.Fatalf("Empty data encryption failed: %v", err)
	}
	
	decrypted, err := sm.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Empty data decryption failed: %v", err)
	}
	
	if decrypted != "" {
		t.Error("Empty data encryption/decryption failed")
	}
}