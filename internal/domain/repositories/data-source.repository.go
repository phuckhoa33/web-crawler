package repositories

import (
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

func (r *DataSourceRepository) FindbyID(id uint) (*models.DataSource, error) {
	var dataSource models.DataSource
	if err := r.DB.First(&dataSource, id).Error; err != nil {
		return nil, err
	}
	return &dataSource, nil
}

func (r *DataSourceRepository) FindOne(property string, value string) (*models.DataSource, error) {
	var dataSource models.DataSource
	if err := r.DB.Where(property+" = ?", value).Error; err != nil {
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

func (r *DataSourceRepository) Delete(id uint) error {
	return r.DB.Delete(&models.DataSource{}, id).Error
}
