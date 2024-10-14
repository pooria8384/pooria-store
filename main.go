package main

import (
	"fmt"
	"my-project/api"
	"my-project/config"
	"my-project/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	defaultDB := config.ConnectDefaultDB()
	defer defaultDB.Close()

	err := config.CreateDatabaseIfNotExists(defaultDB, "seed-user")
	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}

	postgresDb := config.ConnectToSeedUserDB()
	defer postgresDb.Close()

	userStore := db.NewPostGresUserStore(postgresDb)
	userHandler := api.NewUserHandler(userStore)
	authHandler := api.NewAuthHandler(userHandler)

	userStore.CreateUser("pooria", "pooria@pooria.com", "securepass")
	userStore.CreateUser("james", "james@james.com", "154pass")
	userStore.CreateUser("mark", "mark@pmark.com", "3256pass")

	app.Post("/login", authHandler.Login)
	app.Listen(":3000")
}
