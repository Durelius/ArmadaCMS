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
	Id           *int      `gorm:"column:id;primaryKey;autoIncrement"`
	RefreshToken *string   `gorm:"column:refresh_token"`
	ValidFrom    time.Time `gorm:"column:valid_from"`
	ValidTo      time.Time `gorm:"column:valid_to"`
	UserID       int       `gorm:"column:user_id"`
	User         User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Enabled      bool      `gorm:"column:enabled"`
}

type Tokens struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
