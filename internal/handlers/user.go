// internal/handlers/user.go
package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			PublicKey string `json:"publicKey"`
		}
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Save the public key to the database
		_, err := db.Exec("INSERT INTO users (public_key) VALUES ($1)", requestBody.PublicKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered", "publicKey": requestBody.PublicKey})
	}
}
