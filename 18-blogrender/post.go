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

type PostRenderer struct {
	templ *template.Template
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}
	return nil
}
