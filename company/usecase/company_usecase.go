package usecase

import (
	"encoding/csv"
	"strconv"
	"strings"

	"github.com/thiagotrennepohl/business-catalog/company"
	"github.com/thiagotrennepohl/business-catalog/models"
	"github.com/thiagotrennepohl/sdr/sdr"
)

type companyUseCase struct {
	companyRepository company.CompanyRepository
	sdr               sdr.Sdr
}

func NewCompanyUseCase(cr company.CompanyRepository, sdrClient sdr.Sdr) company.CompanyUseCase {
	return &companyUseCase{
		companyRepository: cr,
		sdr:               sdrClient,
	}
}

func (c *companyUseCase) Find(name string, zip int) ([]models.Company, error) {
	name = strings.ToUpper(name)
	companies, err := c.companyRepository.Find(name, zip)
	return companies, err
}

func (c *companyUseCase) UpdateManyCompanies(companies []models.Company) error {
	err := c.companyRepository.Bulk(companies)
	return err
}

func (c *companyUseCase) ReadCsvFile(filePath string) (*csv.Reader, error) {
	csv, err := c.sdr.ReadCSV(filePath)
	return csv, err
}

func (c *companyUseCase) ParseHeaders(csv *csv.Reader) ([]string, error) {
	headers, err := c.sdr.ParseHeaders(csv)
	return headers, err
}

func (c *companyUseCase) Transform(csv *csv.Reader, headers []string) ([]models.Company, error) {
	var companies []models.Company
	co, err := c.sdr.Extract(csv, headers)
	for _, value := range co {
		company, err := c.cleanData(value)
		if err != nil {
			return companies, err
		}
		companies = append(companies, company)
	}
	return companies, err
}

func (c *companyUseCase) cleanData(data map[string]interface{}) (models.Company, error) {
	zip := data["ADDRESSZIP"].(string)
	if valid := c.validateZip(zip); !valid {
		return models.Company{}, ErrorInvalidCSV
	}
	zipCode, err := c.parseInt(zip)
	if err != nil {
		return models.Company{}, ErrorCouldNotParseZipCode
	}

	var company models.Company
	company.Name = strings.ToUpper(data["NAME"].(string))
	company.Website = strings.ToLower(data["WEBSITE"].(string))
	company.AddressZip = zipCode

	return company, err
}

func (c *companyUseCase) validateZip(zipCode string) bool {
	if len(zipCode) != 5 {
		return false
	}
	return true
}

func (c *companyUseCase) parseInt(zipCode string) (int, error) {
	zipCode = strings.Replace(zipCode, "-", "", -1)
	zip, err := strconv.Atoi(zipCode)
	return zip, err
}
