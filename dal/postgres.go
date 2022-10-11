package dal

import (
	"database/sql"
	"fmt"
	"log"
	"prueba/constants"
)

func NewDataBase() *sql.DB {
	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		constants.DBUser,
		constants.DBPassword,
		constants.DBHost,
		constants.DBPort,
		constants.DBName,
	)

	db, err := sql.Open(constants.DBType, database)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connect to database")

	return db
}
