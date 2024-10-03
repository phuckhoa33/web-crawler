package datasource_test

import (
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
	"github.com/stretchr/testify/mock"
)

type MockDataSourceRepository struct {
	mock.Mock
}

func (m *MockDataSourceRepository) FindOne(fields map[string]interface{}) (*models.DataSource, error) {
	args := m.Called(fields)
	if args.Get(0) != nil {
		return args.Get(0).(*models.DataSource), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataSourceRepository) Create(dataSource *models.DataSource) error {
	args := m.Called(dataSource)
	return args.Error(0)
}

func (m *MockDataSourceRepository) FindAll() ([]models.DataSource, error) {
	args := m.Called()
	return args.Get(0).([]models.DataSource), args.Error(1)
}

func (m *MockDataSourceRepository) FindbyID(id uint) (*models.DataSource, error) {

	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.DataSource), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDataSourceRepository) Update(dataSource *models.DataSource) error {
	args := m.Called(dataSource)
	return args.Error(0)
}

func (m *MockDataSourceRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDataSourceRepository) DeleteAll() error {
	args := m.Called()
	return args.Error(0)
}
