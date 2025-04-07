package Structure

import "time"

type NewBlogpost struct {
	Text   string   `json:"text"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}
type Blogpost struct {
	Id        int       `json:"id" db:"id"`
	UserId    int       `json:"userId" db:"user_id"`
	Text      string    `json:"text" db:"text"`
	Title     string    `json:"title" db:"title"`
	Author    string    `json:"author" db:"author"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	Tags      []string  `json:"tags" db:"tags"`
}
