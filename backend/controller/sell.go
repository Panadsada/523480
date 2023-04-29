package controller

import (
	"net/http"

	"github.com/Panadsada/523480Project/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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
		Quantity:   sell.Quantity,
		Cost:		sell.Cost,	
		Payment:	sell.Payment,
		Status:		sell.Status,
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

// GET /sell/:id
func GetSell(c *gin.Context) {
	var sell entity.Sell
	id := c.Param("id")
	if err := entity.DB().Preload("Drug").Preload("Types").Raw("SELECT * FROM selles WHERE id = ?", id).Find(&sell).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sell})
}

// GET /sell
func ListSells(c *gin.Context) {
	var sell []entity.Sell
	if err := entity.DB().Preload("Drug").Preload("Types").Raw("SELECT * FROM sell").Find(&sell).Error; err != nil {
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

func UpdateSell(c *gin.Context) {

	var drug entity.Drug
	var sell entity.Sell

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร drug
	if err := c.ShouldBindJSON(&drug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา drug ด้วย id
	if tx := entity.DB().Where("id = ?", sell.DrugID).First(&drug); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drug not found"})
		return
	}

	// 12: สร้าง sell
	update := entity.Sell{
		Drug:		drug,					// โยงความสัมพันธ์กับ Entity drug
		Quantity:   sell.Quantity,
		Cost:		sell.Cost,	
		Payment:	sell.Payment,
		Status:		sell.Status,
	}

	// ขั้นตอนการ validate ที่นำมาจาก sell test
	if _, err := govalidator.ValidateStruct(update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	if err := entity.DB().Where("id = ?", sell.ID).Updates(&update).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": update})
	
}