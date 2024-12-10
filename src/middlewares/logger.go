package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s- %s - %v ", r.URL.String(), r.Method, time.Now())
		next(w, r)
	}
}
