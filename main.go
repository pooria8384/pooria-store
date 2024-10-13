package main

import (
	"fmt"
	"my-project/config"
	"my-project/db"
)

func main() {

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
	userStore.CreateUser("pooria", "pooria@pooria.com", "securepass")
	userStore.CreateUser("james", "james@james.com", "154pass")
	userStore.CreateUser("mark", "mark@pmark.com", "3256pass")

}
