package resources

import "MytheresaChallenge/models"

type ProductResource struct {
	Sku      string                  `json:"sku"`
	Name     string                  `json:"name"`
	Category string                  `json:"category"`
	Price    ProductPriceResource `json:"price"`
}

type ProductPriceResource struct {
	Original           int64   `json:"original"`
	Final              int64   `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Currency           string  `json:"currency"`
}

//GetProductResource accepts query result as model and returns required response type
func GetProductResource(products []models.Product) []ProductResource {
	var productResources []ProductResource
	for _, product := range products {

		//calculates final price and discount if exists
		p := ProductGetPriceResource{Product:product}
		finalPrice, discountPercentage := p.GetPrice()

		productResources = append(productResources, ProductResource{
			Sku:      product.Sku,
			Name:     product.Name,
			Category: product.Category,
			Price:    ProductPriceResource{
				Original:           product.Price,
				Final:              finalPrice,
				DiscountPercentage: discountPercentage,
				Currency:           "EUR",
			},
		})
	}

	return productResources
}