package controllers

import (
	"net/http"
	"time"

	"github.com/Ryuuukin/ap-assignment1/initializers"
	"github.com/Ryuuukin/ap-assignment1/logging"
	"github.com/Ryuuukin/ap-assignment1/models"
	"github.com/gin-gonic/gin"
)

// create
func UsersCreate(c *gin.Context) {
	var userInfo struct {
		Name  string `gorm:"type:varchar(255)"`
		Email string `gorm:"type:varchar(255)"`
		Game  string `gorm:"type:varchar(255)"`
	}

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Users{
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		Game:      userInfo.Game,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Added new player!",
		"user":    user,
	})

	logging.LogUserCreation(user.Name, user.Game)
}

// read all
func UsersIndex(c *gin.Context) {
	var users []models.Users
	if err := initializers.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "All the players list:",
		"users":   users,
	})

	logging.LogUsersIndex()
}

// read one
func UsersShow(c *gin.Context) {
	id := c.Param("id")

	var user models.Users
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Found the player!",
		"user":    user,
	})

	logging.LogUsersShow(id)
}

// update
func UsersUpdate(c *gin.Context) {
	id := c.Param("id")

	var userInfo struct {
		Name  string `gorm:"type:varchar(255)"`
		Email string `gorm:"type:varchar(255)"`
		Game  string `gorm:"type:varchar(255)"`
	}

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Users
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Model(&user).Updates(models.Users{
		Name:  userInfo.Name,
		Email: userInfo.Email,
		Game:  userInfo.Game,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "updated player's info:",
		"user":    user,
	})

	logging.LogUserUpdate(user.Name, user.Game)
}

// delete
func UsersDelete(c *gin.Context) {
	id := c.Param("id")

	var user models.Users
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Delete(&models.Users{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully deleted the user!",
	})

	logging.LogUserDeletion(user.Name, user.Game)
}

// filter + sort + paginate
// FilteredUsersIndex handles filtering, sorting, and pagination based on JSON request body
func FilteredUsersIndex(c *gin.Context) {
	var (
		users []models.Users
		count int64
	)

	var requestBody struct {
		Game string `gorm:"type:varchar(255)"`
		Page int
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filterGame := requestBody.Game
	page := requestBody.Page
	sortBy := c.DefaultQuery("sort", "id")
	pageSize := 5

	if err := initializers.DB.Model(&models.Users{}).Where("game = ?", filterGame).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := initializers.DB.Model(&models.Users{}).Where("game = ?", filterGame)
	if err := query.Order(sortBy).Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"message":   "Filtered and paginated players list:",
		"users":     users,
		"total":     count,
		"page":      page,
		"page_size": pageSize,
	})

	logging.LogFilteringSortingPaginating(filterGame, sortBy, page)
}
