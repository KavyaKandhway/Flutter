package v1

import (
    "furrble.com/backend/forms"
    "github.com/gin-gonic/gin"
    "furrble.com/backend/models"
)

// Import the userModel from the models
var userModel = new(models.User)

// UserController defines the user controller methods
type UserController struct{}

// Signup controller handles registering a user
func (u *UserController) AddNameEmail(c *gin.Context) {
    var data forms.SignupUserCommand

    // Bind the data from the request body to the SignupUserCommand Struct
    // Also check if all fields are provided
    if c.BindJSON(&data) != nil {
        // specified response
        c.JSON(406, gin.H{"message": "Provide relevant fields"})
        // abort the request
        c.Abort()
        // return nothing
        return
    }
    result, _ := userModel.GetUserByEmail(data.Email)

    // If there happens to be a result respond with a
    // descriptive mesage
    if result.Email != "" {
        c.JSON(403, gin.H{"message": "Email is already in use"})
        c.Abort()
        return
    }

    err := userModel.AddNameEmail(data)

    // Check if there was an error when saving user
    if err != nil {
        c.JSON(400, gin.H{"message": "Problem passing the data"})
        c.Abort()
        return
    }

    c.JSON(201, gin.H{"message": "User Registered in the System"})
}

