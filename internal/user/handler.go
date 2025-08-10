package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for users.
type Handler struct {
	service Service
}

// NewHandler creates a new Handler.
func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

// GetUsers godoc
// @Summary      List users
// @Description  get users
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   User
// @Router       /users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users := h.service.GetAll()
	for i := range users {
		users[i].Password = ""
	}
	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  get string by ID
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  User
// @Failure      404  {string}  string  "not found"
// @Router       /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
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
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary      Create user
// @Description  create a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user  body      User  true  "User"
// @Success      201   {object}  User
// @Router       /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.service.Create(user)
	created.Password = ""
	c.JSON(http.StatusCreated, created)
}

// UpdateUser godoc
// @Summary      Update user
// @Description  update a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int       true  "User ID"
// @Param        user  body      User true  "User"
// @Success      200   {object}  User
// @Failure      404   {string}  string    "not found"
// @Router       /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, ok := h.service.Update(id, user)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	updated.Password = ""
	c.JSON(http.StatusOK, updated)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  delete a user by ID
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      204  {string}  string  ""
// @Failure      404  {string}  string  "not found"
// @Router       /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
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
