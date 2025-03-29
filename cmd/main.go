package main

import (
	"log"
	"github.com/spf13/viper"
	"github.com/AlinaZbk/ToDo-App"
	"github.com/AlinaZbk/ToDo-App/pkg/handler"
	"github.com/AlinaZbk/ToDo-App/pkg/repository"
	"github.com/AlinaZbk/ToDo-App/pkg/service"
)

func main(){
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(todo.Server) //созадем http-сервер
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil { //запускаем сервер
		log.Fatalf("error occured while running http server: %s", err.Error())
	
	}
}

func initConfig() error {
	viper.AddConfigPath("configs") //путь к конфигу
	viper.SetConfigName("config") //имя кофигурационного файла
	return viper.ReadInConfig() //читаем конфиг
}