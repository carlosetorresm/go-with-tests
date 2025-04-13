package blogrender

import (
	"embed"
	"io"
	"text/template"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, post Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}
	if err := templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}
	return nil
}
