package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thiagotrennepohl/business-catalog/company"
	"github.com/thiagotrennepohl/business-catalog/models"
)

type CompanyHTTPHandler struct {
	companyUseCase company.CompanyUseCase
}

func (co *CompanyHTTPHandler) UpdateCompanies(c *gin.Context) {
	companies := []models.Company{}

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	err = c.SaveUploadedFile(file, "assets/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	csv, err := co.companyUseCase.ReadCsvFile("assets/" + file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	headers, err := co.companyUseCase.ParseHeaders(csv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	companies, err = co.companyUseCase.Transform(csv, headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	err = co.companyUseCase.UpdateManyCompanies(companies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (co *CompanyHTTPHandler) SearchCompany(c *gin.Context) {
	companyName := c.Query("name")
	companyZip, err := strconv.Atoi(c.Query("zip"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	companies, err := co.companyUseCase.Find(companyName, companyZip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, companies)
}

func NewHttpHandler(g *gin.Engine, companyUseCase company.CompanyUseCase) {
	handler := &CompanyHTTPHandler{
		companyUseCase: companyUseCase,
	}

	g.POST("/v1/company", handler.UpdateCompanies)
	g.GET("/v1/company", handler.SearchCompany)
}
