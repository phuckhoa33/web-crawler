package repositories

import models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"

type DataSourceRepository interface {
	Create(dataSource *models.DataSource) error
	FindByID(id uint) (*models.DataSource, error)
	FindOne(fields map[string]interface{}) (*models.DataSource, error)
	FindAll() ([]models.DataSource, error)
	Update(dataSource *models.DataSource) error
	Delete(id uint) error
	DeleteAll() error
}
