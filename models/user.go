package models

import (
	"chkdIn-backend-developer/config"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Salt      string
	Password  string
	CreatedAt time.Time
	IsActive  bool
}

type UserSession struct {
	ID       int64
	User     User
	Token    string
	IsActive bool
}

// CreateUser
func CreateUser(user User) error {
	db, err := config.GetDB()
	if err != nil {
		log.Println("CreateUser: Failed while connecting with the database with error: ", err)
		return err
	}
	defer db.Close()

	sqlInsert := `INSERT INTO 
		app_user (
			name,
			email,
			salt,
			password
		)
		VALUES 
			($1, $2, $3, $4)`

	_, err = db.Exec(sqlInsert, user.Name, user.Email, user.Salt, user.Password)
	if err != nil {
		log.Println("CreateUser: Failed while creating user record with an error: ", err)
		return err
	}
	return nil
}

func CreateSession(session UserSession) error {
	db, err := config.GetDB()
	if err != nil {
		log.Println("CreateUser: Failed while connecting with the database with error: ", err)
		return err
	}
	defer db.Close()

	sqlInsert := `INSERT INTO 
			user_session (
			user_id,
			token
		)
		VALUES 
			($1, $2)`

	_, err = db.Exec(sqlInsert, session.User.ID, session.Token)
	if err != nil {
		log.Println("CreateUser: Failed while creating user record with an error: ", err)
		return err
	}
	return nil
}

// GetUserEmail
func GetUserByEmail(email string) (User, error) {
	db, err := config.GetDB()
	if err != nil {
		log.Println("GetUser: Failed while connecting with the database with an error: ", err)
		return User{}, err
	}
	defer db.Close()

	var id sql.NullInt64
	var name sql.NullString
	var password sql.NullString
	var sale sql.NullString
	var created_at sql.NullTime
	var is_active sql.NullBool

	// var user User
	query := `
		SELECT
			id,
			name,
			email,
			password,
			salt,
			created_at,
			is_active
		FROM 
			app_user 
		WHERE 
			email=$1`

	err = db.QueryRow(query, email).Scan(
		&id,
		&name,
		&email,
		&password,
		&sale,
		&created_at,
		&is_active,
	)
	if err != nil {
		log.Println("GetUser: Failed while fetching user for given email id with an error: ", err, "\t email: ", email)
		return User{}, err
	}

	user := User{
		ID:        id.Int64,
		Name:      name.String,
		Email:     email,
		Password:  password.String,
		Salt:      sale.String,
		CreatedAt: created_at.Time,
		IsActive:  is_active.Bool,
	}

	return user, nil
}

func GetUserByID(id int64) (User, error) {
	db, err := config.GetDB()
	if err != nil {
		log.Println("GetUserByID: Failed while connecting with the database with an error: ", err)
		return User{}, err
	}
	defer db.Close()

	var name sql.NullString
	var email sql.NullString
	var password sql.NullString
	var sale sql.NullString
	var created_at sql.NullTime
	var is_active sql.NullBool

	// var user User
	query := `
		SELECT
			id,
			name,
			email,
			password,
			salt,
			created_at,
			is_active
		FROM 
			app_user 
		WHERE 
			id=$1`

	err = db.QueryRow(query, id).Scan(
		&id,
		&name,
		&email,
		&password,
		&sale,
		&created_at,
		&is_active,
	)
	if err != nil {
		log.Println("GetUserByID: Failed while fetching user for given id with an error: ", err, "\t id: ", id)
		return User{}, err
	}

	user := User{
		ID:        id,
		Name:      name.String,
		Email:     email.String,
		Password:  password.String,
		Salt:      sale.String,
		CreatedAt: created_at.Time,
		IsActive:  is_active.Bool,
	}

	return user, nil
}

func GetUserList() ([]User, error) {
	db, err := config.GetDB()
	if err != nil {
		log.Println("GetUserList: failed with an error: ", err)
		return []User{}, err
	}
	defer db.Close()

	query := `SELECT 
					id,
					name,
					email,
					is_active
				FROM
				app_user`

	rows, err := db.Query(query)

	if err != nil {
		log.Println("GetUserList: failed while fetching user list with error: ", err)
		return []User{}, err
	}

	defer rows.Close()

	user := []User{}

	for rows.Next() {

		var id sql.NullInt64
		var name sql.NullString
		var email sql.NullString
		var is_active sql.NullBool

		err := rows.Scan(
			&id,
			&name,
			&email,
			&is_active,
		)
		if err != nil {
			log.Println("GetUserList: failed while scanning with an error: ", err)
			continue
		}

		user = append(user, User{
			ID:       id.Int64,
			Name:     name.String,
			Email:    email.String,
			IsActive: is_active.Bool,
		})
	}

	return user, nil
}

func UpdateUserStatus(user User, status bool) error {
	fmt.Println("status: ", status)
	db, err := config.GetDB()
	if err != nil {
		log.Println("UpdateUserStatus: Failed while connecting with the database with error: ", err)
		return err
	}
	defer db.Close()

	sqlInsert := `UPDATE app_user SET is_active = $1 WHERE id = $2`

	_, err = db.Exec(sqlInsert, status, user.ID)
	if err != nil {
		log.Println("UpdateUserStatus: Failed while changing user status with an error: ", err)
		return err
	}
	fmt.Println("status: ", status)
	return nil
}

func DeleteUserSession(user User) error {
	db, err := config.GetDB()
	if err != nil {
		log.Println("DeleteUserSession: Failed while connecting with the database with error: ", err)
		return err
	}
	defer db.Close()

	sqlInsert := `DELETE FROM user_session WHERE user_id = $1`

	_, err = db.Exec(sqlInsert, user.ID)
	if err != nil {
		log.Println("DeleteUserSession: Failed while changing user status with an error: ", err)
		return err
	}

	return nil
}

func DeleteUser(user User) error {
	_ = DeleteUserSession(user)

	db, err := config.GetDB()
	if err != nil {
		log.Println("DeleteUserStatus: Failed while connecting with the database with error: ", err)
		return err
	}
	defer db.Close()

	sqlInsert := `DELETE FROM app_user WHERE id = $1`

	_, err = db.Exec(sqlInsert, user.ID)
	if err != nil {
		log.Println("DeleteUserStatus: Failed while changing user status with an error: ", err)
		return err
	}

	return nil
}

// IsSessionValid
func IsSessionValid(token string) bool {
	db, err := config.GetDB()
	if err != nil {
		log.Println("IsSessionValid: Failed while connecting with the database with an error: ", err)
		return false
	}
	defer db.Close()

	var id sql.NullInt64
	var userId sql.NullInt32
	var createdAt sql.NullTime

	query := `
		SELECT
			id,
			user_id,
			created_at
		FROM 
			user_session 
		WHERE 
			token=$1`

	err = db.QueryRow(query, token).Scan(
		&id,
		&userId,
		&createdAt,
	)
	if err != nil {
		log.Println("IsSessionValid: Failed while fetching record for given token with an error: ", err, "\t email: ", token)
		return false
	}

	if time.Since(createdAt.Time) > time.Minute*60*24*30 {
		return false
	}

	return true
}
