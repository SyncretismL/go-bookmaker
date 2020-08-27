package main

import (
	"bookmaker/internal/config"
	"bookmaker/internal/postgres"
	"bookmaker/internal/subscribe"
	"bookmaker/pkg/logger"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfiguration("./config/config.json")
	if err != nil {
		log.Fatalf("Could not instantiate config %+s", err)
	}

	newLogger, err := logger.NewLogger(cfg.Log)
	if err != nil {
		log.Fatalf("Could not instantiate log %+s", err)
	}

	db := postgres.New(newLogger, cfg)

	defer db.Close()

	sports, err := postgres.NewSportStorage(db)
	if err != nil {
		newLogger.Fatalf("Could not instantiate statements %+s", err)
	}

	handler := newHandler(newLogger, db.Session)
	subscr := newSubscribe(newLogger, sports)
	client := newClient(newLogger, cfg, sports)

	for _, sport := range cfg.Sports {
		adress := fmt.Sprint(cfg.Provider.Adress + sport.Sport)
		go client.worker(sport.N, adress, sport.Sport)
	}

	r := chi.NewRouter()

	handler.Routers(r)

	srv := &http.Server{
		Addr:    cfg.Http.Adress,
		Handler: r,
	}

	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			newLogger.Fatalf("server stopped %+s", err)
		}
	}()

	dbReady := make(chan struct{})

	go func() {
		for {
			if err := db.Session.Ping(); err != nil {
				time.Sleep(1 * time.Second)
			} else {
				newLogger.Debugf("db is ready")
				dbReady <- struct{}{}
				return
			}
		}
	}()

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	subscribe.RegisterSubscrServiceServer(grpcServer, subscr)
	lis, err := net.Listen("tcp", cfg.Grpc.Adress)
	if err != nil {
		newLogger.Debugf("cannot listen on ", cfg.Grpc.Adress)
	}

	go func() {
		<-dbReady
		newLogger.Debugf("grpc server started")
		err := grpcServer.Serve(lis)
		if err != nil {
			newLogger.Fatalf("grpc server stopped %+s", err)
		}
	}()

	newLogger.Debugf("server started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		newLogger.Debugf("system call:%+v", oscall)
		cancel()
	}()

	<-ctx.Done()

	newLogger.Debugf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	err = srv.Shutdown(ctxShutDown)
	if err != nil {
		newLogger.Fatalf("server shutdown failed:%+s", err)
	}

	grpcServer.GracefulStop()

	newLogger.Debugf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
