package models

type BreedData struct {
	Breeds []CatBreed `json:"data"`
}

type CatBreed struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
}
