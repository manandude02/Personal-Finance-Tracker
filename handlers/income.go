package handlers

import (
	"context"
	"net/http"
	"time"

	"personal-finance-tracker/config"
	"personal-finance-tracker/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

// CreateIncome handles creating a new income entry.
func CreateIncome(c *gin.Context) {
	var income models.Income
	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from JWT claims
	userID := c.GetString("user_id")
	income.UserID, _ = primitive.ObjectIDFromHex(userID)
	income.ID = primitive.NewObjectID()
	income.Date = time.Now()

	collection := config.DB.Collection("incomes")
	_, err := collection.InsertOne(context.TODO(), income)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Income created successfully"})
}

// GetIncomes retrieves all incomes for the logged-in user.
func GetIncome(c *gin.Context) {
	userID := c.GetString("user_id")
	objectID, _ := primitive.ObjectIDFromHex(userID)

	collection := config.DB.Collection("incomes")
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch incomes"})
		return
	}
	defer cursor.Close(context.TODO())

	var incomes []models.Income
	if err = cursor.All(context.TODO(), &incomes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse incomes"})
		return
	}

	c.JSON(http.StatusOK, incomes)
}

// UpdateIncome updates an existing income entry.
func UpdateIncome(c *gin.Context) {
	incomeID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(incomeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid income ID"})
		return
	}

	var incomeUpdate models.Income
	if err := c.ShouldBindJSON(&incomeUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("incomes")
	update := bson.M{"$set": bson.M{
		"source":      incomeUpdate.Source,
		"amount":      incomeUpdate.Amount,
		"description": incomeUpdate.Description,
	}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Income updated successfully"})
}

// DeleteIncome deletes an income entry.
func DeleteIncome(c *gin.Context) {
	incomeID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(incomeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid income ID"})
		return
	}

	collection := config.DB.Collection("incomes")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Income deleted successfully"})
}
