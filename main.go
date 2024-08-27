// main.go
package main

import (
	"bugsmirror/db"
	"bugsmirror/handlers"
	"bugsmirror/migrations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize GORM DB
    db, err := db.NewDB()
    if err != nil {
        panic("failed to connect database")
    }

    // Run migrations
    migrations.RunMigrations(db)

    // Initialize Gin router
    router := gin.Default()

    // Define your routes here
	router.POST("/signup", handlers.CreateUser)
	router.POST("/login", handlers.LoginUser)
    router.POST("/users/:userID/complaints", handlers.AddComplaint)
    router.GET("/users/:userID/complaints", handlers.GetUserComplaints)
    router.GET("/admin/:userID/complaints", handlers.GetAllComplaints)
    router.GET("/users/:userID/complaints/:complaintID", handlers.GetComplaint)
    router.PUT("/users/:userID/complaints/:complaintID/mark-resolved", handlers.MarkComplaintResolved)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "App is working")
	})


    // Start the server
    router.Run(":8080")
}
