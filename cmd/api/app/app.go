package app

import (
	"go.uber.org/fx"
)

func StartApp() {
	app := fx.New(CommonModules, fx.Invoke(func(*Runner) {}))
	app.Run()
}
