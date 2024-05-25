package main

import (
	Api "compClub/internal/api"
	Controller "compClub/internal/controllers"
	DB "compClub/internal/db"
	"compClub/internal/redis"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	err := DB.DbInit()
	if err != nil {
		panic("can't connect to database")
	}
	redis.RedisInit()
	redis.InitPcInfo()
	handler := Api.NewRouter(Controller.NewController())

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		WriteTimeout: 10 * time.Second,
	}
	errors := make(chan error, 1)
	go func() {
		fmt.Println("Запущен")
		errors <- server.ListenAndServe()
	}()

	if err := <-errors; err != nil {
		log.Fatalf("Сервер остановил работу! : %v", err)
	}

	//test
	
}


