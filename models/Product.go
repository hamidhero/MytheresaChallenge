package models

type Product struct {
	Id       int64  `gorm:"column:id" json:"id"`
	Sku      string `gorm:"column:sku" json:"sku"`
	Name     string `gorm:"column:name" json:"name"`
	Category string `gorm:"column:category" json:"category"`
	Price    int64  `gorm:"column:price" json:"price"`
}

//for compatibility with simdb library
func (c Product) ID() (jsonField string, value interface{}) {
	value = c.Id
	jsonField = "id"
	return
}

func (c Product) GetPrice() int64 {
	return c.Price
}