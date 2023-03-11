package cmd

import (
	"github.com/j3yzz/mowz/internal/cmd/migrate"
	"github.com/j3yzz/mowz/internal/cmd/server"
	"github.com/j3yzz/mowz/internal/config"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	ExitFailure = 1
)

func Execute() {
	cfg := config.New()

	root := &cobra.Command{
		Use:   "mowz",
		Short: "a backend web application.",
	}

	migrate.Register(root, cfg)
	server.Register(root, cfg)

	if err := root.Execute(); err != nil {
		log.Println("failed to execute root command")
		os.Exit(ExitFailure)
	}
}
