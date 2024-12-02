package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang-application/internal/helper"
	models "golang-application/internal/model"
	"golang-application/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
	wg          sync.WaitGroup
}

// NewUserHandler initializes a new UserHandler with the provided service.
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		wg:          sync.WaitGroup{},
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userService.CreateUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	type userdata struct {
		ID    int    `json:"id"`
		Name  string `json:"username"`
		Email string `json:"email"`
		Phone int    `json:"phone"`
	}

	jsonReq, err := json.Marshal(userdata{
		ID:    int(resp.Id),
		Name:  req.Name,
		Email: req.Email,
		Phone: int(req.Phone),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := strconv.Itoa(int(resp.Id))
	err = helper.Rdb.Set(userId, jsonReq, 15*time.Minute).Err()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": resp.Message,
	})
}

// UpdateUser handles updating an existing user.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = uint(id)
	resp, err := h.userService.UpdateUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	resp, err := h.userService.DeleteUser(context.Background(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
	})
}

func (h *UserHandler) ListUser(c *gin.Context) {
	resp, err := h.userService.ListUser(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	type userdata struct {
		ID    int    `json:"id"`
		Name  string `json:"username"`
		Email string `json:"email"`
		Phone int    `json:"phone"`
	}
	user, err := helper.Rdb.Get(idStr).Result()
	if err == nil {
		var userdata userdata
		err = json.Unmarshal([]byte(user), &userdata)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Println("User data fetch from Redis")

		c.JSON(http.StatusOK, gin.H{
			"userdata":  userdata,
			"fetch_from": "Redis",
		})
		return
	}

	resp, err := h.userService.GetUser(context.Background(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userData := userdata{
		ID:    int(resp.ID),
		Name:  resp.Name,
		Email: resp.Email,
		Phone: int(resp.Phone),
	}
	jsonReq, err := json.Marshal(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userId := strconv.Itoa(int(resp.ID))
	err = helper.Rdb.Set(userId, jsonReq, 15*time.Minute).Err()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println("User data fetch from Database")
	c.JSON(http.StatusOK, gin.H{
		"userdata": userData,
		"fetch_from": "Database",
	})
}
func (h *UserHandler) ListUserNames(c *gin.Context) {
	var req struct {
		Method   int `json:"method"`
		WaitTime int `json:"wait_time"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.userService.ListUserNames(context.Background(), req.Method, req.WaitTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"names": resp})
}

// func (h *UserHandler) ListUserNames2(c *gin.Context) {
// 	var req struct {
// 		Method   int `json:"method"`
// 		WaitTime int `json:"wait_time"`
// 	}
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	resp, err := h.userService.ListUserNames(ctx, req.Method, req.WaitTime)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"names": resp})
// }
