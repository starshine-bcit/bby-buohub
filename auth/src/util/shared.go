package util

type TokenBody struct {
    AccessToken string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

type AuthBody struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type AuthResponse struct {
    Ok bool `json:"ok"`
    Refreshed bool `json:"refreshed"`
    NewToken string `json:"newToken"`
}

type CreateResponse struct {
    Created bool `json:"created"`
    AccessToken string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
    Valid bool `json:"valid"`
    AccessToken string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
}

type ErrorResponse struct {
    ErrorName string `json:"errorName"`
    ErrorText string `json:"errorText"`
}