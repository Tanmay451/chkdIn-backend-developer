package controllers

import (
	"chkdIn-backend-developer/models"
	"chkdIn-backend-developer/utility"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	err := c.Request.ParseForm()

	if err != nil {
		log.Println("Register: Failed while form parsing with error: ", err)
		c.JSON(http.StatusOK, gin.H{})
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Println("name: ", name, "email: ", email, "password: ", password)

	salt := "somerandomstring" // Need to generate some random string on runtime

	passwordHash, err := utility.HashPassword(password + salt)
	if err != nil {
		log.Println("RegisterUser : Failed to hashing password with error : ", err)
		c.JSON(http.StatusOK, gin.H{"status": "failed", "message": "Failed while hashing the password"})
		return
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Salt:     salt,
		Password: passwordHash,
	}

	err = models.CreateUser(user)
	if err != nil {
		log.Println("Register: Failed while creating user with an error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": "failed", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Authenticate(c *gin.Context) {
	err := c.Request.ParseForm()

	if err != nil {
		log.Println("Authenticate: Failed while form parsing with error: ", err)
		c.JSON(http.StatusOK, gin.H{})
	}

	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.GetUserByEmail(email)
	if err != nil {
		log.Println("Authenticate: Failed while fetching user with an error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": "failed", "error": err})
	}

	salt := "somerandomstring" // Need to generate some random string on runtime

	passwordWithSalt := password + salt

	if !utility.VerifyPassword(passwordWithSalt, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		return
	}

	session := models.UserSession{
		User:  user,
		Token: utility.CreateSession(c, email),
	}

	err = models.CreateSession(session)
	if err != nil {
		log.Println("Authenticate: Failed while user session with an error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": "failed", "error": err})
	}

	fmt.Println("email: ", email, "\npassword: ", password, "\ntoken: ", session.Token, "\nuser id: ", user.ID)

	c.JSON(http.StatusOK, gin.H{})
}

func GetUserList(c *gin.Context) {

	user, err := models.GetUserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "failed", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"length": len(user), "user": user})
}

func UpdateUserStatus(c *gin.Context) {
	err := c.Request.ParseForm()

	if err != nil {
		log.Println("UpdateUserStatus: Failed while form parsing with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
	}

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		log.Println("UpdateUserStatus: Failed while parsing user id with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
		return
	}

	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		log.Println("UpdateUserStatus: Failed while parsing status with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
		return
	}

	user, err := models.GetUserByID(int64(id))
	if err != nil {
		log.Println("UpdateUserStatus: Failed while fetching user with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
		return
	}

	err = models.UpdateUserStatus(user, status)
	if err != nil {
		log.Println("UpdateUserStatus: Failed while updating user status with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteUser(c *gin.Context) {
	err := c.Request.ParseForm()

	if err != nil {
		log.Println("DeleteUser: Failed while form parsing with error: ", err)
		c.JSON(http.StatusOK, gin.H{})
	}

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		log.Println("DeleteUser: Failed while parsing user id with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
		return
	}

	user, err := models.GetUserByID(int64(id))
	if err != nil {
		log.Println("DeleteUser: Failed while fetching user with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
		return
	}

	err = models.DeleteUser(user)
	if err != nil {
		log.Println("DeleteUser: Failed while updating user status with error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
		return
	}

	fmt.Println("id: ", id)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
