package main

import (
	"go-jwt/controllers"
	"go-jwt/database"
	_ "go-jwt/docs"
	"go-jwt/middlewares"
	"go-jwt/models"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the program.
// It initializes the database, sets up the router and starts the server.

// @title Swagger JWT API
// @version 1.0
// @description Create  Go REST API with JWT Authentication in Gin Framework
// @contact.name API Support
// @termsOfService demo.com
// @contact.url http://demo.com/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
//@Schemes http https
// @query.collection.format multi
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
 // Initialize the database
 err := database.InitDatabase()
 if err != nil {
  // Log the error and exit
  log.Fatalln("could not create database", err)
 }
 // Automigrate the User model
// AutoMigrate() automatically migrates our schema, to keep our schema upto date.
 database.GlobalDB.AutoMigrate(&models.User{})
 // Set up the router
 r := setupRouter()
 // Start the server
 r.Run(":8080")
}

// setupRouter sets up the router and adds the routes.
func setupRouter() *gin.Engine {
 // Create a new router
 r := gin.Default()
 // Add a welcome route
 r.GET("/", func(c *gin.Context) {
  c.String(200, "Welcome To This Website")
 })
 // Create a new group for the API
 api := r.Group("/api")
 {
  // Create a new group for the public routes
  public := api.Group("/public")
  {
   // Add the login route
   public.POST("/login", controllers.Login)
   // Add the signup route
   public.POST("/signup", controllers.Signup)
  }
  // Add the signup route
  protected := api.Group("/protected").Use(middlewares.Authz())
  {
   // Add the profile route
   protected.GET("/profile", controllers.Profile)   
  }
 }
 // docs route
 r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
 // Return the router
 return r
}