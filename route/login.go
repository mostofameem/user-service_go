package route

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
)

func (svc *service) Login(ctx context.Context, email, pass string) error {

	dbpass, err := svc.userRepo.GetPass(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("no user found")
		}
		return fmt.Errorf("internal server error: %v", err)
	}

	if hashPassword(pass) != dbpass {
		return fmt.Errorf("wrong username / password")
	}

	return nil
}

func hashPassword(pass string) string {

	h := sha1.New()
	h.Write([]byte(pass))
	hashValue := h.Sum(nil)
	return hex.EncodeToString(hashValue)
}
