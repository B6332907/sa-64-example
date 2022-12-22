package controller

import (
	"net/http"

	"github.com/chanwit/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /roles
func CreateRole(c *gin.Context) {
	var role entity.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": role})
}

// GET /role/:id
func GetRole(c *gin.Context) {
	var role entity.Role
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM roles WHERE id = ?", id).Find(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// GET /role/watched/user/:id
func GetRoleWatchedByUser(c *gin.Context) {
	var role entity.Role
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM roles WHERE owner_id = ? AND title = ?", id, "Watched").Find(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

// GET /roles
func ListRoles(c *gin.Context) {
	var roles []entity.Role
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM roles").Find(&roles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// DELETE /roles/:id
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM roles WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /roles
func UpdateRole(c *gin.Context) {
	var role entity.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", role.ID).First(&role); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
		return
	}

	if err := entity.DB().Save(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}
