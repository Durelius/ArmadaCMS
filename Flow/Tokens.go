package Flow

import (
	"ArmadaCMS/Structure"
	utilities "ArmadaCMS/Utilities"
	"ArmadaCMS/db"
	"log"
	"net/http"
	"strings"
	"time"
)

const validDuration = 7   //refresh token is valid for 7 days
const updateDuration = -5 //update the token when it is 5 days left

func VerifyRefreshToken(tokenString string, userId int) *Structure.RefreshTokenDB {
	rToken, err := db.MatchRefreshToken(tokenString, userId)
	if err != nil {
		log.Println(err)
		return nil
	}

	if time.Now().Unix() > rToken.ValidTo.Unix() {
		log.Println("Token expired")
		return nil
	}

	updateWhen := rToken.ValidTo.AddDate(0, 0, updateDuration)
	if time.Now().Unix() > updateWhen.Unix() {
		db.ExtendRefreshToken(rToken, validDuration)
	}

	return rToken
}

// Returns true if we should use refresh token to generate new access token
func VerifyAccessTokenWebRequest(r *http.Request) (*int, bool) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return nil, true
	}

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, false
	}

	mapClaims, err := utilities.VerifyAccessToken(parts[1])
	if err != nil {
		if err.Error() == "token has invalid claims: token is expired" {
			return nil, true
		}
	}
	floatUserId, ok := (*mapClaims)["user_id"].(float64)
	if !ok {
		return nil, false
	}
	userId := int(floatUserId)

	return &userId, false

}
