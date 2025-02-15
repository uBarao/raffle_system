package utils

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}