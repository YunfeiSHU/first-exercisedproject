package model

import (
	"github.com/golang-jwt/jwt/v5"
)

// JwtCustomClaims jwt自定义负载
type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

// JwtCustomRefreshClaims jwt自定义更新后的负载
type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
