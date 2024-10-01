package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/phuckhoa33/web-crawler/cmd/server"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
	"github.com/phuckhoa33/web-crawler/internal/domain/repositories"
	data_source_request "github.com/phuckhoa33/web-crawler/internal/domain/requests/data-source"
	wrapper_responses "github.com/phuckhoa33/web-crawler/internal/domain/responses"
	data_source_response "github.com/phuckhoa33/web-crawler/internal/domain/responses/data-source"
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

	// Ping the website
	client := http.Client{
		Timeout: time.Duration(h.Server.Config.App.PingTimeout) * time.Second,
	}
	resp, err := client.Get(request.Url)
	if err != nil || resp.StatusCode != http.StatusOK {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Unable to reach the website URL")
		return
	}
	defer resp.Body.Close()

	// Init a new data source
	fields := map[string]interface{}{
		"url": request.Url,
	}
	existedDataSource, repoError := h.Repo.FindOne(fields)
	if repoError != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, repoError.Error())
		return
	}

	if existedDataSource != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Data source already exists")
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

// GetDataSource godoc
// @Summary Get a data source
// @Description Get a data source
// @ID get-data-source
// @Tags data-source-management
// @Accept json
// @Produce json
// @Param dto query data_source_request.GetAllDataSourcesRequest true "Data source object"
// @Success 200 {object} data_source_response.ViewDataSourcesResponse[]
// @Failure 400 {string} string "BAD_REQUEST"
// @Router /data-source [get]
func (h *DataSourceManagementHandler) GetAllDataSource(context *gin.Context) {
	// get request
	request := new(data_source_request.GetAllDataSourcesRequest)
	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check field is empty
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Get data source
	dataSource, err := h.Repo.FindAll()
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// map data source
	response := make([]data_source_response.ViewDataSourcesResponse, 0)
	for _, data := range dataSource {
		response = append(response, data_source_response.ViewDataSourcesResponse{
			Name:     data.Name,
			Url:      data.Url,
			Industry: data.Industry,
			ID:       data.ID,
		})
	}

	wrapper_responses.Response(context, http.StatusOK, response)
}

// DeleteDataSource godoc
// @Summary Delete a data source
// @Description Delete a data source
// @ID delete-data-source
// @Tags data-source-management
// @Accept json
// @Produce json
// @Param id path string true "Data source ID"
// @Success 200
// @Failure 400 {string} string "BAD_REQUEST"
// @Router /data-source/{id} [delete]
func (h *DataSourceManagementHandler) DeleteDataSource(context *gin.Context) {
	// get request
	idStr := context.Param("id")

	// Check field is empty
	if idStr == "" {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "ID is required")
		return
	}

	// Convert id to UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Invalid ID")
		return
	}

	// Get data source
	dataSource, err := h.Repo.FindbyID(id)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	if dataSource == nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "Data source not found")
		return
	}

	// Delete data source
	if err := h.Repo.Delete(id); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
}

// DeleteAllDataSource godoc
// @Summary Delete all data sources
// @Description Delete all data sources
// @ID delete-all-data-sources
// @Tags data-source-management
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {string} string "BAD_REQUEST"
// @Router /data-source/all [delete]
func (h *DataSourceManagementHandler) DeleteAllDataSource(context *gin.Context) {
	// Delete all data source
	if err := h.Repo.DeleteAll(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
}
