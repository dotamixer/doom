package auth

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

var defaultAuth = &Auth{
	opts: defaultOptions(),
}

func BuildAuthFunc() grpc_auth.AuthFunc {
	return defaultAuth.AuthFunc
}

func GenerateAuthToken(uid int64) (string, error) {
	return defaultAuth.GenerateAuth(uid)
}
