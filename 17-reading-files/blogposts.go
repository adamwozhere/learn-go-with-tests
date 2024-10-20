package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

// task is to create a blogsystem software.
// on startup the webserver will read a folder to create some `Post`s
// then a separate `NewHandler` function will use those as a datasource for the blogs server.

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

// 1. check we have the right number of posts by just returning a blank Post
// func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
// 	dir, err := fs.ReadDir(fileSystem, ".")

// 	if err != nil {
// 		return nil, err
// 	}

// 	var posts []Post
// 	for range dir {
// 		posts = append(posts, Post{})
// 	}

// 	return posts, nil
// }

// 2
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err // todo: needs clarification
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
// 	postFile, err := fileSystem.Open(f.Name())
// 	if err != nil {
// 		return Post{}, err
// 	}
// 	defer postFile.Close()

// 	postData, err := io.ReadAll(postFile)
// 	if err != nil {
// 		return Post{}, err
// 	}

// 	post := Post{Title: string(postData)[7:]}
// 	return post, nil
// }

// Refactor the code intoto separate functions for opening the file and parsing it
func getPost(fileSystem fs.FS, filename string) (Post, error) {
	postFile, err := fileSystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

// Not technically needed but can be helpful:
// moving `newPost` and `Post` type into a new `post.go` file to group them logically.
// keeping this here though just for the flow of seeing the process/coding iteration.

// func newPost(postFile io.Reader) (Post, error) {
// 	postData, err := io.ReadAll(postFile)
// 	if err != nil {
// 		return Post{}, err
// 	}

// 	// simply get the title by slicing the string
// 	post := Post{Title: string(postData)[7:]}
// 	return post, nil
// }

// use constants for tag names to extract data with rather than just trimming strings by a magic number
const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

// 3. we can use `Scanner` to read data such as a file of newline-delimited lines of text
// func newPost(postBody io.Reader) (Post, error) {
// 	scanner := bufio.NewScanner(postBody)

// 	// call Scan() to read a line then extract the data using Text() removing the tag name (`Title: ` etc.)
// 	readMetaLine := func(tagName string) string {
// 		scanner.Scan()
// 		return strings.TrimPrefix(scanner.Text(), tagName)
// 	}

// 	// titleLine := readLine()[7:]
// 	// descriptionLine := readLine()[13:]
// 	// after refactoring using const separators:
// 	title := readMetaLine(titleSeparator)
// 	description := readMetaLine(descriptionSeparator)
// 	// add tags: take the line and Split it to turn it into an array
// 	tags := strings.Split(readMetaLine(tagsSeparator), ", ")

// 	scanner.Scan() // ignore a line (the `---`)

// 	// add the body

// 	// scanner.Scan() returns a bool returning true if there's more data to scan, so we can loop it.
// 	// write each Scan() to the buffer using `fmt.Println` which adds a newline ("\n") character at the end.
// 	// we do this as scanner removes newlines from each line -- we then need to trim off the final trailing newline.
// 	buf := bytes.Buffer{}
// 	for scanner.Scan() {
// 		fmt.Fprintln(&buf, scanner.Text())
// 	}
// 	body := strings.TrimSuffix(buf.String(), "\n")

// 	return Post{
// 		Title:       title,
// 		Description: description,
// 		Tags: tags,
// 		Body: body,
// 	}, nil
// }

// 4. Refactor, encapsulating the `readBody` part into it's own function,
// allowing readers to understand *what* is happening in `newPost`,
// without having to concern themselves with implementation specifics.
func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignores one line (the `---` line separator)
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

// Summary:
// this all works now, extracting the meta lines, and the body
// of course there are other considerations such as different filenames, or different metadata,
// but it works and would just be a matter of iterating on futher implementations and testing--the overal design works.
