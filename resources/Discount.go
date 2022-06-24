package resources

import (
	"MytheresaChallenge/models"
	"MytheresaChallenge/utils"
)

type boot struct {
	Product models.Product
}

func (c *boot) GetPrice() (int64, *string) {
	percent := "30%"
	return utils.ApplyDiscount(c.Product.GetPrice(), 30), &percent
}

type sku3 struct {
	Product models.Product
}

func (c *sku3) GetPrice() (int64, *string) {
	percent := "15%"
	return utils.ApplyDiscount(c.Product.GetPrice(), 15), &percent
}

type ProductGetPriceResource struct {
	Product models.Product
}

//Calculates price of any kind of product either with discount or without it
func (c *ProductGetPriceResource) GetPrice() (int64, *string) {
	if c.Product.Category == "boots" {
		p := boot{Product:c.Product}
		return p.GetPrice()
	} else if c.Product.Sku == "000003" {
		p := sku3{Product:c.Product}
		return p.GetPrice()
	} else {
		return c.Product.GetPrice(), nil
	}
}