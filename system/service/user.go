package service

import (
	"context"

	"github.com/crusttech/crust/system/types"
	"github.com/crusttech/crust/system/internal/service"
)

type (
	UserService interface {
		FindByUsername(username string) (*types.User, error)
		FindByEmail(email string) (*types.User, error)
		FindByID(id uint64) (*types.User, error)
		FindByIDs(id ...uint64) (types.UserSet, error)
		Find(filter *types.UserFilter) (types.UserSet, error)
	}
)

var DefaultUser = service.DefaultUser

func User(ctx context.Context) UserService {
	return DefaultUser.With(ctx)
}
