package entities

import (
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/zengineDev/dojo"
	"github.com/zengineDev/dojo/db"
)

type User struct {
	db.Model
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func (u User) FindByEmail(ctx dojo.Context, email string) error {
	u.Init()
	sql, args, err := u.SB.Select("id", "email", "password").
		From("users").Where(squirrel.Eq{"email": email}).ToSql()
	if err != nil {
		return err
	}

	row := u.DB.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			//logx.Log.Error(err) // TODO return an notfound error
			return err
		}
	}

	return nil

}
