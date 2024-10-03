package models

import "github.com/phuckhoa33/web-crawler/internal/domain/models/base"

type Data struct {
	*base.AuditModelBase

	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	DetailedUrl string `json:"detailed_url" gorm:"type:varchar(255);not null"`
	Content     string `json:"content" gorm:"type:text;not null"`
	Tag         string `json:"tag" gorm:"type:varchar(255);not null"`
	PrefixUrl   string `json:"prefix_url" gorm:"type:varchar(255);not null"`
	PostedAt    string `json:"posted_at" gorm:"type:timestamp;not null"`
	MediaUrl    string `json:"media_url" gorm:"type:varchar(255)"`

	// Foreign key field
	DataSourceID uint `json:"data_source_id" gorm:"not null"`

	// Association
	DataSource DataSource `json:"data_source" gorm:"foreignkey:DataSourceID;association_foreignkey:ID"`
}
