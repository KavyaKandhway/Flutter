package routers

import (
	"furrble.com/backend/middlewares"
	"github.com/gin-gonic/gin"
)

//SetupRouter function will perform all route operations
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//Load the middleware in the system
	//TODO : Make the credentials filepath a environment variable
	middleware, err := middlewares.New("/Users/User/Downloads/02.05/backend-master/credentials.json", nil)
	if err != nil {
		panic(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Access Denied",
		})
	})
	//API route for version 1
	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Access Denied",
		})
	})
	//If you want to pass your route through specific middlewares
	v1.Use(middleware.TokenValidation())
	{
		//user := new(apiControllerV1.UserController)
		//v1.POST("/addnameemail", user.AddNameEmail)
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// Handle error response when a route is not defined
	r.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	return r

}
