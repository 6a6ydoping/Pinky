package app

import (
	"github.com/6a6ydoping/Pinky/internal/config"
	"github.com/6a6ydoping/Pinky/internal/handler"
	"github.com/6a6ydoping/Pinky/internal/repository/pgrepo"
	"github.com/6a6ydoping/Pinky/internal/service"
	"github.com/6a6ydoping/Pinky/pkg/httpserver"
	"github.com/6a6ydoping/Pinky/pkg/jwttoken"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	db, err := pgrepo.New(
		pgrepo.WithHost(cfg.DB.Host),
		pgrepo.WithPort(cfg.DB.Port),
		pgrepo.WithDBName(cfg.DB.DBName),
		pgrepo.WithUsername(cfg.DB.Username),
		pgrepo.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	log.Println("connection success")
	token := jwttoken.New([]byte(cfg.Token.SecretKey))
	service := service.New(db, token, cfg)
	handler := handler.New(service)
	server := httpserver.New(
		handler.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)
	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
