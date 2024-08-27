// handlers/user_handler.go
package handlers

import (
	"bugsmirror/db"
	"bugsmirror/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserComplaints(c *gin.Context) {
    userID := c.Param("userID") // Assuming userID is passed in the URL parameter

    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to database"})
        return
    }

    var user models.User
    result := db.Preload("Complaints").First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, user.Complaints)
}

func AddComplaint(c *gin.Context) {
    userID := c.Param("userID") // Assuming userID is passed in the URL parameter

    var newComplaint models.Complaint
    if err := c.ShouldBindJSON(&newComplaint); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to database"})
        return
    }

    // Check if the user exists
    var user models.User
    result := db.First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    // Associate the complaint with the user
    newComplaint.UserID = user.ID

    // Create the complaint
    result = db.Create(&newComplaint)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create complaint"})
        return
    }

    // Return the newly added complaint in the response
    c.JSON(http.StatusCreated, gin.H{"message": "complaint added successfully", "complaint": newComplaint})
}

func GetAllComplaints(c *gin.Context) {
    userID := c.Param("userID") // Assuming userID is passed in the URL parameter

    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to database"})
        return
    }

    var user models.User
    result := db.First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    if user.Status != "admin" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
        return
    }

    var complaints []models.Complaint
    result = db.Find(&complaints)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch complaints"})
        return
    }

    c.JSON(http.StatusOK, complaints)
}

func GetComplaint(c *gin.Context) {
    userID := c.Param("userID")
    complaintID := c.Param("complaintID")

    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to database"})
        return
    }

    var user models.User
    result := db.First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    // Check if the user is authorized to access the complaint
    var complaint models.Complaint
    result = db.Where("id = ? AND user_id = ?", complaintID, userID).First(&complaint)
    if result.Error != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
        return
    }

    c.JSON(http.StatusOK, complaint)
}

func MarkComplaintResolved(c *gin.Context) {
    userID := c.Param("userID") // Assuming userID is passed in the URL parameter
    complaintID := c.Param("complaintID") // Assuming complaintID is passed in the URL parameter

    db, err := db.NewDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to database"})
        return
    }

    // Check if the user is an admin
    var user models.User
    result := db.First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    if user.Status != "admin" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "only admin can mark complaint as resolved"})
        return
    }

    // Mark complaint as resolved
    var complaint models.Complaint
    result = db.Where("id = ?", complaintID).First(&complaint)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "complaint not found"})
        return
    }

    complaint.Resolved = true
    db.Save(&complaint)

    c.JSON(http.StatusOK, gin.H{"message": "complaint marked as resolved"})
}