package views

import "net/http"

func Error(w http.ResponseWriter, statusCode int, err error) {
	var jsonError = map[string]string{"Error": err.Error()}
	ToJSON(w, statusCode, jsonError)

}
