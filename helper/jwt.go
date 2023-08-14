package helper

import (
	"digital-sawit-pro/model"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
)

var secretKey = "digital-sawit-pro"

func GenerateJwt(user model.User) (*string, error) {
	jwtData := &JWTData{
		Id:              *user.Id,
		PhoneNumber:     *user.PhoneNumber,
		FullName:        *user.FullName,
		SuccessfulLogin: *user.SuccessfulLogin,
	}

	stdClaims := &jwt.StandardClaims{
		Id:        ksuid.New().String(),
		ExpiresAt: time.Now().Add(300 * time.Second).Unix(),
		Subject:   *user.Id,
	}

	// generate token
	accessClaims := AccessJWTClaims{
		StandardClaims: stdClaims,
		JWTData:        jwtData,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	var byteSecret = []byte(secretKey)
	accessToken, err := token.SignedString(byteSecret)
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}

type AccessJWTClaims struct {
	*jwt.StandardClaims
	*JWTData
}

type JWTData struct {
	Id              string `json:"id,omitempty" binding:"required"`
	PhoneNumber     string `json:"phone_number,omitempty" binding:"required"`
	FullName        string `json:"full_name"`
	SuccessfulLogin int    `json:"successful_login"`
}

func GetAuthToken(c echo.Context) string {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	tokens := strings.SplitN(authHeader, " ", 2)
	if len(tokens) < 2 || strings.ToLower(tokens[0]) != "bearer" {
		return ""
	}

	return tokens[1]
}

func DecodeClaims(tokenStr string) (*JWTData, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AccessJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, validSignMethod := token.Method.(*jwt.SigningMethodHMAC); !validSignMethod {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		hmacSecret := []byte(secretKey)
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(*AccessJWTClaims); ok && token.Valid {
		return claim.JWTData, nil
	}

	return nil, errors.New("invalid token")
}
