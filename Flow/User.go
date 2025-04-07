package Flow

import (
	"ArmadaCMS/Structure"
	"ArmadaCMS/Utilities"
	"ArmadaCMS/db"
	"errors"
)

func VerificationFlow(user Structure.User) (*Structure.Tokens, error) {
	userId, err := db.VerifyUser(user)
	if err != nil {
		return nil, errors.New("not authenticated (1)")
	}
	refreshToken, err := Utilities.GenerateRefreshToken()
	if err != nil {
		return nil, errors.New("not authenticated (2)")
	}
	if !db.InsertRefreshToken(userId, refreshToken) {
		return nil, errors.New("not authenticated (3)")
	}

	accessToken, _ := Utilities.GenerateAccessToken(userId)

	return &Structure.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
