package util

import (
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// const secretKey string = "acb52590-e620-4e0d-8686-ea613346afca"
var secretKey string = uuid.New().String()

const issuer string = "sakurafisch"

type authCustomClaims struct {
	UserId uint   `json:"userId"`
	Eamil  string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(userId uint, email string) (string, error) {
	claims := &authCustomClaims{
		userId,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// encoding string
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return t, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, myKeyFunc)
}

func myKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
		return nil, errors.New("Invalid token")
	}

	return []byte(secretKey), nil
}
