package main

import (
	"authn/api/handler"
	"authn/repo"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find .env file: ", err)
	}

	dbURI := viper.GetString("DB_URI")

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(dbURI),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("test_auth")
	userRepo := repo.NewMongoUserRepo(db)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.POST("/login", func(c echo.Context) error {
		c.Set("TOKEN_SECRET", viper.GetString("TOKEN_SECRET"))

		return handler.Login(c, userRepo)
	})

	port := viper.GetString("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
