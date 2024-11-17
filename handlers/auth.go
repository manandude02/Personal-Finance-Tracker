package handlers

import (
	"context"
	"net/http"

	"personal-finance-tracker/config"
	"personal-finance-tracker/models"
	"personal-finance-tracker/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch user from MongoDB
	var user models.User
	collection := config.DB.Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"username": input.Username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Check password (replace this with password hash verification in production)
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username already exists
	collection := config.DB.Collection("users")
	count, err := collection.CountDocuments(context.TODO(), bson.M{"username": input.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Insert new user
	input.ID = primitive.NewObjectID()
	_, err = collection.InsertOne(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
