package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
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

func main() {
	var cfg config
	// parse command variables
	flag.IntVar(&cfg.port, "port", 8080, "port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "environment")
	flag.Parse()

	fmt.Printf("Starting API v%s on port %d in %s mode\n", version, cfg.port, cfg.env)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
			Version:     version,
		}
		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	// Listen and Serve
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}

}
