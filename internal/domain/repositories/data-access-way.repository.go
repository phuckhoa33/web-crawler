package repositories

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
)

type DataAccessWayRepository struct {
	DB *gorm.DB
}

func NewDataAccessWayRepository(db *gorm.DB) *DataAccessWayRepository {
	return &DataAccessWayRepository{DB: db}
}

func (r *DataAccessWayRepository) Create(dataAccess *models.DataAccessWay) error {
	return r.DB.Create(dataAccess).Error
}

func (r *DataAccessWayRepository) FindbyID(id uuid.UUID) (*models.DataAccessWay, error) {
	var dataAccess models.DataAccessWay
	if err := r.DB.First(&dataAccess, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	return &dataAccess, nil
}

func (r *DataAccessWayRepository) FindOne(fields map[string]interface{}) (*models.DataAccessWay, error) {
	var dataAccess models.DataAccessWay
	query := r.DB

	for field, value := range fields {
		query = query.Where(field+" = ?", value)
	}

	if err := query.First(&dataAccess).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &dataAccess, nil
}

func (r *DataAccessWayRepository) FindAll() ([]models.DataAccessWay, error) {
	var dataAccesss []models.DataAccessWay
	if err := r.DB.Find(&dataAccesss).Error; err != nil {
		return nil, err
	}
	return dataAccesss, nil
}

func (r *DataAccessWayRepository) Update(dataAccess *models.DataAccessWay) error {
	return r.DB.Save(dataAccess).Error
}

func (r *DataAccessWayRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.DataAccessWay{}, "id = ?", id.String()).Error
}

func (r *DataAccessWayRepository) DeleteAll() error {
	return r.DB.Delete(&models.DataAccessWay{}).Error
}
