package mymiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var confDB = func() *leveldb.DB {
	db, err := leveldb.OpenFile("conf.db", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Put([]byte("password"), []byte("123456"), nil)
	if err != nil {
		log.Fatal(err)
	}
	return db

}()

func LoadConf(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("panelConfig", confDB)
		return next(c)
	}
}
