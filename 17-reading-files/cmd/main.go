package main

import (
	"blogposts"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}

// you can run this file using `go run .` from within the `cmd` folder.
// notice that it's ealy now to change the filesystem that we pass into `NewPostsFromFS`
// and it all works correctly as it uses the same interface as with the tests.
