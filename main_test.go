package main

import (
	"MytheresaChallenge/database"
	"MytheresaChallenge/database/connections"
	"MytheresaChallenge/models"
	"MytheresaChallenge/requests"
	"MytheresaChallenge/resources"
	"MytheresaChallenge/router"
	"MytheresaChallenge/services"
	"MytheresaChallenge/utils"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

/***** utils *****/
func TestApplyDiscount(t *testing.T) {
	price := utils.ApplyDiscount(25000, 40)
	assert.Equal(t, int64(15000), price)
}

func TestNewOutput(t *testing.T) {
	output := utils.NewOutput()
	assert.Equal(t, 200, output.Status)
}

/***** database *****/
func TestInit(t *testing.T) {
	_, err := database.Init()
	assert.Equal(t, nil, err)
}

/***** services *****/
func TestGetProducts(t *testing.T) {
	db, err := database.Init()
	assert.Equal(t, nil, err)

	getProductsInput := requests.ProductRequest{}
	_, err = services.GetProducts(db, getProductsInput)
	assert.Equal(t, nil, err)

	category := "boots"
	getProductsInput = requests.ProductRequest{Category: &category}
	products, err := services.GetProducts(db, getProductsInput)
	assert.Equal(t, nil, err)
	if len(products) == 0 || products[0].Category != "boots" {
		t.Error("wrong output")
	}

	category = "airplanes"
	getProductsInput = requests.ProductRequest{Category: &category}
	_, err = services.GetProducts(db, getProductsInput)
	assert.NotEqual(t, nil, err)

	priceLessThan := "70000"
	getProductsInput = requests.ProductRequest{PriceLessThan: &priceLessThan}
	products, err = services.GetProducts(db, getProductsInput)
	assert.Equal(t, nil, err)
	if len(products) == 0 || products[0].Price > 70000 {
		t.Error("wrong output")
	}

	priceLessThan = "700"
	getProductsInput = requests.ProductRequest{PriceLessThan: &priceLessThan}
	products, err = services.GetProducts(db, getProductsInput)
	assert.NotEqual(t, nil, err)

	priceLessThan = "Not Number"
	getProductsInput = requests.ProductRequest{PriceLessThan: &priceLessThan}
	products, err = services.GetProducts(db, getProductsInput)
	assert.NotEqual(t, nil, err)

	category = "boots"
	priceLessThan = "90000"
	getProductsInput = requests.ProductRequest{Category: &category, PriceLessThan: &priceLessThan}
	products, err = services.GetProducts(db, getProductsInput)
	assert.Equal(t, nil, err)
	if len(products) == 0 || products[0].Category != "boots" || products[0].Price > 90000 {
		t.Error("wrong output")
	}

	category = "airplanes"
	priceLessThan = "90000"
	getProductsInput = requests.ProductRequest{Category: &category, PriceLessThan: &priceLessThan}
	products, err = services.GetProducts(db, getProductsInput)
	assert.NotEqual(t, nil, err)

	category = "boots"
	priceLessThan = "Not Number"
	getProductsInput = requests.ProductRequest{Category: &category, PriceLessThan: &priceLessThan}
	products, err = services.GetProducts(db, getProductsInput)
	assert.NotEqual(t, nil, err)

	category = "airplanes"
	priceLessThan = "Not Number"
	getProductsInput = requests.ProductRequest{Category: &category, PriceLessThan: &priceLessThan}
	products, err = services.GetProducts(db, getProductsInput)
	assert.NotEqual(t, nil, err)
}

/***** resources *****/
func TestGetPrice(t *testing.T)  {
	db, err := database.Init()
	assert.Equal(t, nil, err)

	category := "boots"
	getProductsInput := requests.ProductRequest{Category: &category}
	products, err := services.GetProducts(db, getProductsInput)

	p := resources.ProductGetPriceResource{Product:products[0]}
	price, discount := p.GetPrice()
	assert.NotEqual(t, 0, price)
	assert.NotEqual(t, nil, discount)

	category = "sneakers"
	getProductsInput = requests.ProductRequest{Category: &category}
	products, err = services.GetProducts(db, getProductsInput)

	p = resources.ProductGetPriceResource{Product:products[0]}
	price, discount = p.GetPrice()
	assert.NotEqual(t, 0, price)
	assert.Equal(t, nil, discount)

}

func TestGetProductResource(t *testing.T) {
	var productResourceInput []models.Product
	output := resources.GetProductResource(productResourceInput)
	assert.Equal(t, 0, len(output))

	db, err := database.Init()
	assert.Equal(t, nil, err)

	category := "boots"
	getProductsInput := requests.ProductRequest{Category: &category}
	products, err := services.GetProducts(db, getProductsInput)

	output = resources.GetProductResource(products)

	if len(output) == 0 || output[0].Category != "boots" ||
		output[0].Price.DiscountPercentage == nil ||
		output[0].Price.Final == output[0].Price.Original {
		t.Error("wrong output")
	}
}

/***** controllers *****/
func TestGetProductsEndpoint(t *testing.T) {
	connections.Connect()
	route := router.GetRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/products?category=airplanes", nil)
	route.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	output := utils.Output{}
	err := json.Unmarshal(w.Body.Bytes(), &output)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, output.Status)
	assert.Equal(t, nil, output.Data)
	assert.Equal(t, nil, output.Error)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/api/products?category=boots", nil)
	route.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &output)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, output.Status)
	assert.NotEqual(t, nil, output.Data)
	assert.NotEqual(t, nil, output.Data)
	assert.Equal(t, nil, output.Error)

	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/api/products?priceLessThan=80000", nil)
	route.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &output)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, output.Status)
	assert.NotEqual(t, nil, output.Data)
	assert.NotEqual(t, nil, output.Data)
	assert.Equal(t, nil, output.Error)
}
