package jwttoken

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

type TokenDetail struct {
	AccessToken          string
	ExpirationTimeInUnix int64
}

type AccessDetail struct {
	UserID     int64
	Authorized bool
}

func CreateToken(userId int64) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.ExpirationTimeInUnix = time.Now().Add(time.Minute * 15).Unix()

	var err error
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["expiration_time"] = td.ExpirationTimeInUnix

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	td.AccessToken, err = token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func extractTokenFromRequestHeader(h http.Header) string {
	token := h.Get("Authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func getTokenFromRequest(r *http.Request) (*jwt.Token, error) {
	tokenString := extractTokenFromRequestHeader(r.Header)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong signature method")
		}

		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CheckTokenIsValid(r *http.Request) error {
	token, err := getTokenFromRequest(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetail, error) {
	token, err := getTokenFromRequest(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authorized, ok := claims["authorized"].(bool)
		if !ok {
			return nil, err
		}

		userId, ok := claims["user_id"].(int64)
		if !ok {
			return nil, err
		}

		return &AccessDetail{
			Authorized: authorized,
			UserID:     userId,
		}, nil
	}

	return nil, err
}
