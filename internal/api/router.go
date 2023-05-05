package api

import (
	"github.com/dom/user/internal/config"
	"github.com/dom/user/internal/handler"
	"github.com/dom/user/internal/users"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type SetupRouterParams struct {
}

type Router struct {
	AppRouter *echo.Echo
}

func SetupRouter(db *sqlx.DB) (*Router, error) {
	echo := echo.New()

	jwtConfig := config.DefaultJWTConfig()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	authUserParams := handler.AuthUserHandlerParams{
		Query: users.NewUserQuery(db),
	}
	echo.POST("/auth", handler.NewAuthUserHandler(&authUserParams).Handle)

	echo.GET("/_alive", handler.HealthAlive)

	usersGroup := echo.Group("/users")
	usersGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    jwtConfig.SecretKey,
	}))
	getUsersParams := handler.GetUsersHandlerParams{
		Query: users.NewUserQuery(db),
	}
	usersGroup.GET("/:id", handler.NewGetUsersHandler(&getUsersParams).Handle)
	usersGroup.GET("", handler.NewGetUsersHandler(&getUsersParams).Handle)
	addUserParams := handler.AddUserHandlerParams{
		Cmd: users.NewUserCommand(db),
	}
	usersGroup.POST("/new", handler.NewAddUserHandler(&addUserParams).Handle)
	kafkaParams := handler.KafkaHandlerParams{}
	usersGroup.POST("/kafka", handler.NewKafkaHandler(&kafkaParams).Handle)

	r := &Router{
		AppRouter: echo,
	}

	return r, nil
}
