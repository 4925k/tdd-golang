package blog_test

import (
	"blogTesting/blog"
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		postX = blog.Post{
			Title:       "hello world",
			Body:        "konichiwa hello namaste",
			Description: "saying hello around the world",
			Tags:        []string{"japan", "usa", "nepal"},
		}
	)

	t.Run("render a HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blog.Render(&buf, postX)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1><p>saying hello around the world</p>Tags: <ul><li>japan</li><li>usa</li><li>nepal</li></ul>`
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
