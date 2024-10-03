package datasource_test

import (
	"testing"

	"github.com/phuckhoa33/web-crawler/cmd/server"
	"github.com/phuckhoa33/web-crawler/internal/config"
	data_source_request "github.com/phuckhoa33/web-crawler/internal/domain/requests/data-source"
	data_source_management_service "github.com/phuckhoa33/web-crawler/internal/domain/services/data-source"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceManagementService_CreateDataSource(t *testing.T) {
	mockRepo := new(MockDataSourceRepository)
	service := data_source_management_service.NewDataSourceManagementService(server.NewServer(config.NewConfig()))

	dto := &data_source_request.CreateDataSourceRequest{
		Name:     "Test Source",
		Url:      "http://example.com",
		Industry: "Tech",
	}

	// Set up expectations
	mockRepo.On("FindOne", mock.Anything).Return(nil, nil)
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Call the method
	err := service.Create(*dto)

	// Assert expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
