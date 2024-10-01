package repositories

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	models "github.com/phuckhoa33/web-crawler/internal/domain/models/postgresql"
)

type DataSourceRepository struct {
	DB *gorm.DB
}

func NewDataSourceRepository(db *gorm.DB) *DataSourceRepository {
	return &DataSourceRepository{DB: db}
}

func (r *DataSourceRepository) Create(dataSource *models.DataSource) error {
	return r.DB.Create(dataSource).Error
}

func (r *DataSourceRepository) FindbyID(id uuid.UUID) (*models.DataSource, error) {
	var dataSource models.DataSource
	if err := r.DB.First(&dataSource, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}
	return &dataSource, nil
}

func (r *DataSourceRepository) FindOne(fields map[string]interface{}) (*models.DataSource, error) {
	var dataSource models.DataSource
	query := r.DB

	for field, value := range fields {
		query = query.Where(field+" = ?", value)
	}

	if err := query.First(&dataSource).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &dataSource, nil
}

func (r *DataSourceRepository) FindAll() ([]models.DataSource, error) {
	var dataSources []models.DataSource
	if err := r.DB.Find(&dataSources).Error; err != nil {
		return nil, err
	}
	return dataSources, nil
}

func (r *DataSourceRepository) Update(dataSource *models.DataSource) error {
	return r.DB.Save(dataSource).Error
}

func (r *DataSourceRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.DataSource{}, "id = ?", id.String()).Error
}

func (r *DataSourceRepository) DeleteAll() error {
	return r.DB.Delete(&models.DataSource{}).Error
}
