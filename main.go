package main

import (
	"my-project/config"
	"my-project/handlers"
	"my-project/server"
	"my-project/storer"
)

func init() {
	config.LoadEnv()
}

func main() {

	postgresDb := config.ConnectDB()
	defer postgresDb.Close()

	userStore := storer.NewPostGresUserStore(postgresDb)
	userHandler := handlers.NewUser(userStore)
	authHandler := handlers.NewAuthUser(userHandler)

	userStore.CreateUser("pooria", "pooria@pooria.com", "securepass")
	userStore.CreateUser("james", "james@james.com", "154pass")
	userStore.CreateUser("mark", "mark@pmark.com", "3256pass")

	srv := server.NewServer(userHandler, authHandler)
	if err := srv.Start(":3000"); err != nil {
		panic(err)
	}
}
