package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuckhoa33/web-crawler/cmd/server"
	data_source_request "github.com/phuckhoa33/web-crawler/internal/domain/requests/data-source"
	wrapper_responses "github.com/phuckhoa33/web-crawler/internal/domain/responses"
	data_source_management_service "github.com/phuckhoa33/web-crawler/internal/domain/services/data-source"
)

type DataSourceManagementHandler struct {
	Server  *server.Server
	Service *data_source_management_service.DataSourceManagementService
}

func NewDataSourceManagementHandler(server *server.Server) *DataSourceManagementHandler {
	return &DataSourceManagementHandler{
		Server:  server,
		Service: data_source_management_service.NewDataSourceManagementService(server),
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

	// Create data source
	if err := h.Service.Create(*request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	wrapper_responses.Response(context, http.StatusOK, nil)
}

// GetDataSource godoc
// @Summary Get a data source
// @Description Get a data source
// @ID get-data-source
// @Tags data-source-management
// @Accept json
// @Produce json
// @Param id path string true "Data source ID"
// @Success 200 {object} data_source_response.ViewDataSourcesResponse
// @Failure 400 {string} string "BAD_REQUEST"
// @Router /data-source/{id} [get]
func (h *DataSourceManagementHandler) GetDataSource(context *gin.Context) {
	// get request
	id := context.Param("id")

	// Check field is empty
	if id == "" {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "ID is required")
		return
	}

	// Get data source
	dataSource, err := h.Service.GetOne(id)
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	wrapper_responses.Response(context, http.StatusOK, dataSource)
}

// GetDataSource godoc
// @Summary Get a data source
// @Description Get a data source
// @ID get-data-source
// @Tags data-source-management
// @Accept json
// @Produce json
// @Param dto query data_source_request.GetAllDataSourcesRequest true "Data source object"
// @Success 200 {array} data_source_response.ViewDataSourcesResponse
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

	// Get all data sources
	dataSources, err := h.Service.GetAll()
	if err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	wrapper_responses.Response(context, http.StatusOK, dataSources)
}

// UpdateDataSource godoc
// @Summary Update a data source
// @Description Update a data source
// @ID update-data-source
// @Tags data-source-management
// @Accept json
// @Produce json
// @Param id path string true "Data source ID"
// @Param data-source body data_source_request.UpdateDataSourceRequest true "Data source object"
// @Success 200
// @Failure 400 {string} string "BAD_REQUEST"
// @Router /data-source/{id} [patch]
func (h *DataSourceManagementHandler) UpdateDataSource(context *gin.Context) {
	// get request
	id := context.Param("id")
	request := new(data_source_request.UpdateDataSourceRequest)
	if err := context.Bind(&request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Check field is empty
	if id == "" {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, "ID is required")
		return
	}

	// Check field is empty
	if err := request.Validate(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Update data source
	if err := h.Service.Update(id, *request); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	wrapper_responses.Response(context, http.StatusOK, nil)
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

	// Delete data source
	if err := h.Service.Delete(idStr); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	wrapper_responses.Response(context, http.StatusOK, nil)
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
	if err := h.Service.DeleteAll(); err != nil {
		wrapper_responses.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
}
