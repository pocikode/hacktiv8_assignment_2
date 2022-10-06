package main

import (
	"assignment2/config"
	"assignment2/gin/controller"
	"assignment2/gin/repository/gorm"
	"assignment2/gin/router"
)

func main() {
	err := config.ConnectGorm()
	if err != nil {
		panic(err)
	}

	orm := config.GetGorm()
	orderRepo := gorm.NewOrderRepo(orm)

	orderHandler := controller.NewOrderHandler(orderRepo)

	router := router.NewRouter(orderHandler)
	router.CreateRouter().Run(":4000")
}
