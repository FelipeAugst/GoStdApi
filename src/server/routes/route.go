package routes

import "net/http"

type Route struct {
	Url     string
	Method  string
	Handler http.HandlerFunc
}

func (r Route) GetWildCard() string {
	return r.Method + " " + r.Url
}
