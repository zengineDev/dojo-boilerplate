package entities

import (
	"github.com/gofrs/uuid"
	"github.com/zengineDev/dojo"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func (u *User) GetAuthType() dojo.AuthUserType {
	foo := dojo.UserUserType
	return foo
}

func (u *User) GetAuthID() uuid.UUID {
	return u.ID
}

func (u *User) GetAuthData() interface{} {
	return make(map[string]interface{})
}

func (u User) IsGuest() bool {
	return u.GetAuthType() == dojo.GuestUserType
}
