package repositories

import (
	"database/sql"
	"fmt"
	"prueba/models"
)

type MySqlRepositories struct {
	database *sql.DB
}

func ProductsRepositories(mysql *sql.DB) *MySqlRepositories {
	return &MySqlRepositories{
		database: mysql,
	}
}

func (sq MySqlRepositories) GetProduct(id int) (models.Product, error) {
	value, err := sq.database.Query(
		fmt.Sprintf(
			`select id, name, price from products where id=%d`,
			id,
		),
	)
	if err != nil {
		return models.Product{}, err
	}

	product := models.Product{}
	if value.Next() {
		err = value.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			return models.Product{}, err
		}
	}

	return product, nil
}
