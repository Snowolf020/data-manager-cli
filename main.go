package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"data-manager-cli/config"
	"data-manager-cli/db"
	"data-manager-cli/prisma"
)

func main() {
	loadEnvVariables()

	rootCmd := &cobra.Command{
		Use:   "data-manager-cli",
		Short: "A command-line tool for managing and syncing data between different databases",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Data Manager CLI")
		},
	}

	postgresCmd := &cobra.Command{
		Use:   "postgres",
		Short: "Manage Postgres database",
		Run: func(cmd *cobra.Command, args []string) {
			db.PostgresInit()
		},
	}

	sqliteCmd := &cobra.Command{
		Use:   "sqlite",
		Short: "Manage SQLite database",
		Run: func(cmd *cobra.Command, args []string) {
			db.SqliteInit()
		},
	}

	redisCmd := &cobra.Command{
		Use:   "redis",
		Short: "Manage Redis database",
		Run: func(cmd *cobra.Command, args []string) {
			db.RedisInit()
		},
	}

	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync data between databases",
		Run: func(cmd *cobra.Command, args []string) {
			prisma.SyncData()
		},
	}

	rootCmd.AddCommand(postgresCmd)
	rootCmd.AddCommand(sqliteCmd)
	rootCmd.AddCommand(redisCmd)
	rootCmd.AddCommand(syncCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func loadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}