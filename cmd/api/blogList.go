package main

import (
	"fmt"
	"net/http"
)

func (app *application) blogList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of Blogs to be put here")
}
