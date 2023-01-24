package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary About
// @Accept json
// @Produce json
// @Success 200 {object} response.AboutResponse
// @Failure 400,404,500 {object} response.ErrorResponse
// @Router /monitoring/about [get]
func (h *Handler) about(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this auth service api"})
}

// @Summary Healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} response.StatusResponse
// @Failure 400,404,500 {object} response.ErrorResponse
// @Router /monitoring/healthcheck [get]
func (h *Handler) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
