package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/k07g/sv1/internal/pkg/infrastructure"
	ingrpc "github.com/k07g/sv1/internal/pkg/infrastructure/grpc"
)

func Run() {
	cfg := &infrastructure.Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	fmt.Println("Port:", cfg.Port, "Environment:", cfg.Environment)
	fmt.Println("gRPC Port:", cfg.GRPCPort)

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer cancel()

	addr := fmt.Sprintf(":%s", cfg.Port)
	server := http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 10 * time.Second,
	}

	grpcServer := ingrpc.NewServer(ctx)
	grpcLis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
	if err != nil {
		panic(err)
	}

	go func() {
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalln("gRPCサーバの起動に失敗しました")
		}
	}()

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("HTTPサーバの起動に失敗しました")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Println("シグナル", <-quit, "を受け取ったのでサーバをシャットダウンしています...")

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalln("シャットダウンに失敗しました", err)
	}
	grpcServer.GracefulStop()

	log.Println("サーバのシャットダウンが完了しました")
}
