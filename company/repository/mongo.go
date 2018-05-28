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

type companyRepository struct {
	session *mgo.Session
}

func NewCompanyRepository(conn *mgo.Session) company.CompanyRepository {
	return &companyRepository{
		session: conn,
	}
}

func (cr *companyRepository) Bulk(documents []models.Company) error {
	var wg sync.WaitGroup
	errChan := make(chan error)
	session := cr.session.Copy()
	conn := session.DB("").C("companies")
	defer session.Close()
	for _, document := range documents {
		wg.Add(1)
		go cr.updateRecord(errChan, conn, &document, &wg)
		// close(errChan)
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

func (cr *companyRepository) updateRecord(errChan chan error, conn *mgo.Collection, document *models.Company, wg *sync.WaitGroup) {
	err := conn.Update(bson.M{"namess2": strings.ToLower(document.Name)}, document)
	if err == mgo.ErrNotFound {
		//do nothing
	} else {
		errChan <- err
	}
	wg.Done()
}

func (cr *companyRepository) Find(name string, zip int) ([]models.Company, error) {
	var companies []models.Company
	session := cr.session.Copy()
	// defer session.Close()

	conn := session.DB("").C("companies")

	err := conn.Find(bson.M{"name": bson.M{"$regex": ".*" + name + ".*"}, "addresszip": zip}).All(&companies)
	log.Println(err)
	return companies, err
}
