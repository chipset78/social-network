package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func createSaltedHash(password, salt string) string {
	data := []byte(password + salt)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func HashPassword(password, salt string) (string, error) {
	if salt == "" {
		return "", errors.New("password salt cannot be empty")
	}

	saltedHash := createSaltedHash(password, salt)
	hashed, err := bcrypt.GenerateFromPassword([]byte(saltedHash), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPasswordHash(password, hash, salt string) bool {
	if hash == "" || salt == "" {
		log.Println("Empty hash or salt")
		return false
	}

	saltedHash := createSaltedHash(password, salt)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedHash))
	if err != nil {
		log.Printf("Password mismatch: %v", err)
	}
	return err == nil
}

func GenerateToken(userID, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(jwtSecret))
}

func ValidateToken(tokenString, jwtSecret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"].(string), nil
	}

	return "", jwt.ErrInvalidKey
}
