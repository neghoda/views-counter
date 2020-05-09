package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/neghoda/views-couter/handlers"
	"github.com/neghoda/views-couter/middlewares"
)

var (
	logger *log.Logger
	port   *string
)

func init() {
	port = flag.String("port", ":8080", "Application port")
	logFileLoc := flag.String("lfile", "application.log", "Log file location")

	file, err := os.OpenFile(*logFileLoc, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	homeHandler, err := handlers.NewHomeHandler(logger)
	if err != nil {
		logger.Fatal(err)
	}
	rm := &middlewares.Recovery{logger}

	s := &http.Server{
		Addr:         *port,
		Handler:      rm.WithRecovery(homeHandler),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	shutdown := make(chan os.Signal)

	signal.Notify(shutdown, os.Interrupt)
	signal.Notify(shutdown, os.Kill)

	<-shutdown

	if err := s.Shutdown(context.Background()); err != nil {
		logger.Printf("HTTP server Shutdown: %v", err)
	}

}
