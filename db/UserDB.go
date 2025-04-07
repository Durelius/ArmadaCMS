package db

import (
	"ArmadaCMS/Structure"
	"ArmadaCMS/Utilities"
	"errors"
	"log"
)

func InsertUser(user Structure.User) (*Structure.User, error) {
	const insertUserSQL string = `
	  INSERT INTO user (
	  username, password, title, full_name
	  ) VALUES($1,$2,$3,$4)
	  ;`
	const getUserSQL string = `
	  SELECT id, username, password, title, full_name, created_at 
	  FROM user
	  WHERE username = $1
	  ;`

	_, err := DB.Exec(insertUserSQL, user.Username, Utilities.HashPassword(user.Password), user.Title, user.FullName)
	if err != nil {
		return nil, err
	}

	var userDB Structure.User

	err = DB.Get(&userDB, getUserSQL, user.Username)
	if err != nil {
		log.Fatal(err)
	}

	//transfer over the unhashed password to the user with the ID
	//so we can verify the user and generate tokens in the next step
	userDB.Password = user.Password

	return &userDB, nil
}

func VerifyUser(user Structure.User) (int, error) {
	const getUserString string = `
	  SELECT id, username, password, title, full_name, created_at 
	  FROM user 
	  WHERE username = $1
	  ;`
	var userDB Structure.User

	err := DB.Get(&userDB, getUserString, user.Username)

	if err != nil {
		log.Fatal(err)
	}
	if !Utilities.CheckPasswordHash(user.Password, userDB.Password) {
		return 0, errors.New("invalid username or password")
	}
	return userDB.Id, nil
}

func PopulateUserFromId(id int) (Structure.User, error) {
	emptyUser := Structure.User{}
	const getUserString string = `
SELECT id, username, password, title, full_name, created_at 
FROM user 
WHERE id = $1
	  ;`
	var userDB Structure.User
	err := DB.Get(&userDB, getUserString, id)
	if err != nil {
		log.Println(err)
		return emptyUser, nil
	}
	return userDB, nil
}
