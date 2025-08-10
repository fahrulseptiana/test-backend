package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for products.
type Handler struct {
	service Service
}

// NewHandler creates a new Handler.
func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

// GetProducts godoc
// @Summary      List products
// @Description  get products
// @Tags         products
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   Product
// @Router       /products [get]
func (h *Handler) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetAll())
}

// GetProduct godoc
// @Summary      Get product by ID
// @Description  get product by ID
// @Tags         products
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  Product
// @Failure      404  {string}  string  "not found"
// @Router       /products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	product, ok := h.service.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct godoc
// @Summary      Create product
// @Description  create a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        product  body      Product  true  "Product"
// @Success      201   {object}  Product
// @Router       /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := h.service.Create(product)
	c.JSON(http.StatusCreated, created)
}

// UpdateProduct godoc
// @Summary      Update product
// @Description  update a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int       true  "Product ID"
// @Param        product  body      Product true  "Product"
// @Success      200   {object}  Product
// @Failure      404   {string}  string    "not found"
// @Router       /products/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, ok := h.service.Update(id, product)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteProduct godoc
// @Summary      Delete product
// @Description  delete a product by ID
// @Tags         products
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Product ID"
// @Success      204  {string}  string  ""
// @Failure      404  {string}  string  "not found"
// @Router       /products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
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
