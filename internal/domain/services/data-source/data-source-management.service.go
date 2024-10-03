package data_source_management_service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/phuckhoa33/web-crawler/cmd/server"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
	"github.com/phuckhoa33/web-crawler/internal/domain/repositories"
	data_source_request "github.com/phuckhoa33/web-crawler/internal/domain/requests/data-source"
	data_source_response "github.com/phuckhoa33/web-crawler/internal/domain/responses/data-source"
)

type DataSourceManagementService struct {
	Repo   *repositories.DataSourceRepository
	Server *server.Server
}

func NewDataSourceManagementService(server *server.Server) *DataSourceManagementService {
	return &DataSourceManagementService{
		Repo:   repositories.NewDataSourceRepository(server.DB),
		Server: server,
	}
}

// Create a new data source
func (s *DataSourceManagementService) Create(dto data_source_request.CreateDataSourceRequest) error {
	// Ping the website and check website is existed
	client := http.Client{
		Timeout: time.Duration(s.Server.Config.App.PingTimeout) * time.Second,
	}
	resp, httpErr := client.Get(dto.Url)
	if httpErr != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cannot ping the website %s", dto.Url)
	}
	defer resp.Body.Close()

	// Check existed data source
	fields := map[string]interface{}{
		"url": dto.Url,
	}
	existedDataSource, repoError := s.Repo.FindOne(fields)
	if repoError != nil {
		return fmt.Errorf("cannot find data source with url %s", dto.Url)
	}

	if existedDataSource != nil {
		return fmt.Errorf("data source with url %s existed", dto.Url)
	}

	// Create data source
	dataSource := models.DataSource{
		Name:     dto.Name,
		Url:      dto.Url,
		Industry: dto.Industry,
	}

	if err := s.Repo.Create(&dataSource); err != nil {
		return fmt.Errorf("cannot create data source with url %s", dto.Url)
	}

	return nil
}

// Get all data sources
func (s *DataSourceManagementService) GetAll() ([]data_source_response.ViewDataSourcesResponse, error) {
	// Get data source
	dataSource, err := s.Repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("cannot get all data sources")
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

	return response, nil
}

// Get a data source
func (s *DataSourceManagementService) GetOne(id string) (*data_source_response.ViewDataSourcesResponse, error) {
	// Check existed data source
	fields := map[string]interface{}{
		"id": id,
	}
	dataSource, repoError := s.Repo.FindOne(fields)
	if repoError != nil {
		return nil, fmt.Errorf("cannot find data source with id %s", id)
	}

	if dataSource == nil {
		return nil, fmt.Errorf("data source with id %s not found", id)
	}

	// map data source
	response := data_source_response.ViewDataSourcesResponse{
		Name:     dataSource.Name,
		Url:      dataSource.Url,
		Industry: dataSource.Industry,
		ID:       dataSource.ID,
	}

	return &response, nil
}

// Update a data source
func (s *DataSourceManagementService) Update(id string, dto data_source_request.UpdateDataSourceRequest) error {
	// Check existed data source
	fields := map[string]interface{}{
		"id": id,
	}
	dataSource, repoError := s.Repo.FindOne(fields)
	if repoError != nil {
		return fmt.Errorf("cannot find data source with id %s", id)
	}

	if dataSource == nil {
		return fmt.Errorf("data source with id %s not found", id)
	}

	// Update data source
	dataSource.Name = dto.Name
	dataSource.Url = dto.Url
	dataSource.Industry = dto.Industry

	if err := s.Repo.Update(dataSource); err != nil {
		return fmt.Errorf("cannot update data source with id %s", id)
	}

	return nil

}

// Delete a data source
func (s *DataSourceManagementService) Delete(id string) error {
	// Check existed data source
	fields := map[string]interface{}{
		"id": id,
	}
	dataSource, repoError := s.Repo.FindOne(fields)
	if repoError != nil {
		return fmt.Errorf("cannot find data source with id %s", id)
	}

	if dataSource == nil {
		return fmt.Errorf("data source with id %s not found", id)
	}

	// Delete data source
	if err := s.Repo.Delete(dataSource.ID); err != nil {
		return fmt.Errorf("cannot delete data source with id %s", id)
	}

	return nil
}

// Delete all data sources
func (s *DataSourceManagementService) DeleteAll() error {
	// Delete all data sources
	if err := s.Repo.DeleteAll(); err != nil {
		return fmt.Errorf("cannot delete all data sources")
	}

	return nil
}
