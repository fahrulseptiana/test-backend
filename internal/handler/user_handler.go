package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"test-backend/internal/models"
	"test-backend/internal/service"
)

// UserHandler handles HTTP requests for users.
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// GetUsers godoc
// @Summary      List users
// @Description  get users
// @Tags         users
// @Produce      json
// @Success      200  {array}   models.User
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetAll())
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  get string by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {string}  string  "not found"
// @Router       /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, ok := h.service.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary      Create user
// @Description  create a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User"
// @Success      201   {object}  models.User
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.service.Create(user)
	c.JSON(http.StatusCreated, created)
}

// UpdateUser godoc
// @Summary      Update user
// @Description  update a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "User ID"
// @Param        user  body      models.User true  "User"
// @Success      200   {object}  models.User
// @Failure      404   {string}  string      "not found"
// @Router       /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, ok := h.service.Update(id, user)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  delete a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  {string}  string  ""
// @Failure      404  {string}  string  "not found"
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if !h.service.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
