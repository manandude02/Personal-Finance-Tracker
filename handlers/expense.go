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

// CreateExpense handles creating a new expense entry.
func CreateExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from JWT claims
	userID := c.GetString("user_id")
	expense.UserID, _ = primitive.ObjectIDFromHex(userID)
	expense.ID = primitive.NewObjectID()
	expense.Date = time.Now()

	collection := config.DB.Collection("expenses")
	_, err := collection.InsertOne(context.TODO(), expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense created successfully"})
}

// GetExpenses retrieves all expenses for the logged-in user.
func GetExpenses(c *gin.Context) {
	userID := c.GetString("user_id")
	objectID, _ := primitive.ObjectIDFromHex(userID)

	collection := config.DB.Collection("expenses")
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}
	defer cursor.Close(context.TODO())

	var expenses []models.Expense
	if err = cursor.All(context.TODO(), &expenses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse expenses"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

// UpdateExpense updates an existing expense entry.
func UpdateExpense(c *gin.Context) {
	expenseID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(expenseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	var expenseUpdate models.Expense
	if err := c.ShouldBindJSON(&expenseUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("expenses")
	update := bson.M{"$set": bson.M{
		"category":    expenseUpdate.Category,
		"amount":      expenseUpdate.Amount,
		"description": expenseUpdate.Description,
	}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})
}

// DeleteExpense deletes an expense entry.
func DeleteExpense(c *gin.Context) {
	expenseID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(expenseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	collection := config.DB.Collection("expenses")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}
