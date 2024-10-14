package config

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunMigrations() {
	cmd := exec.Command("atlas", "migrate", "status", "--url", os.Getenv("DATABASE_URL"))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error checking migration status: %s: %v", string(output), err)
		return
	}

	log.Println("Migration status output:", string(output))

	if strings.Contains(string(output), "no migrations to apply") {
		log.Println("No new migrations to apply.")
		return
	}

	cmd = exec.Command("atlas", "migrate", "apply", "--dir", "file://migrations", "--url", os.Getenv("DATABASE_URL"))
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
