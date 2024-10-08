package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Mensurui/bloggingPlatformAPI/internals/data"
	"github.com/Mensurui/bloggingPlatformAPI/internals/data/validator"
)

func (app *application) blogCreateHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tag     []string `json:"tag"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	blog := &data.Blog{
		Title:     input.Title,
		Content:   input.Content,
		Tag:       input.Tag,
		UpdatedAt: time.Now(),
	}

	err = app.models.Blogs.Insert(blog)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/blogs/%d", blog.ID))

	err = app.writeJSON(w, http.StatusOK, blog, headers)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showBlogHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	blog, err := app.models.Blogs.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrorRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, blog, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateBlogHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	blog, err := app.models.Blogs.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrorRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title   *string  `json:"title"`
		Content *string  `json:"content"`
		Tag     []string `json:"tag"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusNoContent, err)
		return
	}

	if input.Title != nil {
		blog.Title = *input.Title
	}

	if input.Content != nil {
		blog.Content = *input.Content
	}
	if input.Tag != nil {
		blog.Tag = input.Tag
	}

	err = app.models.Blogs.Update(blog)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, blog, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) deleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Blogs.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrorRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, "successfully deleted", nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) listBlogHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()
	input.Title = app.readString(qs, "title", "")

	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Sort = app.readString(qs, "sort", "id")
	input.SortSafelist = []string{"id", "title", "content", "-id", "-title", "-content"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.notFoundResponse(w, r)
		return
	}

	blogs, err := app.models.Blogs.GetAll(input.Title) // Pass the title filter here
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, blogs, nil)
	if err != nil {
		app.methodNotAllowedResponse(w, r)
	}
}
