package page

import "fmt"

type Link struct {
	text, href string
}

func NewLink(text, href string) Link {
	return Link{text: text, href: href}
}

func (l Link) String() string {
	return fmt.Sprintf("<a href=\"%s\">%s</a>", l.href, l.text)
}
