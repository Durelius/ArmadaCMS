package Structure

import "time"

//	type RefreshTokenDB struct {
//		Id           *int      `db:"id"`
//		RefreshToken *string   `db:"refresh_token"`
//		ValidFrom    time.Time `db:"valid_from"`
//		ValidTo      time.Time `db:"valid_to"`
//		UserID       int       `db:"user_id"`
//		Enabled      *bool     `db:"enabled"`
//	}
type RefreshTokenDB struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	RefreshToken string    `json:"refreshToken" gorm:"unique;not null"`
	ValidFrom    time.Time `json:"validFrom"`
	ValidTo      time.Time `json:"validTo"`
	UserID       int       `json:"userId" gorm:"not null"`
	User         User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Enabled      bool      `json:"enabled"`
}

type Tokens struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
