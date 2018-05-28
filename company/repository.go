package company

import "github.com/thiagotrennepohl/business-catalog/models"

type CompanyRepository interface {
	Bulk([]models.Company) error
	Find(string, int) ([]models.Company, error)
}
