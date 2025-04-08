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

	refreshTokenDB := Structure.RefreshTokenDB{
		RefreshToken: refreshToken,
		UserID:       userId,
		ValidFrom:    time.Now(),
		ValidTo:      time.Now().Add(7 * 24 * time.Hour), // Valid for 7 days
		Enabled:      true,
	}

	if err := DB.Create(&refreshTokenDB).Error; err != nil {
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

	// Query to match the refresh token
	err := DB.Where("refresh_token = ? AND user_id = ? AND enabled = true", refreshToken, userId).First(&rTokenDb).Error
	if err != nil {
		log.Printf("Error matching refresh token: %v", err)
		return nil, err
	}

	return &rTokenDb, nil
}

func ExtendRefreshToken(rToken *Structure.RefreshTokenDB, validDuration int) (*Structure.RefreshTokenDB, error) {
	if len(rToken.RefreshToken) == 0 {
		log.Println("Refresh Token empty")
		return nil, errors.New("refresh token empty")
	}

	newValidTo := time.Now().Add(time.Duration(validDuration) * 24 * time.Hour)

	err := DB.Model(&Structure.RefreshTokenDB{}).
		Where("refresh_token = ? AND user_id = ? AND enabled = true", rToken.RefreshToken, rToken.UserID).
		Update("valid_to", newValidTo).Error

	if err != nil {
		log.Printf("Error extending refresh token: %v", err)
		return nil, fmt.Errorf("failed to extend token: %w", err)
	}

	rToken.ValidTo = newValidTo
	return rToken, nil
}
