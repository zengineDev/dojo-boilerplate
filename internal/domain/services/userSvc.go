package services

import (
	"context"
	"main/internal/domain/entities"
)

type UserPersistenceStore interface {
	FindByEmail(ctx context.Context, email string) (entities.User, error)
}

type UserSvc struct {
	Store UserPersistenceStore
}

func NewUserSvc(persistenceStore UserPersistenceStore) *UserSvc {
	return &UserSvc{
		Store: persistenceStore,
	}
}

func (svc *UserSvc) FindByEmail(ctx context.Context, email string) (entities.User, error) {
	return svc.Store.FindByEmail(ctx, email)
}
