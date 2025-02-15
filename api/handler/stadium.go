package handler

import (
	pb1 "apigateway/genproto/register"
	pb "apigateway/genproto/stadium"
	"context"
	"fmt"
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
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit := 10 // Default limit
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = l
		}
	}

	page := 1 // Default page
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil {
			page = p
		}
	}
	fmt.Println(limit, page)

	limitInt32 := int32(limit)
	pageInt32 := int32(page)

	req := pb.GetAllStadiumRequest{
		Limit: limitInt32,
		Page:  pageInt32,
	}

	resp, err := h.Stadium.GetAllStadium(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("error retrieving all stadium information %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, resp)
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
