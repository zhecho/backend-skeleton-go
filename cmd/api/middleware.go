package main

import "net/http"

func (app application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// if r.Method == "OPTIONS" {
		// return
		// }
		next.ServeHTTP(w, r)
	})
}
