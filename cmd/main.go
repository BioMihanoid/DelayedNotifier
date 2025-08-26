package main

import (
	"BioMihanoid/DelayedNotifier/internal/api"
	"BioMihanoid/DelayedNotifier/internal/config"
	"BioMihanoid/DelayedNotifier/internal/service"
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conf := config.NewConfig()
	serv := service.NewService()
	handler := api.NewHandler(serv)

	server := &http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: handler.InitRouter(),
	}

	go func() {
		log.Println("Starting server on port:" + conf.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("failed to gracefully shutdown server: %v", err)
	}
	log.Println("Server gracefully stopped")
}
