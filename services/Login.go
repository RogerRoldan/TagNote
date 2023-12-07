package services

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/roger/workhub/models"
	"gorm.io/gorm"
	"time"
)

func Login(db *gorm.DB, credentialsUser models.Login) (models.UserDTO, bool) {
	var user models.User
	response := db.Where("email = ?", credentialsUser.Email).First(&user)

	if response.Error != nil {
		return models.UserDTO{}, false
	}

	if user.Password != Hash256(credentialsUser.Password) {
		return models.UserDTO{}, false
	}

	userDTO := models.MapUserToDTO(user)
	token, error := GenerateToken(userDTO, "secret")
	if error != nil {
		return models.UserDTO{}, false
	}
	userDTO.Token = token

	return userDTO, true
}

func Hash256(password string) string {
	// Convert the password string to a byte slice
	passwordBytes := []byte(password)

	// Create a new SHA-256 hash
	hasher := sha256.New()

	// Write the password bytes to the hash
	hasher.Write(passwordBytes)

	// Get the final hash sum
	hashSum := hasher.Sum(nil)

	// Convert the hash sum to a hexadecimal string
	hashString := hex.EncodeToString(hashSum)

	return hashString
}

func GenerateToken(dto models.UserDTO, secretKey string) (string, error) {
	// Create the claims for the JWT
	claims := jwt.MapClaims{
		"id":       dto.ID,
		"username": dto.Username,
		"email":    dto.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours from now)
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
