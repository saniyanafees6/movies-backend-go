package main

import "net/http"

func (app *application) enableCORS(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Controll-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}