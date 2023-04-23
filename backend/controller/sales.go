package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/Panadsada/523480Project/entity"
)

// POST /sales
func CreateSales(c *gin.Context) {

	var sales entity.Sales
	var sell entity.Sell
	

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร sales
	if err := c.ShouldBindJSON(&sales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา sell ด้วย id
	if tx := entity.DB().Where("id = ?", sales.SellID).First(&sell); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sell not found"})
		return
	}
	

	// 10: สร้าง sales
	pm := entity.Sales{
		Sell:		sell,					// โยงความสัมพันธ์กับ Entity sell
		
	}

	// ขั้นตอนการ validate ที่นำมาจาก sell test
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

// GET /sales/:id
func GetSales(c *gin.Context) {
	var sales entity.Sales
	id := c.Param("id")
	if err := entity.DB().Preload("Sell").Raw("SELECT * FROM sales WHERE id = ?", id).Find(&sales).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sales})
}

// GET /sales
func ListSales(c *gin.Context) {
	var sales []entity.Sales
	if err := entity.DB().Preload("Sell").Raw("SELECT * FROM sales").Find(&sales).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sales})
}

// DELETE /sales/:id
func DeleteSales(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM sales WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sales not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /sales
func UpdateSales(c *gin.Context) {
	var sales entity.Sales
	if err := c.ShouldBindJSON(&sales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", sales.ID).First(&sales); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sales not found"})
		return
	}

	if err := entity.DB().Save(&sales).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sales})
}