package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mensurui/bloggingPlatformAPI/internals/data"
)

func (app *application) blogList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of Blogs to be put here")
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	blog := data.Blog{
		ID:        id,
		Title:     "My Day",
		Content:   "The content is very long so make sure you are ready to GO(pun intented)",
		Tag:       []string{"life", "blog", "vlog"},
		CreatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, blog, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tag     []string `json:"tag"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	fmt.Fprintf(w, "%+v\n", input)
}
