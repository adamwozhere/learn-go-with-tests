package blogposts_test

import (
	// "17-reading-files/blogposts"
	"blogposts"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

// we can use `fstest` tp create an in-memory filesystem for testing
// rather than creating separate folder structure of .md files.
// the `fstest` uses the same interface as `fs` (`io.fs`)

// 1. first just check we get the right number of Posts back.
// func TestNewBlogPosts(t *testing.T) {
// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte("hi")},
// 		"hello-world2.md": {Data: []byte("hola")},
// 	}

// 	posts, err := blogposts.NewPostsFromFS(fs)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(posts) != len(fs) {
// 		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
// 	}
// }

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestFSError(t *testing.T) {
	_, err := blogposts.NewPostsFromFS(StubFailingFS{})

	if err == nil {
		t.Error("expected an error but got nil")
	}
}

// technically not worth writing the error handling test as we don't do
// anything with it, we just propgate it so it's not worth writing the test?

// 2. parse a post and check the title
// func TestNewBlogPosts(t *testing.T) {
// 	fs := fstest.MapFS{
// 		"hello world.md": {Data: []byte("Title: Post 1")},
// 		"helo-world2.md": {Data: []byte("Title: Post 2")},
// 	}

// 	posts, err := blogposts.NewPostsFromFS(fs)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assertPost(t, posts[0], blogposts.Post{Title: "Post 1"})
// }

// 3. extend the test to check the description also
func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
A
B
C`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		// add tags
		Tags: []string{"tdd", "go"},
		// add body
		Body: `Hello
World`,
	})
}

// helper assert function
func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
