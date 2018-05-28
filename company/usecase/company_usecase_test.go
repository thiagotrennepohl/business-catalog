package usecase

import (
	"encoding/csv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thiagotrennepohl/business-catalog/company/mocks"
	"github.com/thiagotrennepohl/business-catalog/models"
)

func TestReadCsvFile(t *testing.T) {
	var csvReader *csv.Reader

	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)
	companySdrMock.On("ReadCSV", mock.AnythingOfType("string")).Return(csvReader, nil)

	companyUseCase := NewCompanyUseCase(companyRepositoryMock, companySdrMock)
	result, err := companyUseCase.ReadCsvFile("somePath")

	assert.NoError(t, err)
	assert.ObjectsAreEqualValues(csvReader, result)
}

func TestParseHeaders(t *testing.T) {
	var csvFile *csv.Reader
	var headers []string

	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)

	companySdrMock.On("ParseHeaders", mock.Anything).Return(headers, nil)

	companyUseCase := NewCompanyUseCase(companyRepositoryMock, companySdrMock)
	result, err := companyUseCase.ParseHeaders(csvFile)

	assert.NoError(t, err)
	assert.ObjectsAreEqualValues(headers, result)

}

func TestInvalidZipCode(t *testing.T) {
	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)

	companyUseCase := &companyUseCase{
		companyRepository: companyRepositoryMock,
		sdr:               companySdrMock,
	}
	valid := companyUseCase.validateZip("123")

	assert.False(t, valid)
}

func TestValidZipCode(t *testing.T) {
	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)

	companyUseCase := &companyUseCase{
		companyRepository: companyRepositoryMock,
		sdr:               companySdrMock,
	}
	valid := companyUseCase.validateZip("12345")

	assert.True(t, valid)
}

func TestZipCodeParsing(t *testing.T) {
	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)

	companyUseCase := &companyUseCase{
		companyRepository: companyRepositoryMock,
		sdr:               companySdrMock,
	}
	valid, err := companyUseCase.parseInt("12345")
	assert.NoError(t, err)
	assert.Equal(t, 12345, valid)
}

func TestDataCleaning(t *testing.T) {
	dataParam := map[string]interface{}{
		"NAME":       "Some company",
		"WEBSITE":    "HTTP://SOMEWEBSITE",
		"ADDRESSZIP": "12345",
	}

	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)

	companyUseCase := &companyUseCase{
		companyRepository: companyRepositoryMock,
		sdr:               companySdrMock,
	}

	company, err := companyUseCase.cleanData(dataParam)

	assert.NoError(t, err)
	assert.Equal(t, company.Name, "SOME COMPANY")
	assert.Equal(t, company.Website, "http://somewebsite")
	assert.Equal(t, company.AddressZip, 12345)
}

func TestInvalidZipOnClean(t *testing.T) {
	dataParam := map[string]interface{}{
		"NAME":       "Some company",
		"WEBSITE":    "HTTP://SOMEWEBSITE",
		"ADDRESSZIP": "1234",
	}

	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)

	companyUseCase := &companyUseCase{
		companyRepository: companyRepositoryMock,
		sdr:               companySdrMock,
	}

	_, err := companyUseCase.cleanData(dataParam)

	assert.Error(t, err)
}

// Is working as expected but I feel it this isn't a useful test
func TestZipParsingError(t *testing.T) {
	dataParam := map[string]interface{}{
		"NAME":       "Some company",
		"WEBSITE":    "HTTP://SOMEWEBSITE",
		"ADDRESSZIP": "somenumber",
	}

	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)

	companyUseCase := &companyUseCase{
		companyRepository: companyRepositoryMock,
		sdr:               companySdrMock,
	}

	company, err := companyUseCase.cleanData(dataParam)

	assert.Error(t, err, err)
	assert.Zero(t, company)
}

func TestTransform(t *testing.T) {
	var csvReader *csv.Reader
	var headers []string
	var jsonData []map[string]interface{}
	var companies []models.Company

	companySdrMock := new(mocks.Sdr)
	companyRepositoryMock := new(mocks.CompanyRepository)
	companySdrMock.On("Extract", mock.Anything, mock.AnythingOfType("[]string")).Return(jsonData, nil)

	companyUseCase := NewCompanyUseCase(companyRepositoryMock, companySdrMock)
	result, err := companyUseCase.Transform(csvReader, headers)

	assert.NoError(t, err)
	assert.Equal(t, companies, result)
	assert.Len(t, companies, 0)
}
func TestUpdateManyCompanies(t *testing.T) {
	companyMockRepo := new(mocks.CompanyRepository)
	companySdrMock := new(mocks.Sdr)

	company1 := models.Company{
		Name:       "Yawoen",
		AddressZip: 12345,
		Website:    "yawoen.com",
	}

	company2 := models.Company{
		Name:       "Boring company",
		AddressZip: 99999,
		Website:    "boringcompanycom",
	}

	mockCompanies := make([]models.Company, 0)
	mockCompanies = append(mockCompanies, company1, company2)
	companyMockRepo.On("Bulk", mock.Anything).Return(nil)

	u := NewCompanyUseCase(companyMockRepo, companySdrMock)

	err := u.UpdateManyCompanies(mockCompanies)
	assert.Equal(t, err, nil)
	assert.NoError(t, err)
}

func TestFind(t *testing.T) {
	companyMockRepo := new(mocks.CompanyRepository)
	companySdrMock := new(mocks.Sdr)

	mockCompany := models.Company{
		Name:       "Yawoen",
		AddressZip: 12345,
	}

	mockCompanies := make([]models.Company, 0)
	mockCompanies = append(mockCompanies, mockCompany)
	companyMockRepo.On("Find", mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(mockCompanies, nil)

	u := NewCompanyUseCase(companyMockRepo, companySdrMock)

	companies, err := u.Find("yawoen", 12345)
	assert.Equal(t, companies, mockCompanies)
	assert.NoError(t, err)
	assert.Len(t, companies, len(mockCompanies))
	// cursor := "12"
	// list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
	// cursorExpected := strconv.Itoa(int(mockArticle.ID))
	// assert.Equal(t, cursorExpected, nextCursor)
	// assert.NotEmpty(t, nextCursor)
	// assert.NoError(t, err)
	// assert.Len(t, list, len(mockListArtilce))

	// mockArticleRepo.AssertExpectations(t)
	// mockAuthorrepo.AssertExpectations(t)
}
