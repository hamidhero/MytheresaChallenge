package database

import (
	"MytheresaChallenge/models"
	"encoding/json"
	db "github.com/sonyarouje/simdb"
	"io/ioutil"
	"path/filepath"
)

//Init initializes database
func Init() (*db.Driver, error) {
	productJson := "./products.json"
	dir, err := filepath.Abs(filepath.Dir(productJson))
	if err != nil {
		return nil, err
	}

	//selects of creates product model
	dbDir := dir + "/database/product"
	dbLocal, err := db.New(dbDir)
	if err != nil {
		return nil, err
	}

	_, err = ioutil.ReadFile(dbDir + "/product")
	if err != nil {
		//if database does not exist, fill it with json data
		content, err := ioutil.ReadFile(dir + "/database/" + productJson)
		if err != nil {
			return nil, err
		}

		type ProductsFromJson struct {
			Products []models.Product `json:"products"`
		}

		var data ProductsFromJson
		err = json.Unmarshal(content, &data)
		if err != nil {
			return nil, err
		}

		for _, product := range data.Products {
			err = dbLocal.Insert(product)
			if err != nil {
				return nil, err
			}
		}
	}

	return dbLocal, nil
}
