package domain

import (
	"github.com/lmurature/api-boilerplate/cmd/api/domain/users"
	"go.uber.org/fx"
)

var DaoModule = fx.Options(fx.Provide(users.NewUserDao))
