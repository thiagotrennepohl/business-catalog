package http_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	companyHandler "github.com/thiagotrennepohl/business-catalog/company/delivery/http"
	companyRepository "github.com/thiagotrennepohl/business-catalog/company/repository"
	companyUseCase "github.com/thiagotrennepohl/business-catalog/company/usecase"
	"github.com/thiagotrennepohl/sdr/sdr"
	mgo "gopkg.in/mgo.v2"
)

func TestSucessFindCompany(t *testing.T) {
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/testbulk")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: ';'})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)

	response := httptest.NewRecorder()
	endpoint := "/v1/company"

	req, _ := http.NewRequest("GET", endpoint, nil)
	queryParams := req.URL.Query()
	queryParams.Add("name", "t")
	queryParams.Add("zip", "12345")
	req.URL.RawQuery = queryParams.Encode()

	ginEngine.ServeHTTP(response, req)
	respBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.Code, string(respBody))
}

func TestFindCompanyWithoutZipParam(t *testing.T) {
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/testbulk")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: ';'})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)

	response := httptest.NewRecorder()
	endpoint := "/v1/company"

	req, _ := http.NewRequest("GET", endpoint, nil)
	queryParams := req.URL.Query()
	queryParams.Add("name", "t")
	req.URL.RawQuery = queryParams.Encode()

	ginEngine.ServeHTTP(response, req)
	respBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.Code, string(respBody))
}

func TestFindCompanyWithoutNameParam(t *testing.T) {
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/testbulk")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: ';'})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)

	response := httptest.NewRecorder()
	endpoint := "/v1/company"

	req, _ := http.NewRequest("GET", endpoint, nil)
	queryParams := req.URL.Query()
	req.URL.RawQuery = queryParams.Encode()

	ginEngine.ServeHTTP(response, req)
	respBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.Code, string(respBody))
}

func TestUpdateCompanies(t *testing.T) {

	mongoSession, err := mgo.Dial("mongodb://localhost:27017/testbulk")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: ';'})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)

	response := httptest.NewRecorder()
	endpoint := "/v1/company"

	file, err := os.Open("../../../assets/q2_clientData.csv")
	if err != nil {
		t.Errorf("Failed %e", err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()
	part, err := writer.CreateFormFile("data", "q2_clientData.csv")
	if err != nil {
		t.Errorf("Failed while creating form %e", err)
	}
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Contet-Type", writer.FormDataContentType())

	ginEngine.ServeHTTP(response, req)
	respBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.Code, string(respBody))

}
