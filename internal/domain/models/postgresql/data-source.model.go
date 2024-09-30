package models

import (
	"github.com/phuckhoa33/web-crawler/internal/domain/models/base"
	data_source_enum "github.com/phuckhoa33/web-crawler/internal/enums/data-source"
)

type DataSource struct {
	*base.AuditModelBase
	Name     string                    `json:"name" gorm:"type:varchar(255);not null"`
	Url      string                    `json:"url" gorm:"type:varchar(255);not null"`
	Industry data_source_enum.Industry `json:"industry" gorm:"type:varchar(255);not null"`
}
