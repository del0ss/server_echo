package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Note struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Text    string `json:"text"`
}

var Notes = []Note{}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.GET("/notes", func(c echo.Context) error {
		return c.Render(http.StatusOK, "notes.html", nil)
	})

	e.GET("/create_notes", func(c echo.Context) error {
		return c.Render(http.StatusOK, "create.html", nil)
	})

	e.POST("/create_notes", func(c echo.Context) error {
		name := c.FormValue("name")
		surname := c.FormValue("surname")
		text := c.FormValue("text")
		Notes = append(Notes, Note{name, surname, text})
		return c.Render(http.StatusOK, "create.html", nil)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
