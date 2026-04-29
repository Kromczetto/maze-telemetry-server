package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func saveRun(c *gin.Context) {

	var run Run

	if err := c.ShouldBindJSON(&run); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	if run.Time <= 0 || run.Cells <= 0 || run.Algorithm == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid run data",
		})

		return
	}

	run.CreatedAt = time.Now()

	_, err := runCollection.InsertOne(context.Background(), run)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": "Failed to save run",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Run saved successfully",
	})
}

func getAllRuns(c *gin.Context) {

	cursor, err := runCollection.Find(context.Background(), map[string]interface{}{})

	if err != nil {
		c.JSON(500, gin.H {
			"error": "Failed to retrieve runs",
		})

		return
	}

	var runs []Run

	if err := cursor.All(context.Background(), &runs); err != nil {
		c.JSON(500, gin.H {
			"error": "Failed to parse runs",
		})
		
		return
	}

	c.JSON(200, runs)

}