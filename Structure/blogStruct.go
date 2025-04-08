package Structure

import (
	"time"
)

type NewBlogpost struct {
	Text   string   `json:"text"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}
type Blogpost struct {
	ID        int           `json:"id"  gorm:"primaryKey"`
	UserID    int           `json:"userId"  gorm:"not null"`
	User      User          `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Text      string        `json:"text" `
	Title     string        `json:"title" `
	Author    string        `json:"author" `
	CreatedAt time.Time     `json:"createdAt" `
	Tags      []BlogpostTag `json:"tags" gorm:"foreignKey:BlogpostID"`
}

type BlogpostTag struct {
	ID         int
	BlogpostID int // foreign key field
	Tag        string
}
