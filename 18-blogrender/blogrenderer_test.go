package blogrender_test

import (
	"bytes"
	"io"
	"testing"

	blogrender "example.com/hello/18-blogrender"
	approvals "github.com/approvals/go-approval-tests"
)

var (
	aPost = blogrender.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
)

func TestRender(t *testing.T) {
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrender.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blogrender.Render(io.Discard, aPost)
	}
}
