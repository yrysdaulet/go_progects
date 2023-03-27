package main

import (
	"github.com/yrysdaulet/go_progects/shop"
	"github.com/yrysdaulet/go_progects/shop/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(shop.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
