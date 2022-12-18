package auth

import (
	"context"
	"errors"
)

var (
	ErrorInvalidToken = errors.New("invalid token")
)

type TokenInfo interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresAt() int64
	EncodeToJSON() ([]byte, error)
}

type Auth interface {
	GenerateToken(ctx context.Context, userId string) (TokenInfo, error)
	DestroyToken(ctx context.Context, accessToken string) error
	ParseUserID(ctx context.Context, accessToken string) (string, error)
	Release() error
}
