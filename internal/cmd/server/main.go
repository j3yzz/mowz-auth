package server

import (
	"github.com/j3yzz/mowz/internal/config"
	"github.com/j3yzz/mowz/internal/db"
	"github.com/spf13/cobra"
	"log"
)

func main(cfg config.Config) {
	_, err := db.New(cfg.Database)
	if err != nil {
		log.Fatal("database init failed.")
	}
}

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "run server to serve the requests.",
			Run: func(cmd *cobra.Command, args []string) {
				main(cfg)
			},
		},
	)
}
