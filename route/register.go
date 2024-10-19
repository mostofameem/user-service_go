package route

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"time"
	"user-service/db"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Dob      string
	Type     string
}

func (svc *service) Register(ctx context.Context, user User) error {
	_, err := svc.userRepo.GetProfileWithEmail(ctx, user.Email)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("internal Server error")
	}

	err = svc.userRepo.Register(&ctx, db.User{
		Name:       user.Name,
		Email:      user.Email,
		Pass:       user.Password,
		Dob:        user.Dob,
		Type:       user.Type,
		Is_active:  "false",
		Created_at: time.Now(),
		Updated_at: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func generateSecretCode() string {
	const min = 10000
	const max = 99999

	// Calculate the range size
	rangeSize := max - min + 1

	// Generate a random number in the range 0 to (rangeSize - 1)
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(rangeSize)))
	if err != nil {
		// Handle error
		fmt.Println("Error generating random number:", err)
		return ""
	}

	// Add the min value to shift the range to [min, max]
	code := int(nBig.Int64()) + min

	return fmt.Sprintf("%05d", code)
}
