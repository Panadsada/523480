package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"github.com/Panadsada/523480Project/entity"
)

// POST /drug
func CreateDrug(c *gin.Context) {

	var drug entity.Drug
	var unit entity.Unit

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร drug
	if err := c.ShouldBindJSON(&drug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา unit ด้วย id
	if tx := entity.DB().Where("id = ?", drug.UnitID).First(&unit); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unit not found"})
		return
	}

	// 12: สร้าง drug
	dt := entity.Drug{
		Unit:        unit,  // โยงความสัมพันธ์กับ Entity unit
		Code:         drug.Code,
		Name:         drug.Name,	
		Amount: 	  drug.Amount,
		Price:     	drug.Price,
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

// GET /drug/:id
func GetDrug(c *gin.Context) {
	var drug entity.Drug
	id := c.Param("id")
	if err := entity.DB().Preload("Unit").Raw("SELECT * FROM drugs WHERE id = ?", id).Find(&drug).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": drug})
}


// GET /drug /s
func ListDrugs(c *gin.Context) {
	var drug []entity.Drug
	if err := entity.DB().Preload("Unit").Raw("SELECT * FROM drugs").Find(&drug).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": drug})
}

// DELETE /drug/:id
func DeleteDrug(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM drugs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drug not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}


func UpdateDrug(c *gin.Context) {

	var drug entity.Drug
	var unit entity.Unit

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร drug
	if err := c.ShouldBindJSON(&drug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 11: ค้นหา unit ด้วย id
	if tx := entity.DB().Where("id = ?", drug.UnitID).First(&unit); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unit not found"})
		return
	}

	// 12: สร้าง drug
	update := entity.Drug{
		Unit:        unit,  // โยงความสัมพันธ์กับ Entity unit
		Code:         drug.Code,
		Name:         drug.Name,	
		Amount: 	  drug.Amount,
		Price:     drug.Price,
	}


	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	if err := entity.DB().Where("id = ?", drug.ID).Updates(&update).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": update})
	
}