package web

import (
	"html/template"
	"io"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
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

func parseTemplates(base_path string) *template.Template {
	templates := template.New("")
	err := filepath.Walk(base_path, func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templates.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		log.Fatalf("Error loading templates: %v\n", err)
	}

	return templates
}

func SetupServer() *echo.Echo {
	// Template Parsing

	templates := parseTemplates(template_files)

	// Database Setup
	err := db.SetupDB("db/waypoint.db")
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
