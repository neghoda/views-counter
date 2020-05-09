package main

import (
	"flag"
	"log"
	"net/http"
	"os"

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
	http.ListenAndServe(*port, rm.WithRecovery(homeHandler))
}
