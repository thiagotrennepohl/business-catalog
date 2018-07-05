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

//NewCompanyUseCase is responsible for creating a new companyUseCase which implements the interface
//company.CompanyUseCase
func NewCompanyUseCase(cr company.CompanyRepository, sdrClient sdr.Sdr) company.CompanyUseCase {
	return &companyUseCase{
		companyRepository: cr,
		sdr:               sdrClient,
	}
}

//Find performs a company search using name and zip
func (c *companyUseCase) Find(name string, zip int) ([]models.Company, error) {
	name = strings.ToUpper(name)
	companies, err := c.companyRepository.Find(name, zip)
	return companies, err
}

//UpdateManyCompaies performs a bulk update to all companies stored in mongo
func (c *companyUseCase) UpdateManyCompanies(companies []models.Company) error {
	err := c.companyRepository.Bulk(companies)
	return err
}

//ReadCsvFile reads a csvFile and returns it's reader pointer
func (c *companyUseCase) ReadCsvFile(filePath string) (*csv.Reader, error) {
	csv, err := c.sdr.ReadCSV(filePath)
	return csv, err
}

//ParseHeaders validates and fixes csv headers by removing "-" replacing spaces for "_"
//and making sure all char's are lowercase
func (c *companyUseCase) ParseHeaders(csv *csv.Reader) ([]string, error) {
	headers, err := c.sdr.ParseHeaders(csv)
	if err != nil {
		return []string{""}, err
	}
	err = c.validadeHeaders(headers)
	return headers, err
}

func (c *companyUseCase) validadeHeaders(headers []string) error {
	if len(headers) > 3 {
		return ErrorTooManyHeaders
	}
	for _, header := range headers {
		switch header {
		case NameHeader:
			continue
		case AddressZipHeader:
			continue
		case WebSiteHeader:
			continue
		default:
			return ErrorInvalidHeaders
		}
	}
	return nil
}

//Transform takes an csv.Reader as input and returns an slice of Companies
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
	zip := data[AddressZipHeader].(string)
	if valid := c.validateZip(zip); !valid {
		return models.Company{}, ErrorInvalidCSV
	}
	zipCode, err := c.parseInt(zip)
	if err != nil {
		return models.Company{}, ErrorCouldNotParseZipCode
	}

	var company models.Company
	company.Name = strings.ToUpper(data[NameHeader].(string))
	company.Website = strings.ToLower(data[WebSiteHeader].(string))
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
