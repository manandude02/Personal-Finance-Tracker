package handlers

import (
	"context"
	"net/http"

	"personal-finance-tracker/config"
	"personal-finance-tracker/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

// CreateGoal handles creating a new financial goal.
func CreateGoal(c *gin.Context) {
	var goal models.Goal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from JWT claims
	userID := c.GetString("user_id")
	goal.UserID, _ = primitive.ObjectIDFromHex(userID)
	goal.ID = primitive.NewObjectID()

	collection := config.DB.Collection("goals")
	_, err := collection.InsertOne(context.TODO(), goal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal created successfully"})
}

// GetGoals retrieves all goals for the logged-in user.
func GetGoals(c *gin.Context) {
	userID := c.GetString("user_id")
	objectID, _ := primitive.ObjectIDFromHex(userID)

	collection := config.DB.Collection("goals")
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch goals"})
		return
	}
	defer cursor.Close(context.TODO())

	var goals []models.Goal
	if err = cursor.All(context.TODO(), &goals); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse goals"})
		return
	}

	c.JSON(http.StatusOK, goals)
}

// UpdateGoal updates an existing financial goal.
func UpdateGoal(c *gin.Context) {
	goalID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(goalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid goal ID"})
		return
	}

	var goalUpdate models.Goal
	if err := c.ShouldBindJSON(&goalUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("goals")
	update := bson.M{"$set": bson.M{
		"target":      goalUpdate.Target,
		"description": goalUpdate.Description,
		"completed":   goalUpdate.Completed,
	}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal updated successfully"})
}

// DeleteGoal deletes a financial goal.
func DeleteGoal(c *gin.Context) {
	goalID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(goalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid goal ID"})
		return
	}

	collection := config.DB.Collection("goals")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal deleted successfully"})
}
