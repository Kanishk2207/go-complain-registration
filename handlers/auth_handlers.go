// handlers/user_handler.go
package handlers

import (
	"bugsmirror/db"
	"bugsmirror/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    // Parse JSON request body into User struct
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Initialize GORM DB
    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to database"})
        return
    }

    // Create the user
    result := db.Create(&newUser)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func LoginUser(c *gin.Context) {
    var loginUser models.User
    if err := c.ShouldBindJSON(&loginUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to database"})
        return
    }


    var user models.User
    result := db.Where("email = ?", loginUser.Email).First(&user)
    if result.Error != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    if user.Password != loginUser.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    // Return user information in the response
    c.JSON(http.StatusOK, gin.H{
        "message": "login successful",
        "user": gin.H{
            "id":    user.ID,
            "email": user.Email,
			"status": user.Status,
            // Include any other user information you want to expose
        },
    })
}