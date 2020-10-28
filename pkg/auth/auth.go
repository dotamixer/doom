package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth struct {
	opts *Options
}

var unauthenticatedErr = status.Errorf(codes.Unauthenticated, "Request unauthenticated")

var CTXUID = "uid"

func (a *Auth) AuthFunc(ctx context.Context) (context.Context, error) {
	var (
		err    error
		token  *jwt.Token
		claims = &Claims{}
	)
	tokenString := metautils.ExtractIncoming(ctx).Get(a.opts.HeaderAuthorize)
	if tokenString == "" {
		return ctx, unauthenticatedErr
	}

	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.opts.Secret), nil
	})

	if err != nil {
		return ctx, unauthenticatedErr
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		ctx = context.WithValue(ctx, CTXUID, claims.UID)
	} else {
		return ctx, unauthenticatedErr
	}

	return ctx, err
}

func (a *Auth) GenerateAuth(uid int64) (string, error) {
	claims := &Claims{
		UID:      uid,
		ExpireAt: time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.opts.Secret))
}
