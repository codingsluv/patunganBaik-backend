package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/codingsluv/crowdfounding/auth"
	"github.com/codingsluv/crowdfounding/campaign"
	"github.com/codingsluv/crowdfounding/handler"
	"github.com/codingsluv/crowdfounding/helper"
	"github.com/codingsluv/crowdfounding/payment"
	"github.com/codingsluv/crowdfounding/transaction"
	"github.com/codingsluv/crowdfounding/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// import "gorm.io/driver/mysql"
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "root:codingsluv@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewSevice(userRepository)
	campaingService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaingService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")

	// * User Endpoints
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/check_email", userHandler.CheckEmail)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// * Campaign Endpoints
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)

	// * Transaction Endpoints
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransaction)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransaction)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)

	router.Run(":8080")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ? Mendapatkan token dari header authorization
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//? Melakukan validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claimToken, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//? Mendapatkan id user berdasarkan token
		userID := int(claimToken["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.ApiResponse("Unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}

}
