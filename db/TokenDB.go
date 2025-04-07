package db

import (
	"ArmadaCMS/Structure"
	"errors"
	"fmt"
	"log"
	"time"
)

func InsertRefreshToken(userId int, refreshToken string) bool {
	if userId == 0 {
		return false
	}

	const insertJwtTokenString string = `
    INSERT INTO token (
        refresh_token, 
        user_id, 
        valid_from, 
        valid_to
    ) VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '7 days')`

	_, err := DB.Exec(insertJwtTokenString, refreshToken, userId)
	if err != nil {
		log.Printf("Error inserting refresh token: %v", err)
		return false
	}
	return true
}

func MatchRefreshToken(refreshToken string, userId int) (*Structure.RefreshTokenDB, error) {
	var rTokenDb Structure.RefreshTokenDB

	if len(refreshToken) == 0 {
		log.Println("Refresh Token empty")
		return nil, errors.New("refresh token empty")
	}

	const matchRefreshTokenQuery string = `
    SELECT refresh_token, user_id, valid_from, valid_to
    FROM token
    WHERE refresh_token = $1
    AND user_id = $2
    AND enabled = true
    `

	err := DB.Get(&rTokenDb, matchRefreshTokenQuery, refreshToken, userId)
	if err != nil {
		log.Printf("Error matching refresh token: %v", err)
		return nil, err
	}

	return &rTokenDb, nil
}
func ExtendRefreshToken(rToken *Structure.RefreshTokenDB, validDuration int) (*Structure.RefreshTokenDB, error) {
	if rToken.RefreshToken == nil {
		log.Println("Refresh Token empty")
		return nil, errors.New("refresh token empty")
	}

	const extendRefreshTokenQuery string = `
    UPDATE token
    SET valid_to = CURRENT_TIMESTAMP + $1::interval
    WHERE refresh_token = $2
    AND user_id = $3
    AND enabled = true
    RETURNING valid_to
    `

	interval := fmt.Sprintf("%d days", validDuration)

	var newValidTo time.Time
	err := DB.QueryRow(
		extendRefreshTokenQuery,
		interval,
		rToken.RefreshToken,
		rToken.UserID,
	).Scan(&newValidTo)

	if err != nil {
		log.Printf("Error extending refresh token: %v", err)
		return nil, fmt.Errorf("failed to extend token: %w", err)
	}

	updatedToken := *rToken
	updatedToken.ValidTo = newValidTo

	return &updatedToken, nil
}
