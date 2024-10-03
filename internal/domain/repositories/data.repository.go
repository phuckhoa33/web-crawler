package repositories

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
)

type DataRepository struct {
	DB *gorm.DB
}

func NewDataRepository(db *gorm.DB) *DataRepository {
	return &DataRepository{DB: db}
}

func (r *DataRepository) Create(data *models.Data) error {
	return r.DB.Create(data).Error
}

func (r *DataRepository) FindbyID(id uuid.UUID) (*models.Data, error) {
	var data models.Data
	if err := r.DB.First(&data, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *DataRepository) FindOne(fields map[string]interface{}) (*models.Data, error) {
	var data models.Data
	query := r.DB

	for field, value := range fields {
		query = query.Where(field+" = ?", value)
	}

	if err := query.First(&data).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}

func (r *DataRepository) FindAll() ([]models.Data, error) {
	var datas []models.Data
	if err := r.DB.Find(&datas).Error; err != nil {
		return nil, err
	}
	return datas, nil
}

func (r *DataRepository) Update(data *models.Data) error {
	return r.DB.Save(data).Error
}

func (r *DataRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Data{}, "id = ?", id.String()).Error
}

func (r *DataRepository) DeleteAll() error {
	return r.DB.Delete(&models.Data{}).Error
}
