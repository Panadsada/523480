package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"github.com/Panadsada/523480Project/entity"
)

// POST /trade
func CreateTrade(c *gin.Context) {

	var admin entity.Admin
	var trade entity.Trade
	var user entity.User
	var drug entity.Drug
	var status entity.Status

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร trade
	if err := c.ShouldBindJSON(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	// 8: ค้นหา drug ด้วย id
	if tx := entity.DB().Where("id = ?", trade.DrugID).First(&drug); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drug not found"})
		return
	}
	
	// 9: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", trade.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	// 10: ค้นหา admin ด้วย id
	if tx := entity.DB().Where("id = ?", trade.AdminID).First(&admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin not found"})
		return
	}

	
	// 11: ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", trade.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 12: สร้าง trade
	dt := entity.Trade{
		Admin:		admin, // โยงความสัมพันธ์กับ Entity admin
		User:		user,  // โยงความสัมพันธ์กับ Entity user
		QUANTITY:	trade.QUANTITY,	
		COST:		trade.COST,
		Total:		trade.Total,
		Evidence:	trade.Evidence,
		TradeTime:	trade.TradeTime,
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(dt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&dt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dt})
}

// GET /trade/:id
func GetTrade(c *gin.Context) {
	var trade entity.Trade
	id := c.Param("id")
	if err := entity.DB().Preload("Admin").Preload("User").Preload("Drug").Preload("Status").Raw("SELECT * FROM trades WHERE id = ?", id).Find(&trade).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": trade})
}


// GET /trade /s
func ListTrades(c *gin.Context) {
	var trade []entity.Trade
	if err := entity.DB().Preload("Admin").Preload("User").Preload("Drug").Preload("Status").Raw("SELECT * FROM trades").Find(&trade).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": trade})
}

// DELETE /trade/:id
func DeleteTrade(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM trades WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "trade not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}


func UpdateTrade(c *gin.Context) {

	var admin entity.Admin
	var trade entity.Trade
	var user entity.User
	var drug entity.Drug
	var status entity.Status

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร trade
	if err := c.ShouldBindJSON(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 8: ค้นหา drug ด้วย id
	if tx := entity.DB().Where("id = ?", trade.DrugID).First(&drug); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drug not found"})
		return
	}

	// 9: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", trade.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	// 10: ค้นหา admin ด้วย id
	if tx := entity.DB().Where("id = ?", trade.AdminID).First(&admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin not found"})
		return
	}

	
	// 11: ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", trade.UserID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 12: สร้าง trade
	update := entity.Trade{
		Admin:		admin, // โยงความสัมพันธ์กับ Entity admin
		User:		user,  // โยงความสัมพันธ์กับ Entity user
		QUANTITY:	trade.QUANTITY,	
		COST:		trade.COST,
		Total:		trade.Total,
		Evidence:	trade.Evidence,
		TradeTime:	trade.TradeTime,
	}


	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	if err := entity.DB().Where("id = ?", trade.ID).Updates(&update).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": update})
	
}