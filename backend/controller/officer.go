package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/chanwit/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /watch_videos
func CreateOfficer(c *gin.Context) {

	var officer entity.Officer
	var gender entity.Gender
	var prefix entity.Prefix
	var role entity.Role

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร officer
	if err := c.ShouldBindJSON(&officer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา prefix ด้วย id
	if tx := entity.DB().Where("id = ?", officer.Prefix_ID).First(&prefix); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prefix not found"})
		return
	}

	// 10: ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", officer.Gender_ID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 11: ค้นหา role ด้วย id
	if tx := entity.DB().Where("id = ?", officer.Role_ID).First(&role); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
		return
	}
	// 12: สร้าง WatchVideo
	wv := entity.Officer{
		Prefix:   prefix,           // โยงความสัมพันธ์กับ Entity prefix
		Gender:   gender,           // โยงความสัมพันธ์กับ Entity gender
		Role:     role,             // โยงความสัมพันธ์กับ Entity role
		Name:     officer.Name,     // ตั้งค่าฟิลด์ Name
		Age:      officer.Age,      // ตั้งค่าฟิลด์ Age
		Phone:    officer.Phone,    // ตั้งค่าฟิลด์ Phone
		Email:    officer.Email,    // ตั้งค่าฟิลด์ Email
		Password: officer.Password, // ตั้งค่าฟิลด์ Password
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(wv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /watchvideo/:id
func GetOfficer(c *gin.Context) {
	var officer entity.Officer
	id := c.Param("id")
	if err := entity.DB().Preload("Prefix").Preload("Gender").Preload("Role").Raw("SELECT * FROM officers WHERE id = ?", id).Find(&officer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": officer})
}

// GET /watch_videos
func ListOfficers(c *gin.Context) {
	var officer []entity.Officer
	if err := entity.DB().Preload("Prefix").Preload("Gender").Preload("Role").Raw("SELECT * FROM officers").Find(&officer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officer})
}

// DELETE /watch_videos/:id
func DeleteOfficer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM officers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdateOfficer(c *gin.Context) {
	var officer entity.Officer
	if err := c.ShouldBindJSON(&officer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", officer.ID).First(&officer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "officer not found"})
		return
	}

	if err := entity.DB().Save(&officer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": officer})
}
