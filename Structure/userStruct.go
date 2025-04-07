package Structure

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Title     string    `json:"title" db:"title"`
	FullName  string    `json:"fullName" db:"full_name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type UserLogin struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
