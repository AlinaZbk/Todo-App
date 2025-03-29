package main

import (
	"log"

	"github.com/AlinaZbk/ToDo-App"
	"github.com/AlinaZbk/ToDo-App/pkg/handler"
	"github.com/AlinaZbk/ToDo-App/pkg/repository"
	"github.com/AlinaZbk/ToDo-App/pkg/service"
)

func main(){
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(todo.Server) //созадем http-сервер
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil { //запускаем сервер
		log.Fatalf("error occured while running http server: %s", err.Error())
	
	}
}