package auth

import "time"

type Claims struct {
	UID      int64 `json:"uid"`
	ExpireAt int64 `json:"expire_at"`
}

func (c *Claims) Valid() error {
	if time.Now().Unix() > c.ExpireAt {
		return unauthenticatedErr
	}

	return nil
}
