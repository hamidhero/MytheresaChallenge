package connections

import (
	"MytheresaChallenge/database"
	"github.com/sonyarouje/simdb"
	"log"
)

var DB *simdb.Driver
var err error

func Connect()  {
	if DB, err = database.Init(); err != nil {
		log.Print("Database not initialized")
	}
}

