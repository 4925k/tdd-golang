package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSep       = "Title: "
	descriptionSep = "Description: "
	tagSep         = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostFromFS(files fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(files, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(files, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(files fs.FS, name string) (Post, error) {
	file, err := files.Open(name)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	return newPost(file)
}

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readLine(titleSep),
		Description: readLine(descriptionSep),
		Tags:        strings.Split(readLine(tagSep), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
