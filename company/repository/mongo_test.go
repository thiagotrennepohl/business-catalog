package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagotrennepohl/business-catalog/company/repository"
	"github.com/thiagotrennepohl/business-catalog/models"
	mgo "gopkg.in/mgo.v2"
)

func TestUpdateMany(t *testing.T) {
	mongosession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		t.Error(err)
	}
	companyRepo := repository.NewCompanyRepository(mongosession)

	companies := []models.Company{
		models.Company{
			Name:       "TEST",
			AddressZip: 12345,
			Website:    "test.com",
		},
	}

	err = companyRepo.Bulk(companies)

	assert.NoError(t, err, err)
}

func TestFindCompany(t *testing.T) {
	company := models.Company{
		Name:       "TEST",
		AddressZip: 12345,
		Website:    "test.com",
	}

	mongosession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		t.Error(err)
	}

	conn := mongosession.DB("").C("companies")
	conn.Insert(company)

	companyRepo := repository.NewCompanyRepository(mongosession)

	result, err := companyRepo.Find("T", 12345)

	if len(result) == 0 {
		t.Errorf("The result is invalid!")
	}

	assert.NoError(t, err)

	assert.Equal(t, company.Name, result[0].Name)
}
