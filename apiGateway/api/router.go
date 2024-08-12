package api

import (
	"api/api/handler"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Personal Finance Tracker API
// @version         1
// @description     Personal Finance Tracker
// @in header
// @in header
// @name Authorization
func New(h handler.Handler) *gin.Engine {
	r := gin.New()

	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		RequestHeaders: "Authorization, Origin, Content-Type",
		Methods:        "POST, GET, PUT, DELETE, OPTION",
	}))

	r.Use(traceRequest)
	budget_service := r.Group("/budget_service")
	{

		account := budget_service.Group("/v1")
		{

			account.POST("/account", h.CreateAccountHandler)
			account.GET("/account/:id", h.GetAccountHandler)
			account.GET("/accounts", h.GetAllAccountsHandler)
			account.PUT("/account/:id", h.UpdateAccountHandler)
			account.DELETE("/account/:id", h.DeleteAccountHandler)
		}

		budget := budget_service.Group("/v2")

		{
			budget.POST("/budget", h.CreateBudgetHandler)
			budget.GET("/budget/:id", h.GetBudgetHandler)
			budget.GET("/budgets", h.GetAllBudgetsHandler)
			budget.PUT("/budget/:id", h.UpdateBudgetHandler)
			budget.DELETE("/budget/:id", h.DeleteBudgetHandler)

		}
		category := budget_service.Group("/v3")

		{
			category.POST("/category", h.CreateCategoryHandler)
			category.GET("/category/:id", h.GetCategoryHandler)
			category.GET("/categories", h.GetAllCategoriesHandler)
			category.PUT("/category/:id", h.UpdateCategoryHandler)
			category.DELETE("/category/:id", h.DeleteCategoryHandler)

		}
		goal := budget_service.Group("/v4")
		{
			goal.POST("/goal", h.CreateGoalHandler)
			goal.GET("/goal/:id", h.GetGoalHandler)
			goal.GET("/goals", h.GetAllGoalsHandler)
			goal.GET("/goals/report-progress", h.GenerateGoalProgressReportHandler)
			goal.PUT("/goal/:id", h.UpdateGoalHandler)
			goal.DELETE("/goal/:id", h.DeleteGoalHandler)

		}
		transaction := budget_service.Group("/v5")
		{
			transaction.POST("/transaction", h.CreateTransactionHandler)
			transaction.GET("/transaction/:id", h.GetTransactionHandler)
			transaction.GET("/transactions", h.GetAllTransactionsHandler)
			transaction.GET("/transactions/spend", h.GetUserSpendingHandler)
			transaction.GET("/transactions/income", h.GetUserIncomeHandler)
			transaction.PUT("/transaction/:id", h.UpdateTransactionHandler)
			transaction.DELETE("/transaction/:id", h.DeleteTransactionHandler)

		}
	}
	user := r.Group("/user_service/v6/user")
	{
		user.GET("/userprofile/:id", h.UserProfile)
		user.PUT("/updateprofile", h.UpdateUserProfile)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func traceRequest(c *gin.Context) {
	beforeRequest(c)

	c.Next()

	afterRequest(c)
}

func beforeRequest(c *gin.Context) {
	startTime := time.Now()

	c.Set("start_time", startTime)

	log.Println("start time:", startTime.Format("2006-01-02 15:04:05.0000"), "path:", c.Request.URL.Path)
}

func afterRequest(c *gin.Context) {

	startTime, exists := c.Get("start_time")
	if !exists {
		startTime = time.Now()
	}

	duration := time.Since(startTime.(time.Time)).Seconds()

	log.Println("end time:", time.Now().Format("2006-01-02 15:04:05.0000"), "duration:", duration, " second", "method:", c.Request.Method)
	fmt.Println()
}
