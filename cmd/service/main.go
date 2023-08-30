package main

import (
	"context"
	"errors"
	"github.com/Ypxd/WebService/internal/repository"
	"github.com/Ypxd/WebService/internal/server"
	"github.com/Ypxd/WebService/internal/service"
	http2 "github.com/Ypxd/WebService/internal/transport/http"
	"github.com/Ypxd/WebService/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := utils.GetConfig()

	repos, conn, err := repository.NewRepo()
	if err != nil {
		log.Fatalf("error occured load repository: %s", err.Error())
	}

	services := service.NewService(repos, conn)

	handlers := http2.NewHandlers(services)

	srv := server.NewServer(handlers)
	go func() {
		log.Printf("API server listeing at: %s:%d",
			cfg.Server.Host,
			cfg.Server.Port)

		err := srv.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Stop(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err.Error())
	}
}
