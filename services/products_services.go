package services

import (
	"prueba/models"
	"prueba/repositories"
)

type ProductService struct {
	mysql *repositories.MySqlRepositories
}

func MysqlProductService(mysql *repositories.MySqlRepositories) *ProductService {
	return &ProductService{
		mysql: mysql,
	}
}

func (sq *ProductService) GetProduct(id int) (models.Product, error) {
	product, err := sq.mysql.GetProduct(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
