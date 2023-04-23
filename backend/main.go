package main

import (
	"github.com/Panadsada/523480Project/controller"
	"github.com/Panadsada/523480Project/entity"
	"github.com/Panadsada/523480Project/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			// Trade Routes
			protected.GET("/trades", controller.ListTrades)
			protected.GET("/trade/:id", controller.GetTrade)
			protected.PATCH("/trades", controller.UpdateTrade)
			protected.DELETE("/trades/:id", controller.DeleteTrade)

			// Drug Routes
			protected.GET("/drugs", controller.ListDrugs)
			protected.GET("/drug/:id", controller.GetDrug)
			protected.POST("/drugs", controller.CreateDrug)
			protected.PATCH("/drugs", controller.UpdateDrug)
			protected.DELETE("/drugs/:id", controller.DeleteDrug)

			// Sell Routes
			protected.GET("/sells", controller.ListSells)
			protected.GET("/sell/:id", controller.GetSell)
			protected.POST("/sells", controller.CreateSell)
			protected.PATCH("/sells", controller.UpdateSell)
			protected.DELETE("/sells/:id", controller.DeleteSell)

			// Sales Routes
			protected.GET("/sales", controller.ListSales)
			protected.GET("/sales/:id", controller.GetSales)
			protected.POST("/sales", controller.CreateSales)
			protected.PATCH("/sales", controller.UpdateSales)
			protected.DELETE("/sales/:id", controller.DeleteSales)

			
		}
	}

	//ไว้สมัคร user doctor
	r.POST("/users", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
