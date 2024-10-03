package repositories

import models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"

type DataAccessWayRepository interface {
	Create(dataAccess *models.DataAccessWay) error
	FindByID(id uint) (*models.DataAccessWay, error)
	FindOne(fields map[string]interface{}) (*models.DataAccessWay, error)
	FindAll() ([]models.DataAccessWay, error)
	Update(dataAccess *models.DataAccessWay) error
	Delete(id uint) error
	DeleteAll() error
}
