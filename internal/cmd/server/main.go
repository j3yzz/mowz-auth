package server

import (
	"errors"
	"github.com/j3yzz/mowz/internal/config"
	"github.com/j3yzz/mowz/internal/db"
	"github.com/j3yzz/mowz/internal/http/handler"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(cfg config.Config) {
	_, err := db.New(cfg.Database)
	if err != nil {
		log.Fatal("database init failed.")
	}

	app := echo.New()

	handler.Health{}.Register(app.Group(""))

	if err := app.Start(":8082"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("echo init failed")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
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
