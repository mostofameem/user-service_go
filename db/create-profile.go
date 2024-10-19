package db

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"log/slog"
	"time"
	"user-service/logger"
)

type User struct {
	Id         int       `db:"id"`
	Name       string    `db:"name"`
	Email      string    `db:"email" `
	Pass       string    `db:"password"`
	Dob        string    `db:"dob"`
	Type       string    `db:"type"`
	Is_active  string    `db:"is_active"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

func (r *UserTypeRepo) Register(ctx *context.Context, usr User) error {
	columns := map[string]interface{}{
		"name":       usr.Name,
		"email":      usr.Email,
		"pass":       hashPassword(usr.Pass),
		"dob":        usr.Dob,
		"type":       usr.Type,
		"is_active":  false,
		"created_at": usr.Created_at,
		"updated_at": usr.Updated_at,
	}
	var colNames []string
	var colValues []any

	for colName, colVal := range columns {
		colNames = append(colNames, colName)
		colValues = append(colValues, colVal)
	}

	query, args, err := NewQueryBuilder().
		Insert(r.table).
		Columns(colNames...).
		Values(colValues...).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		slog.Error(
			"Failed to create New user insert query",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"query": query,
				"args":  args,
			}),
		)
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

func hashPassword(pass string) string {

	h := sha1.New()
	h.Write([]byte(pass))
	hashValue := h.Sum(nil)
	return hex.EncodeToString(hashValue)
}
