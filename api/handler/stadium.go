package handler

import (
	pb1 "apigateway/genproto/register"
	pb "apigateway/genproto/stadium"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStadium(c *gin.Context) {
	stadium := pb.CreateStadiumRequest{}
	if err := c.ShouldBindJSON(&stadium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.UserService.CheckUserId(context.Background(), &pb1.CheckUserIdRequest{Id: stadium.UserId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.IsExist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	resp1, err := h.Stadium.CreateStadium(context.Background(), &stadium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp1)
}

func (h *Handler) UpdateStadium(c *gin.Context) {
	stadium := pb.UpdateRequest{}
	if err := c.ShouldBindJSON(&stadium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.UserService.CheckUserId(context.Background(), &pb1.CheckUserIdRequest{Id: stadium.UserId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.IsExist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	resp1, err := h.Stadium.UpdateStadium(context.Background(), &stadium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp1)
}

func (h *Handler) GetStadium(c *gin.Context) {
	stadium := pb.GetStadiumRequest{}
	if err := c.ShouldBindJSON(&stadium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Stadium.GetStadium(context.Background(), &stadium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetStadiums(c *gin.Context) {
	stadium := pb.GetStadiumsRequest{}
	if err := c.ShouldBindJSON(&stadium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Stadium.GetStadiums(context.Background(), &stadium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllStadium(c *gin.Context) {
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	// Proto request obyektini yaratish
	stadium := pb.GetAllStadiumRequest{
		Limit: int32(limit),
		Page:  int32(page),
	}

	resp, err := h.Stadium.GetAllStadium(context.Background(), &stadium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteStadium(c *gin.Context) {
	stadium := pb.DeleteStadiumRequest{}
	if err := c.ShouldBindJSON(&stadium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Stadium.DeleteStadium(context.Background(), &stadium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateOrderStadium(c *gin.Context) {
	order := pb.CreateOrderStadiumRequest{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.UserService.CheckUserId(context.Background(), &pb1.CheckUserIdRequest{Id: order.UserId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.IsExist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	resp1, err := h.Stadium.CreateOrderStadium(context.Background(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp1)
}

func (h *Handler) GetOrderStadiums(c *gin.Context) {
	order := pb.GetOrderStadiumsRequest{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Stadium.GetOrderStadiums(context.Background(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetOrderStadium(c *gin.Context) {
	order := pb.GetOrderStadiumRequest{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Stadium.GetOrderStadium(context.Background(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateOrderStadium(c *gin.Context) {
	order := pb.UpdateOrderStadiumRequest{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.UserService.CheckUserId(context.Background(), &pb1.CheckUserIdRequest{Id: order.UserId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if resp.IsExist == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	resp1, err := h.Stadium.UpdateOrderStadium(context.Background(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp1)
}

func (h *Handler) DeleteOrderStadium(c *gin.Context) {
	order := pb.DeleteOrderStadiumRequest{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Stadium.DeleteOrderStadium(context.Background(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetDeletedOrderStadiums(c *gin.Context) {
	order := pb.GetDeletedOrderStadiumsRequest{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Stadium.GetDeletedOrderStadiums(context.Background(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
