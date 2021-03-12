package books

import (
	"encoding/json"
	"strings"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b *Book) CategoryByLength() string {
	if b.Pages >= 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}

func NewBookFromJSON(str string) (*Book, error) {
	var b Book
	if err := json.NewDecoder(strings.NewReader(str)).Decode(&b); err != nil {
		return nil, err
	}
	return &b, nil
}
