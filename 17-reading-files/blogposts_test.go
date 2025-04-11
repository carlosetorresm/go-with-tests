package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "example.com/hello/17-reading-files"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("open files", func(t *testing.T) {
		const (
			firstBody  = "Title: Post 1\nDescription: Description 1"
			secondBody = "Title: Post 2\nDescription: Description 2"
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		got := posts[0]
		want := blogposts.Post{Title: "Post 1", Description: "Description 1"}

		assertPost(t, got, want)
	})

	t.Run("should fail to open file", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Error("Should've shown an error")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
