package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}

// HealthCheck godoc
// @Summary Health check
// @Description Check the health of the server
// @Tags health-check
// @Accept json
func (h *HealthCheckHandler) HealthCheck(context *gin.Context) {
	fmt.Println("Health check")
}
