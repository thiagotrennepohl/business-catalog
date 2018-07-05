package http_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	companyHandler "github.com/thiagotrennepohl/business-catalog/company/delivery/http"
	companyRepository "github.com/thiagotrennepohl/business-catalog/company/repository"
	companyUseCase "github.com/thiagotrennepohl/business-catalog/company/usecase"
	"github.com/thiagotrennepohl/sdr/sdr"
	mgo "gopkg.in/mgo.v2"
)

const CSVDelimiter = ";"

func TestSucessFindCompany(t *testing.T) {
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: CSVDelimiter})
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
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: CSVDelimiter})
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
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: CSVDelimiter})
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

	mongoSession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: CSVDelimiter})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)

	response := httptest.NewRecorder()
	endpoint := "/v1/company"

	//File handling for http form
	fileDir, _ := filepath.Abs("../../../assets/")
	fileName := "/q2_clientData.csv"
	filePath := fileDir + fileName

	//Create Temp Dir for testing
	_, err = os.Stat("assets")
	if err != nil {
		err = os.MkdirAll("assets", 0755)
		if err != nil {
			assert.Fail(t, err.Error())
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Failed %e", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("data", filePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ginEngine.ServeHTTP(response, req)
	respBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.Code, string(respBody))

	err = os.RemoveAll("assets")
	if err != nil {
		assert.Fail(t, err.Error())
	}
}

func TestUpdateCompaniesWithWorngFormData(t *testing.T) {

	mongoSession, err := mgo.Dial("mongodb://localhost:27017/testbulk")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: CSVDelimiter})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)

	response := httptest.NewRecorder()
	endpoint := "/v1/company"

	//File handling for http form
	fileDir, _ := filepath.Abs("../../../assets/")
	fileName := "/q2_clientData.csv"
	filePath := fileDir + fileName

	//Create Temp Dir for testing
	_, err = os.Stat("assets")
	if err != nil {
		err = os.MkdirAll("assets", 0755)
		if err != nil {
			assert.Fail(t, err.Error())
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Failed %e", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ginEngine.ServeHTTP(response, req)
	respBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.Code, string(respBody))

	err = os.RemoveAll("assets")
	if err != nil {
		assert.Fail(t, err.Error())
	}
}

func TestUpdateCompaniesWithoutAssetsFolder(t *testing.T) {

	mongoSession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: CSVDelimiter})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)

	response := httptest.NewRecorder()
	endpoint := "/v1/company"

	//File handling for http form
	fileDir, _ := filepath.Abs("../../../assets/")
	fileName := "/q2_clientData.csv"
	filePath := fileDir + fileName

	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Failed %e", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("data", filePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ginEngine.ServeHTTP(response, req)
	respBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.Code, string(respBody))
}

func TestUpdateCompaniesWithWrongHeaders(t *testing.T) {
	mongoSession, err := mgo.Dial("mongodb://localhost:27017/yawoen")
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()
	companyRepo := companyRepository.NewCompanyRepository(mongoSession)
	sdr := sdr.NewSdr(sdr.SdrConfig{CommaDelimiter: CSVDelimiter})
	companyUcase := companyUseCase.NewCompanyUseCase(companyRepo, sdr)
	companyHandler.NewHttpHandler(ginEngine, companyUcase)

	response := httptest.NewRecorder()
	endpoint := "/v1/company"

	//File handling for http form
	fileDir, _ := filepath.Abs("../../../assets/")
	fileName := "/q1_wrongHeaders.csv"
	filePath := fileDir + fileName

	_, err = os.Stat("assets")
	if err != nil {
		err = os.MkdirAll("assets", 0755)
		if err != nil {
			assert.Fail(t, err.Error())
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Failed %e", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("data", filePath)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	ginEngine.ServeHTTP(response, req)
	respBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.Code, string(respBody))

	err = os.RemoveAll("assets")
	if err != nil {
		assert.Fail(t, err.Error())
	}
}
