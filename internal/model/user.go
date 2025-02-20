package model

import (
	db "erp-api/database"
	"errors"
)

type User struct {
	UserID int64				`json:"user_id"`
	Username string			`json:"username"`
	Password string			`json:"password"`
	Role string					`json:"role"`
}

type UserRequest struct {
	Username string			`json:"username"`
	Password string			`json:"password"`
}


func GetUser(userReq UserRequest) (*User, error){
	userQuery := 
	`
	SELECT *
	FROM users
	WHERE username=$1
	`

	row := db.DB.QueryRow(userQuery, userReq.Username)

	var user User
	err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.Role)
	if err != nil{
		return nil, err
	}

	return &user, nil
}

func VerifyUser(userReq UserRequest, user *User) (error){
	if userReq.Password != user.Password {
		return errors.New("password does not match")
	}
	return nil
}