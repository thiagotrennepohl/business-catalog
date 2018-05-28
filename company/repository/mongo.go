package repository

import (
	"log"
	"strings"
	"sync"

	"github.com/thiagotrennepohl/business-catalog/company"
	"github.com/thiagotrennepohl/business-catalog/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const companyCollection = "companies"

type companyRepository struct {
	session *mgo.Session
}

//NewCompanyRepository returns a Interface for mongo ops
func NewCompanyRepository(conn *mgo.Session) company.CompanyRepository {
	return &companyRepository{
		session: conn,
	}
}

//Bulk inserts many documents at once
func (cr *companyRepository) Bulk(documents []models.Company) error {
	var wg sync.WaitGroup
	errChan := make(chan error)
	session := cr.session.Copy()
	conn := session.DB("").C(companyCollection)
	defer session.Close()
	for _, document := range documents {
		wg.Add(1)
		go cr.updateRecord(errChan, conn, document, &wg)
	}

	go func(errChan chan error, wg *sync.WaitGroup) {
		wg.Wait()
		close(errChan)
	}(errChan, &wg)

	for e := range errChan {
		if e != nil {
			return e
		}
	}
	return nil
}

func (cr *companyRepository) updateRecord(errChan chan error, conn *mgo.Collection, document models.Company, wg *sync.WaitGroup) {
	err := conn.Update(bson.M{"name": strings.ToLower(document.Name)}, document)
	if err == mgo.ErrNotFound {
		//do nothing
	} else {
		errChan <- err
	}
	wg.Done()
}

//Find finds a document using an AND operator -> name and addresszip
func (cr *companyRepository) Find(name string, zip int) ([]models.Company, error) {
	var companies []models.Company
	session := cr.session.Copy()
	// defer session.Close()

	conn := session.DB("").C(companyCollection)

	err := conn.Find(bson.M{"name": bson.M{"$regex": ".*" + name + ".*"}, "addresszip": zip}).All(&companies)
	log.Println(err)
	return companies, err
}
