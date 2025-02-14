package api

import (
	"apigateway/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	h := handler.NewHandler()

	s := router.Group("/stadium")
	s.POST("/createstadium", h.CreateStadium)
	s.POST("/updatestadium", h.UpdateStadium)
	s.GET("/getstadium", h.GetStadium)
	s.GET("/getstadiums", h.GetStadiums)
	s.DELETE("/deletestadium", h.DeleteStadium)
	s.POST("/createorderstadium", h.CreateOrderStadium)
	s.GET("/getorderstadium", h.GetOrderStadium)
	s.GET("/getorderstadiums", h.GetOrderStadiums)
	s.PUT("/updateorderstadium", h.UpdateOrderStadium)
	s.DELETE("/deleteorderstadium", h.DeleteOrderStadium)
	s.GET("/getdeleteorderstadium", h.GetDeletedOrderStadiums)
	s.GET("/getallstadium", h.GetAllStadium)

	u := router.Group("/user")
	u.POST("/tobeanadmin", h.Tobeanadmin)
	u.PUT("/updaterole", h.UpdateRole)
	u.DELETE("/deleteregister", h.DeleteRegister)
	u.GET("/getregisters", h.GetRegisters)

	return router
}
