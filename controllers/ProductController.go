package controllers

import (
	"MytheresaChallenge/database/connections"
	"MytheresaChallenge/requests"
	"MytheresaChallenge/resources"
	"MytheresaChallenge/services"
	"MytheresaChallenge/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetProducts gets query parameters from url and returns related products in json type
func GetProducts(c *gin.Context)  {
	output := utils.NewOutput()

	request := requests.ProductRequest{}

	//get query parameters if exist
	category, ok := c.GetQuery("category")
	if ok {
		request.Category = &category
	}

	priceLessThan, ok := c.GetQuery("priceLessThan")
	if ok {
		request.PriceLessThan = &priceLessThan
	}

	//performs query and returns model
	products, err := services.GetProducts(connections.DB, request)
	if err!=nil && err.Error() != "record not found" {
		utils.SetError("NO_ITEM_FOUND", c, output, http.StatusExpectationFailed, http.StatusExpectationFailed)
		return
	}

	//transforms result models to output format and apply discount if needed
	productsOutput := resources.GetProductResource(products)

	output.Data = productsOutput

	//returns final response as a json
	c.JSON(http.StatusOK, output)
	return
}
