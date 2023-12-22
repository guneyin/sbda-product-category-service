package main

import (
	"github.com/guneyin/sbda-product-category-service/service"
	"log"
)

func main() {
	svc, err := service.NewService()
	if err != nil {
		log.Fatal(err)
	}

	err = svc.Register()
	if err != nil {
		log.Fatal(err)
	}

	if err = svc.Serve(); err != nil {
		log.Fatal(err)
	}
}
