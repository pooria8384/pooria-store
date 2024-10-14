package config

import (
	"log"
	"os"
	"os/exec"
)

func RunMigrations() {
	cmd := exec.Command("atlas", "migrate", "apply", "--dir", "file://migrations", "--url", os.Getenv("DATABASE_URL"))
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
