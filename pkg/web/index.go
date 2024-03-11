package web

import (
	"html/template"
	"io"
	"log"
	"path"
	"waypoint/pkg/db"
	"waypoint/pkg/web/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const template_files string = "views"

type TemplateRenderer struct {
	template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}

func SetupServer() *echo.Echo {
	// Template Parsing
	templates, err := template.ParseGlob(path.Join(template_files, "*.html"))
	if err != nil {
		log.Fatalf("Error loading templates: %v\n", err)
	}

	// Database Setup
	err = db.SetupDB("db/waypoint.db")
	if err != nil {
		log.Fatalf("Error creating db: %v\n", err)
	}

	// Setup Echo server
	e := echo.New()

	e.Renderer = &TemplateRenderer{
		template: templates,
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static Files
	e.Static("/js", "public/js")
	e.Static("/css", "public/css")
	e.Static("/img", "public/img")

	// Routing
	routes.Routing(e)

	return e
}
