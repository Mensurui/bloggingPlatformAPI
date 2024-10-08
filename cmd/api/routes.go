package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	route := httprouter.New()
	route.HandlerFunc(http.MethodPost, "/v1/blog", app.blogCreateHandler)
	route.HandlerFunc(http.MethodGet, "/v1/blog/:id", app.showBlogHandler)
	route.HandlerFunc(http.MethodPatch, "/v1/blog/:id", app.updateBlogHandler)
	route.HandlerFunc(http.MethodDelete, "/v1/blog/:id", app.deleteBlogHandler)
	route.HandlerFunc(http.MethodGet, "/v1/blog", app.listBlogHandler)

	return route
}
