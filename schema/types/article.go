package types

import "time"

type Article struct {
	ID           uint      `json:"id,omitempty"`
	Created_at   time.Time `json:"created_at"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug"`
	Content      string    `json:"content"`
	Published_at time.Time `json:"published_at"`
	Public       bool      `json:"public"`
	Thumbmail    string    `json:"thumbnail"`
}

func (a Article) Save() {
	a.Slug = slugify(a.Title)
	// TODO: save model
}
