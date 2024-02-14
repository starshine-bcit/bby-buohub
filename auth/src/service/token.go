package service

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/starshine-bcit/bby-buohub/auth/util"
)

type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ParsingError struct {
	Token *jwt.Token
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("Could not parse token. Token: %+v\n", e.Token)
}

func NewAccessToken(claims *UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodES384, claims)
	return accessToken.SignedString(util.PrivKey)
}

func NewRefreshToken(claims *UserClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES384, claims)
	return refreshToken.SignedString(util.PrivKey)
}

func ParseAccessToken(accessToken string) (*jwt.Token, *UserClaims, error) {
	claims := new(UserClaims)
	parsedAccessToken, err := jwt.ParseWithClaims(
		accessToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return util.PubKey, nil
		},
	)
	if err != nil {
		util.WarningLogger.Printf("Error when validating access token. err: %v\n", err.Error())
		return nil, nil, &ParsingError{parsedAccessToken}
	}
	// if !parsedAccessToken.Valid {
	// 	util.WarningLogger.Println("Parsed access token is not valid")
	// 	return nil, &TokenError{claims}
	// }
	return parsedAccessToken, claims, nil
}

func ParseRefreshToken(refreshToken string) (*jwt.Token, *UserClaims, error) {
	claims := new(UserClaims)
	parsedRefreshToken, err := jwt.ParseWithClaims(
		refreshToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return util.PubKey, nil
		},
	)
	if err != nil {
		util.WarningLogger.Printf("Error when validating refresh token. err: %v\n", err.Error())
		return nil, nil, &ParsingError{parsedRefreshToken}
	}
	// if !parsedRefreshToken.Valid {
	// 	util.WarningLogger.Println("Parsed access token is not valid")
	// 	return nil, &TokenError{claims}
	// }
	return parsedRefreshToken, claims, nil
}
