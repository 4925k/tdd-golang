package blog_test

import (
	blog "blogTesting/blog"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFs struct {
}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("open failed")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
hellooooooo`
		secondBody = `Title: Post 2
Description: Description 2
Tags: dibek
---
hwlooooe`
	)

	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blog.NewPostFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blog.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body:        `hellooooooo`,
	})

}

func assertPost(t *testing.T, got, want blog.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}
