package data_source_route

import (
	"github.com/gin-gonic/gin"
	"github.com/phuckhoa33/web-crawler/cmd/server"
	handlers "github.com/phuckhoa33/web-crawler/internal/api/handlers/data-source"
)

func ConfigureUserDataSourceRoute(server *server.Server, route *gin.RouterGroup) {
	// Import handler
	handler := handlers.NewDataSourceManagementHandler(server)

	// Configure route
	dataSourceRoute := route.Group("/data-source")
	{
		dataSourceRoute.POST("", handler.CreateDataSource)
	}
}
