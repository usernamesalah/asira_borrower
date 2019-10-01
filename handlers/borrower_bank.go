package handlers

import (
	"asira_borrower/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func BorrowerBankService(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	borrowerModel := models.Borrower{}

	borrowerID, _ := strconv.Atoi(claims["jti"].(string))
	err := borrowerModel.FindbyID(borrowerID)
	if err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "Akun tidak ditemukan")
	}

	bank := models.Bank{}
	bank.FindbyID(int(borrowerModel.Bank.Int64))

	type Filter struct {
		NameOR []string `json:"name" condition:"OR"`
	}
	var service []string
	jMarshal, _ := json.Marshal(bank.Services)
	if err := json.Unmarshal(jMarshal, &service); err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "Service Tidak Ditemukan")
	}

	bankService := models.BankService{}
	result, err := bankService.FilterSearch(&Filter{
		NameOR: service,
	})

	return c.JSON(http.StatusOK, result)
}

func BorrowerBankServiceDetails(c echo.Context) error {
	defer c.Request().Body.Close()
	bServices := models.BankService{}

	serviceID, _ := strconv.Atoi(c.Param("service_id"))
	err := bServices.FindbyID(serviceID)
	if err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "Service Tidak Ditemukan")
	}
	return c.JSON(http.StatusOK, bServices)
}

func BorrowerBankProduct(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	borrowerModel := models.Borrower{}

	borrowerID, _ := strconv.Atoi(claims["jti"].(string))
	err := borrowerModel.FindbyID(borrowerID)
	if err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "Akun tidak ditemukan")
	}

	bank := models.Bank{}
	bank.FindbyID(int(borrowerModel.Bank.Int64))

	type Filter struct {
		NameOR []string `json:"name" condition:"OR"`
	}
	var product []string
	jMarshal, _ := json.Marshal(bank.Products)
	if err := json.Unmarshal(jMarshal, &product); err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "Service Product Tidak Ditemukan")
	}

	bankProduct := models.BankProduct{}
	result, err := bankProduct.FilterSearch(&Filter{
		NameOR: product,
	})

	return c.JSON(http.StatusOK, result)
}

func BorrowerBankProductDetails(c echo.Context) error {
	defer c.Request().Body.Close()
	bankProduct := models.BankProduct{}

	productID, _ := strconv.Atoi(c.Param("product_id"))
	err := bankProduct.FindbyID(productID)
	if err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "Service Product Tidak Ditemukan")
	}
	return c.JSON(http.StatusOK, bankProduct)
}
