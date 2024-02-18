package service

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/starshine-bcit/bby-buohub/auth/util"
	"gorm.io/gorm"
)

var Usernames = []string{}

type TokenError struct {
	Claims *UserClaims
}

type HeaderError struct {
	Header map[string][]string
}

func (e *TokenError) Error() string {
	return fmt.Sprintf("Parsed token not valid. Claims: %+v\n", e.Claims)
}

func (e *HeaderError) Error() string {
	return fmt.Sprintf("Header not valid. Header: %+v", e.Header)
}

func GenUserAccessClaims(username string) *UserClaims {
	claims := &UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate((time.Now().Add(15 * time.Minute))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	return claims
}

func GenUserRefreshClaims(username string) *UserClaims {
	claims := &UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate((time.Now().Add(24 * time.Hour))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	return claims
}

func ValidateUserClaims(token *jwt.Token, claims *UserClaims, db *gorm.DB) error {
	if !token.Valid {
		return &TokenError{claims}
	}
	username := claims.Username
	if !slices.Contains(Usernames, username) {
		if exists := CheckUsername(db, username); exists {
			Usernames = append(Usernames, username)
		} else {
			return errors.New("claimed username does not exist")
		}
	}
	return nil
}

func ValidateAuthRequest(b *util.TokenBody, db *gorm.DB) (*util.TokenBody, error) {
	accessToken, accessClaims, err := ParseAccessToken(b.AccessToken)
	if err != nil {
		refreshToken, refreshClaims, err := ParseRefreshToken(b.RefreshToken)
		if err != nil {
			return nil, err
		}
		err = ValidateUserClaims(refreshToken, refreshClaims, db)
		if err != nil {
			return nil, err
		} else {
			newClaims := GenUserAccessClaims(accessClaims.Username)
			newToken, err := NewAccessToken(newClaims)
			if err != nil {
				return nil, err
			} else {
				b.AccessToken = newToken
				return b, nil
			}
		}
	} else {
		err = ValidateUserClaims(accessToken, accessClaims, db)
		if err != nil {
			return b, nil
		}
	}
	return nil, nil
}
