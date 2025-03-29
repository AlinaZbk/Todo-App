package main

import (
	"log"
	"github.com/AlinaZbk/ToDo-App"
	"github.com/AlinaZbk/ToDo-App/pkg/handler"
)

func main(){
	handlers := new(handler.Handler) //создаем обработчик http-запросов
	srv := new(todo.Server) //созадем http-сервер
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil { //запускаем сервер
		log.Fatalf("error occured while running http server: %s", err.Error())
	
	}
}