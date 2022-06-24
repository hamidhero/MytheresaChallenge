package services

import (
	"MytheresaChallenge/models"
	"MytheresaChallenge/requests"
	"github.com/sonyarouje/simdb"
	"strconv"
)

//GetProducts accepts query parameters, fetch related data and returns model or error if it has one
func GetProducts(database *simdb.Driver, request requests.ProductRequest) ([]models.Product, error) {
	//initiate the query
	db := database.Open(models.Product{})

	//add conditions if required
	if request.Category != nil {
		db = db.Where("category","=", *request.Category)
	}

	if request.PriceLessThan != nil {
		price, _ := strconv.ParseInt(*request.PriceLessThan,10, 64)
		db = db.Where("price","<=", price)
	}

	//get data from database
	var products []models.Product
	err := db.Get().AsEntity(&products)

	//as required, number of results must not exceed 5
	if len(products) > 5 {
		return products[:5], err
	}

	return products, err
}
