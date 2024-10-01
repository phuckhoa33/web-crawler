package data_source_request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	data_source_enum "github.com/phuckhoa33/web-crawler/internal/enums/data-source"
)

type CreateDataSourceRequest struct {
	Name     string                    `json:"name" validate:"required" example:"VnExpress"`
	Url      string                    `json:"url" validate:"required" example:"https://vnexpress.net"`
	Industry data_source_enum.Industry `json:"industry" validate:"required" example:"News"`
}

func (r *CreateDataSourceRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&r.Url, validation.Required, is.URL),
		validation.Field(&r.Industry, validation.Required, validation.In(
			data_source_enum.IndustryEducation,
			data_source_enum.IndustryFinance,
			data_source_enum.IndustryTechnology,
			data_source_enum.IndustryHealth,
			data_source_enum.IndustryEducation,
			data_source_enum.IndustryNews,
		)),
	)
}
