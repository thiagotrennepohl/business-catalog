package main

import (
	"github.com/gin-gonic/gin"
	companyHandler "github.com/thiagotrennepohl/business-catalog/company/delivery/http"
	companyRepository "github.com/thiagotrennepohl/business-catalog/company/repository"
	companyUseCase "github.com/thiagotrennepohl/business-catalog/company/usecase"
	"github.com/thiagotrennepohl/sdr/sdr"
	mgo "gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session

func main() {
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: ';'})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)
	ginEngine.Run()
}
