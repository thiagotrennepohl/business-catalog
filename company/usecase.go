package company

import (
	"encoding/csv"

	"github.com/thiagotrennepohl/business-catalog/models"
)

type CompanyUseCase interface {
	ReadCsvFile(string) (*csv.Reader, error)
	ParseHeaders(csv *csv.Reader) ([]string, error)
	UpdateManyCompanies(companies []models.Company) error
	Transform(*csv.Reader, []string) ([]models.Company, error)
	Find(name string, zip int) ([]models.Company, error)
}
