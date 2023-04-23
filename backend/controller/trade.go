package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Panadsada/523480Project/entity"

)

// POST /trades

func CreateTrade(c *gin.Context) {

	var trade entity.Trade

	if err := c.ShouldBindJSON(&trade); err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	if err := entity.DB().Create(&trade).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}

	c.JSON(http.StatusOK, gin.H{"data": trade})

}

// GET /trade/:id

func GetTrade(c *gin.Context) {

	var trade entity.Trade

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM trades WHERE id = ?", id).Scan(&trade).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": trade})

}

// GET /trades

func ListTrades(c *gin.Context) {

	var trades []entity.Trade

	if err := entity.DB().Raw("SELECT * FROM trades").Scan(&trades).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": trades})

}

// PATCH /trades

func UpdateTrade(c *gin.Context) {

	var trade entity.Trade

	if err := c.ShouldBindJSON(&trade); err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	if tx := entity.DB().Where("id = ?", trade.ID).First(&trade); tx.RowsAffected == 0 {

		   c.JSON(http.StatusBadRequest, gin.H{"error": "trade not found"})

		   return

	}



	if err := entity.DB().Save(&trade).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": trade})

}

// DELETE /trades/:id

func DeleteTrade(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM trades WHERE id = ?", id); tx.RowsAffected == 0 {

		   c.JSON(http.StatusBadRequest, gin.H{"error": "trade not found"})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": id})

}