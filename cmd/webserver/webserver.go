package main

import (
	"github.com/chris-metz/dbutler/lib/db"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type Templates struct {
	templates map[string]*template.Template
}

func NewTemplates() *Templates {
	var templates = make(map[string]*template.Template)
	templates["home"] = template.Must(template.ParseFiles("templates/layout.gohtml", "templates/home.gohtml"))
	templates["hello"] = template.Must(template.ParseFiles("templates/layout.gohtml", "templates/hello.gohtml"))
	return &Templates{
		templates: templates,
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].ExecuteTemplate(w, "layout", data)
}

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", "")
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "yo!")
}

func main() {
	dbHandler := db.NewDbHandler()
	defer dbHandler.Shutdown()
	dbHandler.ReCreateSchema()
	dbHandler.SeedDatabase()

	e := echo.New()
	e.Renderer = NewTemplates()

	e.GET("/", Home)
	e.GET("/hello", Hello)
	e.Logger.Fatal(e.Start(":8582"))
}
