package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Kami0rn/Carne/controller"

	"github.com/Kami0rn/Carne/entity"
)

const PORT = "8082"


func main() {

	entity.SetupDatabase()

	r := gin.Default()



	r.Use(CORSMiddleware())

	// Customer Routes

	r.GET("/customers", controller.ListCustomers)

	r.GET("/customer/:id", controller.GetCustomer)

	r.GET("/customer/hash/:hashed_password", controller.GetCustomerByHash)

	r.POST("/customers", controller.CreateCustomer)

	r.PATCH("/customers", controller.UpdateCustomer)

	r.DELETE("/customers/:id", controller.DeleteCustomer)

	r.GET("/requests", controller.ListCustomers)

	r.GET("/request/:id", controller.GetCustomer)

	// r.GET("/customer/hash/:hashed_password", controller.GetCustomerByHash)

	r.POST("/requests", controller.CreateCustomer)

	r.PATCH("/requests", controller.UpdateCustomer)

	r.DELETE("/requests/:id", controller.DeleteCustomer)
	

	r.Run("localhost: " + PORT)

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