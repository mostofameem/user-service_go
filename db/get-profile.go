package db

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

func (repo *UserTypeRepo) GetProfile(ctx context.Context, id int) (*User, error) {
	query, args, err := NewQueryBuilder().
		Select("id,name,email,dob,type,is_active").
		From(repo.table).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	var userInfo User
	err = repo.db.QueryRowContext(ctx, query, args...).Scan(
		&userInfo.Id,
		&userInfo.Name,
		&userInfo.Email,
		&userInfo.Dob,
		&userInfo.Type,
		&userInfo.Is_active,
	)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &userInfo, nil
}
