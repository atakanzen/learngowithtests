package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagsSeperator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tag string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tag)
	}

	return Post{
		Title:       readMetaLine(titleSeperator),
		Description: readMetaLine(descriptionSeperator),
		Tags:        strings.Split(readMetaLine(tagsSeperator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // Ignore the `---` line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
