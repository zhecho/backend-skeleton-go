package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

// when it render as json use lower letter representation
type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	// parse command variables
	flag.IntVar(&cfg.port, "port", 8080, "port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "environment")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}
	logger.Printf("Starting API v%s on port %d in %s mode\n", version, cfg.port, cfg.env)

	// Listen and Serve
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
