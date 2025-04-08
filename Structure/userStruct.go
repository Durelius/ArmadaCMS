package Structure

import "time"

//	type User struct {
//		Id        int       `json:"id" db:"id"`
//		Username  string    `json:"username" db:"username"`
//		Password  string    `json:"password" db:"password"`
//		Title     string    `json:"title" db:"title"`
//		FullName  string    `json:"fullName" db:"full_name"`
//		CreatedAt time.Time `json:"createdAt" db:"created_at"`
//	}
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Title     string    `json:"title"`
	FullName  string    `json:"fullName"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
