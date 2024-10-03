package models

import "github.com/phuckhoa33/web-crawler/internal/domain/models/base"

type DataAccessWay struct {
	*base.ModelBase

	StartId           int    `json:"start_number" gorm:"type:int;not null"`
	EndId             int    `json:"end_number" gorm:"type:int;not null"`
	StartPage         int    `json:"start_page" gorm:"type:int;not null"`
	EndPage           int    `json:"end_page" gorm:"type:int;not null"`
	Step              int    `json:"step" gorm:"type:int;not null"`
	Url               string `json:"url" gorm:"type:varchar(255);not null"`
	TitleClassName    string `json:"title_class_name" gorm:"type:varchar(255);not null"`
	ContentClassName  string `json:"content_class_name" gorm:"type:varchar(255);not null"`
	TagClassName      string `json:"tag_class_name" gorm:"type:varchar(255);not null"`
	PostedAtClassName string `json:"posted_at_class_name" gorm:"type:varchar(255);not null"`
	MediaUrlClassName string `json:"media_url_class_name" gorm:"type:varchar(255)"`

	// Foreign key field
	DataSourceID uint `json:"data_source_id" gorm:"not null"`

	// Association
	DataSource DataSource `json:"data_source" gorm:"foreignkey:DataSourceID;association_foreignkey:ID"`
}
