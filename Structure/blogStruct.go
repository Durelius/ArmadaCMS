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
	ID        int           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    int           `gorm:"column:user_id" json:"-"`
	User      User          `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Text      string        `gorm:"column:text" json:"text"`
	Title     string        `gorm:"column:title" json:"title"`
	Author    string        `gorm:"column:author" json:"author"`
	CreatedAt time.Time     `gorm:"column:created_at" json:"createdAt"`
	Tags      []BlogpostTag `gorm:"foreignKey:BlogpostID"`
}
type BlogpostTag struct {
	ID         int
	BlogpostID int // foreign key field
	Tag        string
}
