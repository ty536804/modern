package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
)

const (
	MaxCons int = 100
	MinCons int = 2
)

func init()  {
	db, err = sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/shici?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("database error:",err)
		panic(err)
	}

	db.SetMaxOpenConns(MaxCons)
	db.SetMaxIdleConns(MinCons)


	err = db.Ping()

	if err != nil {
		panic(err)
	}

}

func checkError(err error) bool  {
	if err != nil {
		return true
	}
	return false
}