package page

import "strings"

type Head struct {
}

func (h Head) String() string {
	elems := []string{
		"<head>",
		"<meta charset=\"UTF-8\">",
		"<title>Infi-site</title>",
	}
	return strings.Join(elems, "\n\t") + "\n</head>"
}

type Body struct {
	content string
}

func NewPage(content string) Page {
	return Page{Head: Head{}, Body: Body{content: content}}
}

func (b Body) String() string {
	return b.content
}

type Page struct {
	Head Head
	Body Body
}

func (p Page) String() string {
	elems := []string{
		"<!DOCTYPE html>",
		p.Head.String(),
		p.Body.String(),
	}
	return strings.Join(elems, "\n")
}

func (p Page) Bytes() []byte {
	return []byte(p.String())
}
