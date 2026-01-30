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
)

func Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer cancel()

	addr := fmt.Sprintf(":%s", "8081")
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
