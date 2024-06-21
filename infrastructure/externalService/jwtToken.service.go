package externalService

import (
	"NetFarm/application/interfaces/iinfrastructure/iexternalServices"
	"NetFarm/shared/models/common"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtTokenService struct {
}

var jwtKey = []byte("chave_secreta_do_jwt")

func NewJwtTokenService() iexternalServices.IJwtTokenService {
	return &JwtTokenService{}
}

func (s *JwtTokenService) GenerateToken(email string, userId string, claim []string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &common.Claims{
		Email:  email,
		UserID: userId,
		Claims: claim,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JwtTokenService) ValidateToken(tokenString string) (*common.Claims, error) {
	claims := &common.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
