package server

import (
	"prueba/controllers"
	"prueba/dal"
)

func Routes() {
	mysql := dal.NewDataBase()
	userController := controllers.ProductMySqlController(mysql)

	e.GET("/product/:id", userController.GetProduct)

}
