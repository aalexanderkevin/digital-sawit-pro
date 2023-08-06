package handler

import (
	"digital-sawit-pro/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/segmentio/ksuid"
)

func generateJwt(user model.User) (*string, error) {
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
	var secretKey = []byte("digital-sawit-pro")
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
