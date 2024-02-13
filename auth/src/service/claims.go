package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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

func GenUserAccessClaims(username string) (*UserClaims) {
    claims := &UserClaims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate((time.Now().Add(15 * time.Minute))),
            IssuedAt: jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }
    return claims
}

func GenUserRefreshClaims(username string) (*UserClaims) {
    claims := &UserClaims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate((time.Now().Add(24 * time.Hour))),
            IssuedAt: jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }
    return claims
}

func ValidateUserClaims(token *jwt.Token, claims *UserClaims) (error) {
    if !token.Valid {
        return &TokenError{claims}
    }
    // check username in db/cache
    return nil
}

func ValidateRequestHeader(h map[string][]string) (map[string][]string, error) {
    accessTokens, ok := h["Accesstoken"]
    if !ok || len(accessTokens) < 1 {
        return nil, &HeaderError{h}
    }
    refreshTokens, ok := h["Refreshtoken"]
    if !ok || len(refreshTokens) < 1 {
        return nil, &HeaderError{h}
    }
    accessToken, accessClaims, err := ParseAccessToken(accessTokens[0])
    if err != nil {
        return nil, err
    }
    refreshToken, refreshClaims, err := ParseRefreshToken(refreshTokens[0])
    err = ValidateUserClaims(accessToken, accessClaims)
    if err != nil {
        err = ValidateUserClaims(refreshToken, refreshClaims)
        if err != nil {
            return nil, err
        } else {
            newClaims := GenUserAccessClaims(accessClaims.Username)
            newToken, err := NewAccessToken(*newClaims)
            if err != nil {
                return nil, err
            } else {
                h["Accesstoken"][0] = newToken
                return h, nil
            }
        }
    }
    return nil, nil
}