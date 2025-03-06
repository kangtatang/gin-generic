package handlers

import (
	"fmt"
	"net/http"

	"go-gin-generic/models"
	"go-gin-generic/repositories"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repo *repositories.GenericRepository[models.User]
}

// GetAllUsers handler untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID handler untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var uid uint
	fmt.Sscanf(id, "%d", &uid) // Konversi string ke uint

	user, err := h.Repo.GetByID(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser handler untuk membuat user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
