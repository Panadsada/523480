package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/Panadsada/523480Project/entity"
)

// POST /sell
func CreateSell(c *gin.Context) {

	var sell entity.Sell
	var drug entity.Drug
	

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร sell
	if err := c.ShouldBindJSON(&sell); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา drug ด้วย id
	if tx := entity.DB().Where("id = ?", sell.DrugID).First(&drug); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drug not found"})
		return
	}
	

	// 10: สร้าง sell
	pm := entity.Sell{
		Drug:		drug,					// โยงความสัมพันธ์กับ Entity drug
		
	}

	// ขั้นตอนการ validate ที่นำมาจาก drug test
	if _, err := govalidator.ValidateStruct(pm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pm})
}

// GET /sell/:id
func GetSell(c *gin.Context) {
	var sell entity.Sell
	id := c.Param("id")
	if err := entity.DB().Preload("Drug").Raw("SELECT * FROM selles WHERE id = ?", id).Find(&sell).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sell})
}

// GET /sell
func ListSells(c *gin.Context) {
	var sell []entity.Sell
	if err := entity.DB().Preload("Drug").Raw("SELECT * FROM sell").Find(&sell).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sell})
}

// DELETE /sell/:id
func DeleteSell(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM sells WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sell not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /sells
func UpdateSell(c *gin.Context) {
	var sell entity.Sell
	if err := c.ShouldBindJSON(&sell); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", sell.ID).First(&sell); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sell not found"})
		return
	}

	if err := entity.DB().Save(&sell).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sell})
}