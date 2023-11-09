package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"testTask/internal/models"
)

type IService interface {
	Unmarshal(data []byte) (map[string][]string, error)
	WriteToFile(breedsByCountry map[string][]string) error
}

type Service struct {
}

func New() IService {
	return &Service{}
}

func (s *Service) Unmarshal(data []byte) (map[string][]string, error) {

	var breedData models.BreedData

	err := json.Unmarshal(data, &breedData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	breedsByCountry := s.sort(&breedData)

	return breedsByCountry, nil

}

func (s *Service) sort(breedData *models.BreedData) map[string][]string {
	breedsByCountry := make(map[string][]string)
	for _, breed := range breedData.Breeds {
		breedsByCountry[breed.Country] = append(breedsByCountry[breed.Country], breed.Breed)
	}

	for _, breeds := range breedsByCountry {
		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i]) < len(breeds[j])
		})
	}

	return breedsByCountry
}

func (s *Service) WriteToFile(breedsByCountry map[string][]string) error {
	outFile, err := json.MarshalIndent(breedsByCountry, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	err = ioutil.WriteFile("out.json", outFile, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
