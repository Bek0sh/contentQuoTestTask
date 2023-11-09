package main

import (
	"fmt"
	"log"
	"testTask/internal/handler"
	"testTask/internal/service"
)

var h *handler.Handler

func init() {
	srv := service.New()
	h = handler.New(srv)
}

func main() {

	err := h.StartApp()
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println("Data written to out.json")
}
