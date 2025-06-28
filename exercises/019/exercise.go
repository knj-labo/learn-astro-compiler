package main

import (
	"fmt"
	"log"
)

/*
Exercise 019: 暗号化とセキュリティ

このエクササイズでは、Goの暗号化パッケージを使用してセキュアなアプリケーションを構築することを学びます：

1. 暗号化とハッシュ化
   - AES暗号化/復号化
   - ハッシュ関数の使用
   - デジタル署名

2. セキュアな通信
   - TLS設定
   - 証明書の管理
   - セキュアなHTTPクライアント

3. 認証とセッション管理
   - JWTトークン
   - パスワードハッシュ化
   - セッション管理

期待される動作:
- 強固な暗号化機能
- セキュアな通信
- 適切な認証システム
*/

func main() {
	fmt.Println("Exercise 019: Encryption and Security")
	
	// セキュリティマネージャーを初期化
	security := NewSecurityManager()
	
	// データ暗号化のデモ
	fmt.Println("=== Data Encryption Demo ===")
	plaintext := "This is sensitive information that needs to be encrypted."
	
	encrypted, err := security.Encrypt(plaintext)
	if err != nil {
		log.Printf("Encryption error: %v", err)
	} else {
		fmt.Printf("Encrypted: %s\n", encrypted)
		
		decrypted, err := security.Decrypt(encrypted)
		if err != nil {
			log.Printf("Decryption error: %v", err)
		} else {
			fmt.Printf("Decrypted: %s\n", decrypted)
		}
	}
	
	// パスワードハッシュ化のデモ
	fmt.Println("\n=== Password Hashing Demo ===")
	password := "MySecurePassword123!"
	
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		log.Printf("Password hashing error: %v", err)
	} else {
		fmt.Printf("Hashed password: %s\n", hashedPassword)
		
		// パスワード検証
		isValid := security.VerifyPassword(password, hashedPassword)
		fmt.Printf("Password verification: %t\n", isValid)
		
		// 間違ったパスワードでの検証
		isInvalid := security.VerifyPassword("WrongPassword", hashedPassword)
		fmt.Printf("Wrong password verification: %t\n", isInvalid)
	}
	
	// JWTトークンのデモ
	fmt.Println("\n=== JWT Token Demo ===")
	claims := map[string]interface{}{
		"user_id": 123,
		"username": "alice",
		"role": "admin",
	}
	
	token, err := security.GenerateJWT(claims)
	if err != nil {
		log.Printf("JWT generation error: %v", err)
	} else {
		fmt.Printf("Generated JWT: %s\n", token[:50]+"...")
		
		// トークン検証
		decodedClaims, err := security.VerifyJWT(token)
		if err != nil {
			log.Printf("JWT verification error: %v", err)
		} else {
			fmt.Printf("Decoded claims: %+v\n", decodedClaims)
		}
	}
	
	// ファイル暗号化のデモ
	fmt.Println("\n=== File Encryption Demo ===")
	content := "Secret file content that should be protected."
	filename := "secret.txt"
	encryptedFilename := "secret.txt.enc"
	
	err = security.EncryptFile(filename, encryptedFilename, content)
	if err != nil {
		log.Printf("File encryption error: %v", err)
	} else {
		fmt.Printf("File encrypted: %s\n", encryptedFilename)
		
		decryptedContent, err := security.DecryptFile(encryptedFilename)
		if err != nil {
			log.Printf("File decryption error: %v", err)
		} else {
			fmt.Printf("Decrypted content: %s\n", decryptedContent)
		}
	}
	
	// デジタル署名のデモ
	fmt.Println("\n=== Digital Signature Demo ===")
	message := "Important message that needs to be signed."
	
	signature, err := security.SignMessage(message)
	if err != nil {
		log.Printf("Signing error: %v", err)
	} else {
		fmt.Printf("Message signature: %s\n", signature[:50]+"...")
		
		// 署名検証
		isValid := security.VerifySignature(message, signature)
		fmt.Printf("Signature verification: %t\n", isValid)
	}
	
	// セキュアハッシュのデモ
	fmt.Println("\n=== Secure Hash Demo ===")
	data := "Data to be hashed for integrity verification."
	
	hash := security.SecureHash(data)
	fmt.Printf("SHA-256 hash: %s\n", hash)
	
	// 同じデータの再ハッシュ
	hash2 := security.SecureHash(data)
	fmt.Printf("Hashes match: %t\n", hash == hash2)
	
	// 異なるデータのハッシュ
	hash3 := security.SecureHash(data + " modified")
	fmt.Printf("Different data hashes match: %t\n", hash == hash3)
	
	fmt.Println("\n=== Security Demo Complete ===")
}

// SecurityManager構造体
type SecurityManager struct {
	encryptionKey []byte
	jwtSecret     []byte
}

// NewSecurityManager関数の実装
func NewSecurityManager() *SecurityManager {
	// TODO: 実装する
	// ヒント:
	// 1. SecurityManager構造体を初期化
	// 2. 暗号化キーとJWTシークレットを生成
	// 3. セキュアな乱数生成を使用
	return nil
}

// Encrypt メソッドの実装
func (sm *SecurityManager) Encrypt(plaintext string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. AES暗号化を使用
	// 2. crypto/aesパッケージを使用
	// 3. 初期化ベクトル(IV)を生成
	// 4. base64エンコードして返す
	return "", nil
}

// Decrypt メソッドの実装
func (sm *SecurityManager) Decrypt(ciphertext string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. base64デコード
	// 2. IVを抽出
	// 3. AES復号化
	// 4. プレーンテキストを返す
	return "", nil
}

// HashPassword メソッドの実装
func (sm *SecurityManager) HashPassword(password string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. bcryptパッケージを使用
	// 2. 適切なコストを設定
	// 3. ソルト付きハッシュを生成
	return "", nil
}

// VerifyPassword メソッドの実装
func (sm *SecurityManager) VerifyPassword(password, hashedPassword string) bool {
	// TODO: 実装する
	// ヒント:
	// 1. bcrypt.CompareHashAndPasswordを使用
	// 2. エラーハンドリングを実装
	return false
}

// GenerateJWT メソッドの実装
func (sm *SecurityManager) GenerateJWT(claims map[string]interface{}) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. github.com/golang-jwt/jwtパッケージを使用
	// 2. 有効期限を設定
	// 3. HMACで署名
	return "", nil
}

// VerifyJWT メソッドの実装
func (sm *SecurityManager) VerifyJWT(tokenString string) (map[string]interface{}, error) {
	// TODO: 実装する
	// ヒント:
	// 1. トークンをパース
	// 2. 署名を検証
	// 3. クレームを返す
	return nil, nil
}

// EncryptFile メソッドの実装
func (sm *SecurityManager) EncryptFile(inputFile, outputFile, content string) error {
	// TODO: 実装する
	// ヒント:
	// 1. コンテンツを暗号化
	// 2. 暗号化されたデータをファイルに保存
	// 3. エラーハンドリングを実装
	return nil
}

// DecryptFile メソッドの実装
func (sm *SecurityManager) DecryptFile(filename string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. ファイルから暗号化データを読み取り
	// 2. データを復号化
	// 3. プレーンテキストを返す
	return "", nil
}

// SignMessage メソッドの実装
func (sm *SecurityManager) SignMessage(message string) (string, error) {
	// TODO: 実装する
	// ヒント:
	// 1. crypto/rsaパッケージを使用
	// 2. RSA秘密鍵で署名
	// 3. base64エンコードして返す
	return "", nil
}

// VerifySignature メソッドの実装
func (sm *SecurityManager) VerifySignature(message, signature string) bool {
	// TODO: 実装する
	// ヒント:
	// 1. base64デコード
	// 2. RSA公開鍵で検証
	// 3. 署名の有効性を返す
	return false
}

// SecureHash メソッドの実装
func (sm *SecurityManager) SecureHash(data string) string {
	// TODO: 実装する
	// ヒント:
	// 1. crypto/sha256パッケージを使用
	// 2. ハッシュを計算
	// 3. 16進文字列で返す
	return ""
}

// GenerateRandomKey ヘルパー関数の実装
func GenerateRandomKey(length int) ([]byte, error) {
	// TODO: 実装する
	// ヒント:
	// 1. crypto/randパッケージを使用
	// 2. セキュアな乱数を生成
	// 3. 指定された長さのキーを作成
	return nil, nil
}