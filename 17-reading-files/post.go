package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	scanner.Scan()
	titleLine := scanner.Text()
	scanner.Scan()
	DescriptionLine := scanner.Text()
	post := Post{Title: titleLine[7:], Description: DescriptionLine[13:]}
	return post, nil
}
