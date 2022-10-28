package main

import (
	"context"
	"foo/bootstrap"
	"foo/facades"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// bootstrap
	if err := bootstrap.Boot(); err != nil {
		panic(err.Error())
	}
	server := &http.Server{
		Addr:    facades.Viper.GetString("ADDR"),
		Handler: facades.Handler,
	}
	// start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err.Error())
		}
	}()
	zap.L().Info("server started!")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err.Error())
	}
}
