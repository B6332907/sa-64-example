package main

import (
	"github.com/B6332907/Project_SE/controller"
	"github.com/B6332907/sa-64-example/entity"
	"github.com/B6332907/sa-64-example/middlewares"
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
			// Gender Routes
			protected.GET("/genders", controller.ListGenders)
			protected.GET("/gender/:id", controller.GetGender)
			protected.PATCH("/genders", controller.UpdateGender)
			protected.DELETE("/genders/:id", controller.DeleteGender)

			// Officer Routes
			protected.GET("/officers", controller.ListOfficers)
			protected.GET("/officer/:id", controller.GetOfficer)
			protected.POST("/officers", controller.CreateOfficer)
			protected.PATCH("/officers", controller.UpdateOfficer)
			protected.DELETE("/officers/:id", controller.DeleteOfficer)

			// Prefix Routes
			protected.GET("/prefixs", controller.ListPrefixs)
			protected.GET("/prefix/:id", controller.GetPrefix)
			protected.POST("/prefixs", controller.CreatePrefix)
			protected.PATCH("/prefixs", controller.UpdatePrefix)
			protected.DELETE("/prefixs/:id", controller.DeletePrefix)

			// Role Routes
			protected.GET("/roles", controller.ListRoles)
			protected.GET("/role/:id", controller.GetRole)
			protected.POST("/roles", controller.CreateRole)
			protected.PATCH("/roles", controller.UpdateRole)
			protected.DELETE("/roles/:id", controller.DeleteRole)

			// Patiend Routes
			protected.GET("/patiends", controller.ListPatiends)
			protected.GET("/patiend/:id", controller.GetPatiend)
			protected.POST("/patiends", controller.CreatePatiend)
			protected.PATCH("/patiends", controller.UpdatePatiend)
			protected.DELETE("/patiendrs/:id", controller.DeletePatiend)

		}
	}

	// Gender Routes
	//r.POST("/genders", controller.CreateGender)

	// Authentication Routes
	//r.POST("/login", controller.Login)

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
