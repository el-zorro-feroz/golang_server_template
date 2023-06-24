package cmd

import (
	"cleanarch/src/features/user"
	"context"
	"fmt"

	"go.uber.org/fx"
)

func loop(lifecycle fx.Lifecycle, uu user.UserUsecase) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			user, err := uu.CreateUser("Administrator")

			if err != nil {
				return err
			}

			fmt.Printf("user.Name: %v\n", user.Name)

			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("OnStop Event")

			return nil
		},
	})
}

func RunServer() error {
	app := fx.New(
		fx.Provide(user.NewUserDelivery),
		fx.Invoke(loop),
	)

	if err := app.Err(); err != nil {
		return err
	}

	app.Run()

	return nil
}
