package main

import (
	"os"

	"github.com/gin-gonic/gin"
	companyHandler "github.com/thiagotrennepohl/business-catalog/company/delivery/http"
	companyRepository "github.com/thiagotrennepohl/business-catalog/company/repository"
	companyUseCase "github.com/thiagotrennepohl/business-catalog/company/usecase"
	"github.com/thiagotrennepohl/sdr/sdr"
	mgo "gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session
var mongoAddr = os.Getenv("MONGO_ADDR")
var CSVDelimiter = ";"

func init() {
	if mongoAddr == "" {
		mongoAddr = "mongodb://localhost:27017/yawoen"
	}
	if CSVDelimiter == "" {
		CSVDelimiter = ""
	}
}

func main() {
	mongoSession, err := mgo.Dial(mongoAddr)
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: CSVDelimiter})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)
	ginEngine.Run()
}
