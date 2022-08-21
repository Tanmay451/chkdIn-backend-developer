package utility

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword create hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verify password
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateSession setting cookies as per jwt
func CreateSession(c *gin.Context, email string) string {
	var SessionID string
	maxAge := 365 * 24 * 60 * 60 * 5 //seconds

	domain := os.Getenv("HOST")
	SessionID, _ = CreateJWT(email)
	c.SetCookie("tokenString", SessionID, maxAge, "/", domain, false, true)
	return SessionID
}

// GetTimeForCookies to get time of 30 min
func GetTimeForCookies() time.Time {
	return time.Now().Add((365 * 5) * 24 * time.Hour)
}

// JWTAuthClaims Structure to store JWT
type JWTAuthClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// CreateJWT will create JWT string and send
func CreateJWT(email string) (string, error) {
	expirationTime := GetTimeForCookies()

	claims := JWTAuthClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println("token :	", token)

	var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Println("tokenString :	", err)
	}
	log.Println("tokenString :	", tokenString)

	return tokenString, err
}
