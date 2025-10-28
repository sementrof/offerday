package App

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sementrof/offerday/internal/api"
	v1 "github.com/sementrof/offerday/internal/api/v1"
	"github.com/sementrof/offerday/internal/config"
	"github.com/sementrof/offerday/internal/deps"
	"go.uber.org/zap"
)

func Run() {
	ctx := context.Background()
	cfg := config.LoadConfig()

	depends, err := deps.ProvideDependencies(ctx, cfg)

	if err != nil {
		log.Fatalf("Failed to provide dependencies: %v", err)

	}
	apiSetup := v1.NewApi(depends)
	router := api.SetupRouter(apiSetup, depends.Logger)

	go func() {
		if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
			depends.Logger.Error("Can't run server", zap.Error(err))
			return
		}
		depends.Logger.Info("Server is running on port", zap.String("port", cfg.Port))
	}()
	depends.Logger.Info("Server is running", zap.String("port", cfg.Port))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

}
