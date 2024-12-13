package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-web-graph-app/internal/db"
	"go-web-graph-app/internal/models"
)

// GetGraphNodes retrieves nodes with indegree greater than 5 and within the specified range.
func GetGraphNodes(c *gin.Context) {
	nodes, err := db.GetFilteredNodes(5, -100, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve nodes"})
		return
	}

	c.JSON(http.StatusOK, nodes)
}