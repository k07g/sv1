package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/k07g/sv1/internal/pkg/infrastructure"
)

func Run() {
	cfg := &infrastructure.Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	fmt.Println("Port:", cfg.Port, "Environment:", cfg.Environment)

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer cancel()

	addr := fmt.Sprintf(":%s", cfg.Port)
	server := http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Println("シグナル", <-quit, "を受け取ったのでサーバをシャットダウンしています...")

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalln("failed to gracefully shutdown", err)
	}

	log.Println("サーバのシャットダウンが完了しました")
}
