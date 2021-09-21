package blogposts_test

import (
	"learngowithtests/rf/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
I love go`
	secondBody = `Title: Post 2
Description: Description 2
Tags: ruby, web
---
All we have to decide is what to do with the time that is given us.`
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":   {Data: []byte(firstBody)},
		"hello-world-2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	assertNoErr(t, err)
	assertPostsLength(t, len(posts), len(fs))
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body:        "I love go",
	})
}

func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertPostsLength(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d posts, want %d", got, want)
	}
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
