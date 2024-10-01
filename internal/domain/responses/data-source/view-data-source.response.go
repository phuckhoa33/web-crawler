package data_source_response

import (
	"github.com/google/uuid"
	data_source_enum "github.com/phuckhoa33/web-crawler/internal/enums/data-source"
)

type ViewDataSourcesResponse struct {
	ID       uuid.UUID                 `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name     string                    `json:"name" example:"VnExpress"`
	Url      string                    `json:"url" example:"https://vnexpress.net"`
	Industry data_source_enum.Industry `json:"industry" example:"News"`
}
