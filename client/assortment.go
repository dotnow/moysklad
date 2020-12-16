package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dotnow/moysklad/entities"
)

// AssortmentRequest структура для запросов сущности 'assortment'
type AssortmentRequest struct{ *Request }

// Assortment устанавливает нужный endpoint
func (client *APIClient) Assortment(params Params) *AssortmentRequest {
	return &AssortmentRequest{client.newRequest("entity/assortment", params)}
}

// GetList возвращает список сущностей
func (request *AssortmentRequest) GetList() (*entities.Assortment, error) {

	response, size, err := request.getList()
	if err != nil {
		return nil, err
	}

	assortment := &entities.Assortment{}
	responseItems := make([]map[string]interface{}, size)

	err = json.Unmarshal(response, &responseItems)
	if err != nil {
		return nil, err
	}

	for _, item := range responseItems {
		bytes, err := json.Marshal(item)
		if err != nil {
			log.Println(err)
			continue
		}

		responseData, err := getResponseData(bytes)
		if err != nil {
			log.Println(err)
			continue
		}

		switch responseData.Meta.Type {

		case "product":

			product := entities.ProductEntity{}

			err = json.Unmarshal(bytes, &product)
			if err != nil {
				log.Println(err)
				continue
			}

			assortment.Products = append(assortment.Products, product)

		case "variant":

			variant := entities.Variant{}

			err = json.Unmarshal(bytes, &variant)
			if err != nil {
				log.Println(err)
				continue
			}

			assortment.Variants = append(assortment.Variants, variant)

		case "service":

			service := entities.Service{}

			err = json.Unmarshal(bytes, &service)
			if err != nil {
				log.Println(err)
				continue
			}

			assortment.Services = append(assortment.Services, service)

		case "bundle":

			bundle := entities.Bundle{}

			err = json.Unmarshal(bytes, &bundle)
			if err != nil {
				log.Println(err)
				continue
			}

			assortment.Bundles = append(assortment.Bundles, bundle)

		default:
			fmt.Printf("Неизвестный тип объекта: %s\n", responseData.Meta.Type)
		}
	}

	return assortment, nil
}
