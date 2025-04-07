package Structure

import "time"

type RefreshTokenDB struct {
	Id           *int      `db:"id"`
	RefreshToken *string   `db:"refresh_token"`
	ValidFrom    time.Time `db:"valid_from"`
	ValidTo      time.Time `db:"valid_to"`
	UserID       int       `db:"user_id"`
	Enabled      *bool     `db:"enabled"`
}
type Tokens struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
