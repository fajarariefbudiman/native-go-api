package app

import (
	"belajar-golang-api/helper"
	"database/sql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:(localhost:3306)/category_repository")
	helper.PanicIfError(err)

	return db
}
