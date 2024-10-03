package repositories

import models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"

type DataRepository interface {
	Create(data *models.Data) error
	FindByID(id uint) (*models.Data, error)
	FindOne(fields map[string]interface{}) (*models.Data, error)
	FindAll() ([]models.Data, error)
	Update(data *models.Data) error
	Delete(id uint) error
	DeleteAll() error
}
