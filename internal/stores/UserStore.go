package stores

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/zengineDev/dojo/db"
	"main/internal/domain/entities"
)

type UserStore struct {
	db.PostgresStore
}

func (store UserStore) FindByEmail(ctx context.Context, email string) (entities.User, error) {
	var result entities.User
	sql, args, err := store.SB.Select("id", "email", "password").
		From("users").Where(squirrel.Eq{"email": email}).ToSql()
	if err != nil {
		return result, err
	}

	row := store.DB.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&result.ID, &result.Email, &result.Password)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			//logx.Log.Error(err) // TODO return an notfound error
			return result, err
		}
	}

	return result, nil
}
