package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " <password>")
	fmt.Println("Exmaple: " + os.Args[0] + " VerySafePassword123")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

// Store in env
var secretKey = "87de5bcb-a9f3-4d44-9f86-5971b452050c"

func salt() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func hashPassword(plainText string, salt string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, plainText+salt)
	hashedValue := hash.Sum(nil)
	return hex.EncodeToString(hashedValue)
}

func main() {
	password := checkArgs()
	salt := salt() // vault this "HCEdS15fWdy6gdkep9GSnsXAhXw0g5ewct0pVuEVCvY="
	hashedPassword := hashPassword(password, salt)
	fmt.Println("Password: " + password)
	fmt.Println("Salt: " + salt)
	fmt.Println("Hashed Password: " + hashedPassword)
}
