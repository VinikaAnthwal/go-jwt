package controllers

import (
	"go-jwt/auth"
	"go-jwt/database"
	"go-jwt/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginPayload login body
// LoginPayload is a struct that contains the fields for a user's login credentials
type LoginPayload struct {
 Email    string `json:"email" binding:"required"`
 Password string `json:"password" binding:"required"`
}

// LoginResponse token response
// LoginResponse is a struct that contains the fields for a user's login response
type LoginResponse struct {
 Token        string `json:"token"`
 RefreshToken string `json:"refreshtoken"`
}

// Signup is a function that handles user signup
// It takes in a gin context as an argument and binds the user data from the request body to a user struct
// It then hashes the user's password and creates a user record in the database
// If successful, it returns a 200 status code with a success message
// If unsuccessful, it returns a 400 or 500 status code with an error message

func Signup(c *gin.Context) {
 var user models.User
 err := c.ShouldBindJSON(&user)
 if err != nil {
  log.Println(err)
  c.JSON(400, gin.H{
   "Error": "Invalid Inputs ",
  })
  c.Abort()
  return
 }
 err = user.HashPassword(user.Password)
 if err != nil {
  log.Println(err.Error())
  c.JSON(500, gin.H{
   "Error": "Error Hashing Password",
  })
  c.Abort()
  return
 }
 err = user.CreateUserRecord()
 if err != nil {
  log.Println(err)
  c.JSON(500, gin.H{
   "Error": "Error Creating User",
  })
  c.Abort()
  return
 }
 c.JSON(200, gin.H{
  "Message": "Sucessfully Register",
 })
}

// Login is a function that handles user login
// It takes in a gin context as an argument and binds the user data from the request body to a LoginPayload struct
// It then checks if the user exists in the database and if the password is correct
// If successful, it generates a token and a refresh token and returns a 200 status code with the token and refresh token
// If unsuccessful, it returns a 401 or 500 status code with an error message

func Login(c *gin.Context) {
 var payload LoginPayload
 var user models.User
 err := c.ShouldBindJSON(&payload)
 if err != nil {
  c.JSON(400, gin.H{
   "Error": "Invalid Inputs",
  })
  c.Abort()
  return
 }
 result := database.GlobalDB.Where("email = ?", payload.Email).First(&user)
 if result.Error == gorm.ErrRecordNotFound {
  c.JSON(401, gin.H{
   "Error": "Invalid User Credentials",
  })
  c.Abort()
  return
 }
 err = user.CheckPassword(payload.Password)
 if err != nil {
  log.Println(err)
  c.JSON(401, gin.H{
   "Error": "Invalid User Credentials",
  })
  c.Abort()
  return
 }
 jwtWrapper := auth.JwtWrapper{
  SecretKey:       "verysecretkey",
  Issuer:          "AuthService",
  ExpirationMinutes: 1,
  ExpirationHours: 12,
 }
 signedToken, err := jwtWrapper.GenerateToken(user.Email)
 if err != nil {
  log.Println(err)
  c.JSON(500, gin.H{
   "Error": "Error Signing Token",
  })
  c.Abort()
  return
 }
 signedtoken, err := jwtWrapper.RefreshToken(user.Email)
 if err != nil {
  log.Println(err)
  c.JSON(500, gin.H{
   "Error": "Error Signing Token",
  })
  c.Abort()
  return
 }
 tokenResponse := LoginResponse{
  Token:        signedToken,
  RefreshToken: signedtoken,
 }
 c.JSON(200, tokenResponse)
}