package migrate

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/j3yzz/mowz/internal/config"
	"github.com/j3yzz/mowz/internal/db"
	"github.com/spf13/cobra"
	"log"
)

func main(cfg config.Config) {
	database, err := db.New(cfg.Database)
	if err != nil {
		log.Fatal("database init failed", err)
	}
	sqlDB, _ := database.DB()
	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mowz",
		driver,
	)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Print("database is up to date")
			return
		}

		log.Fatalf(err.Error())
	}
}

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "migrate",
			Short: "database migration",
			Run: func(cmd *cobra.Command, args []string) {
				main(cfg)
			},
		},
	)
}
