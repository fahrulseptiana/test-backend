package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"test-backend/internal/user"
)

type Handler struct {
	service user.Service
	jwtKey  []byte
}

func NewHandler(s user.Service, key []byte) *Handler {
	return &Handler{service: s, jwtKey: key}
}

type Credentials struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary      Register user
// @Description  register a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      Credentials  true  "Credentials"
// @Success      201  {object} map[string]string
// @Router       /register [post]
func (h *Handler) Register(c *gin.Context) {
	var req Credentials
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.service.Create(user.User{Name: req.Name, Email: req.Email, Password: req.Password})
	token, err := h.generateToken(created)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

// Login godoc
// @Summary      Login user
// @Description  authenticate a user and return JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      Credentials  true  "Credentials"
// @Success      200  {object} map[string]string
// @Failure      401  {string} string "invalid credentials"
// @Router       /login [post]
func (h *Handler) Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, ok := h.service.Authenticate(creds.Email, creds.Password)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := h.generateToken(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) generateToken(u user.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.jwtKey)
}

func JWTMiddleware(key []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return key, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}
