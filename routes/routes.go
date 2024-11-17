package routes

import (
	"personal-finance-tracker/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Authentication routes
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)

	// Expense routes
	router.GET("/expenses", handlers.GetExpenses)
	router.POST("/expenses", handlers.CreateExpense)
	router.PUT("/expenses/:id", handlers.UpdateExpense)
	router.DELETE("/expenses/:id", handlers.DeleteExpense)

	// Income routes
	router.GET("/income", handlers.GetIncome)
	router.POST("/income", handlers.CreateIncome)
	router.PUT("/income/:id", handlers.UpdateIncome)
	router.DELETE("/income/:id", handlers.DeleteIncome)

	// Goal routes
	router.GET("/goals", handlers.GetGoals)
	router.POST("/goals", handlers.CreateGoal)
	router.PUT("/goals/:id", handlers.UpdateGoal)
	router.DELETE("/goals/:id", handlers.DeleteGoal)
}
