package app

import (
	"api/app/items"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

const (
	port string = ":8080"
)

// StartApp ...
func StartApp() {
	r = gin.Default()
	db := configDataBase()
	items.Configure(r, db)
	r.Run(port)
}

func configDataBase() *sql.DB {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "user", "userpwd", "db", "3306", "db"))

	fmt.Println("HERE ======")
	fmt.Println("connection STRING = ", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "user", "userpwd", "db", "3306", "db"))
	fmt.Println("err = ", err)
	fmt.Println("THERE ======")

	if err != nil {
		panic("Could not connect to the db")
	} else {
		return db
	}
}
