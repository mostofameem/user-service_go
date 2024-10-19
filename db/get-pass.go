package db

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

func (repo *UserTypeRepo) GetPass(ctx context.Context, email string) (string, error) {
	query, args, err := NewQueryBuilder().
		Select("pass").
		From(repo.table).
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	var pass string
	err = repo.db.QueryRowContext(ctx, query, args...).Scan(&pass)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	return pass, nil
}
