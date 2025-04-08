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
	Id        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"column:username;unique"`
	Password  string    `json:"password" gorm:"column:password"`
	Title     string    `json:"title" gorm:"column:title"`
	FullName  string    `json:"fullName" gorm:"column:full_name"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

type UserLogin struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
