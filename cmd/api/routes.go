package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	route := httprouter.New()
	route.HandlerFunc(http.MethodGet, "/v1/bloglist", app.blogList)
	route.HandlerFunc(http.MethodGet, "/v1/blog/:id", app.blog)
	route.HandlerFunc(http.MethodPost, "/v1/blog", app.createMovieHandler)

	return route
}
