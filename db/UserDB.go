package db

import (
	"ArmadaCMS/Structure"
	"ArmadaCMS/Utilities"
	"errors"
)

func InsertUser(user Structure.User) (*Structure.User, error) {
	// Hash the password before inserting
	user.Password = Utilities.HashPassword(user.Password)

	// Insert user into the database
	if err := DB.Create(&user).Error; err != nil {
		return nil, err
	}

	// Fetch the user to return with the auto-generated ID
	var userDB Structure.User
	if err := DB.Where("username = ?", user.Username).First(&userDB).Error; err != nil {
		return nil, err
	}

	// Return the user with the password (hashed)
	return &userDB, nil
}

func VerifyUser(user Structure.User) (int, error) {
	var userDB Structure.User

	if err := DB.Where("username = ?", user.Username).First(&userDB).Error; err != nil {
		return 0, errors.New("invalid username or password")
	}

	// Check if the password matches
	if !Utilities.CheckPasswordHash(user.Password, userDB.Password) {
		return 0, errors.New("invalid username or password")
	}

	return userDB.Id, nil
}

func PopulateUserFromId(id int) (Structure.User, error) {
	var userDB Structure.User

	if err := DB.Where("id = ?", id).First(&userDB).Error; err != nil {
		return Structure.User{}, err
	}

	return userDB, nil
}
