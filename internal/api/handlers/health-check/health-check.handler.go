package handlers

import "github.com/gin-gonic/gin"

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) HealthCheck(context *gin.Context) {
}
