package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuckhoa33/web-crawler/cmd/server"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
	"github.com/phuckhoa33/web-crawler/internal/domain/repositories"
	data_source_request "github.com/phuckhoa33/web-crawler/internal/domain/requests/data-source"
	wrapper_responses "github.com/phuckhoa33/web-crawler/internal/domain/responses"
)

type DataSourceManagementHandler struct {
	Server *server.Server
	Repo   *repositories.DataSourceRepository
}

func NewDataSourceManagementHandler(server *server.Server) *DataSourceManagementHandler {
	return &DataSourceManagementHandler{
		Server: server,
		Repo:   repositories.NewDataSourceRepository(server.DB),
	}
}

// CreateDataSource godoc
// @Summary Create a data source
// @Description Create a data source
// @ID create-data-source
// @Tags data-source-management
// @Accept json
// @Produce json
// @Param data-source body data_source_request.CreateDataSourceRequest true "Data source object"
// @Success 200
// @Failure 400 {string} string "BAD_REQUEST"
// @Router /data-source [post]
func (h *DataSourceManagementHandler) CreateDataSource(context *gin.Context) {
	// get request
	request := new(data_source_request.CreateDataSourceRequest)
	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check field is empty
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Init a new data source
	_, err := h.Repo.FindOne("Name", request.Name)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Create data source
	dataSource := models.DataSource{
		Name:     request.Name,
		Url:      request.Url,
		Industry: request.Industry,
	}

	if err := h.Repo.Create(&dataSource); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
}
