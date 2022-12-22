package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/chanwit/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /watch_videos
func CreatePatiend(c *gin.Context) {

	var patiend entity.Patiend
	var gender entity.Gender
	var prefix entity.Prefix
	var policing entity.Policing

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร patiend
	if err := c.ShouldBindJSON(&patiend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา prefix ด้วย id
	if tx := entity.DB().Where("id = ?", patiend.Prefix_ID).First(&prefix); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prefix not found"})
		return
	}

	// 10: ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", patiend.Gender_ID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 11: ค้นหา policing ด้วย id
	if tx := entity.DB().Where("id = ?", patiend.Policing_ID).First(&policing); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "policing not found"})
		return
	}
	// 12: สร้าง WatchVideo
	wv := entity.Patiend{
		Prefix:        prefix,                // โยงความสัมพันธ์กับ Entity prefix
		Gender:        gender,                // โยงความสัมพันธ์กับ Entity gender
		Policing:      policing,              // โยงความสัมพันธ์กับ Entity policing
		Name:          patiend.Name,          // ตั้งค่าฟิลด์ Name
		Age:           patiend.Age,           // ตั้งค่าฟิลด์ Age
		Phone:         patiend.Phone,         // ตั้งค่าฟิลด์ Phone
		Date_of_Birth: patiend.Date_of_Birth, // ตั้งค่าฟิลด์ Email
		Address:       patiend.Address,       // ตั้งค่าฟิลด์ Password
		ID_Card:       patiend.ID_Card,
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
func GetPatiend(c *gin.Context) {
	var patiend entity.Patiend
	id := c.Param("id")
	if err := entity.DB().Preload("Prefix").Preload("Gender").Preload("Policing").Raw("SELECT * FROM patiends WHERE id = ?", id).Find(&patiend).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patiend})
}

// GET /watch_videos
func ListPatiends(c *gin.Context) {
	var patiend []entity.Patiend
	if err := entity.DB().Preload("Prefix").Preload("Gender").Preload("Policing").Raw("SELECT * FROM patiends").Find(&patiend).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patiend})
}

// DELETE /watch_videos/:id
func DeletePatiend(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM patiends WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patiend not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdatePatiend(c *gin.Context) {
	var patiend entity.Patiend
	if err := c.ShouldBindJSON(&patiend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", patiend.ID).First(&patiend); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patiend not found"})
		return
	}

	if err := entity.DB().Save(&patiend).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patiend})
}
