package handlers

import (
	"asira_borrower/asira"
	"asira_borrower/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

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

	db := asira.App.DB
	var results []models.Product
	var count int

	err = db.Table("banks b").
		Select("p.*").
		Joins("INNER JOIN borrowers bo ON bo.bank = b.id").
		Joins("INNER JOIN services s ON s.id IN (SELECT UNNEST(b.services))").
		Joins("INNER JOIN products p ON p.service_id = s.id AND p.service_id IN (SELECT UNNEST(b.products))").
		Where("bo.id = ?", borrowerID).Find(&results).Count(&count).Error

	if err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "Service Product Tidak Ditemukan")
	}

	type Result struct {
		TotalData int              `json:"total_data"`
		Data      []models.Product `json:"data"`
	}

	return c.JSON(http.StatusOK, &Result{TotalData: count, Data: results})
}

func BorrowerBankProductDetails(c echo.Context) error {
	defer c.Request().Body.Close()
	bankProduct := models.Product{}

	productID, _ := strconv.Atoi(c.Param("product_id"))
	err := bankProduct.FindbyID(productID)
	if err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "Service Product Tidak Ditemukan")
	}
	return c.JSON(http.StatusOK, bankProduct)
}
