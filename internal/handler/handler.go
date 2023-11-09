package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testTask/internal/service"
)

const url = "https://catfact.ninja/breeds"

type Handler struct {
	srv service.IService
}

func New(srv service.IService) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) StartApp() error {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to find url")
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading resonse body: %w", err)
	}

	breedsByCountry, err := h.srv.Unmarshal(data)
	if err != nil {
		return err
	}

	err = h.srv.WriteToFile(breedsByCountry)
	if err != nil {
		return err
	}

	return nil
}
