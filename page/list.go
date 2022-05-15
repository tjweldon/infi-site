package page

import (
	"fmt"
	"strings"
)

type Item struct {
	content fmt.Stringer
}

func NewItem(content fmt.Stringer) Item {
	return Item{content: content}
}

func (i Item) String() string {
	return fmt.Sprintf("<li>%s</li>", i.content)
}

type List struct {
	items []Item
}

func ListFromCurrentPath(path string, count int) List {
	items := make([]Item, count)
	for i := range items {
		relPath := strings.TrimRight(path, "/") + "/" + fmt.Sprintf("%x", i)
		link := NewLink(relPath, relPath)
		items[i] = NewItem(link)
	}
	return List{items: items}
}

func (l List) String() string {
	out := ""
	out += fmt.Sprintf("<ul>\n")
	for _, item := range l.items {
		out += fmt.Sprintf("\t%s\n", item.String())
	}
	out += fmt.Sprintf("</ul>\n")
	return out
}

func (l List) Bytes() []byte {
	return []byte(l.String())
}
