package client

import (
	"encoding/json"
	"strings"

	"github.com/dotnow/moysklad/entities"
)

// ProductRequest структура для запросов сущности 'product'
type ProductRequest struct{ *Request }

// Product устанавливает нужный endpoint
func (client *APIClient) Product(params Params) *ProductRequest {
	return &ProductRequest{client.newRequest("entity/product", params)}
}

// WithRelations устанавливает связи
func (request *ProductRequest) WithRelations(relations ...string) *ProductRequest {

	if len(relations) == 0 {
		return request
	}

	expand := strings.Join(relations, ",")

	request.params.AddQuery("expand", expand)

	return request
}

// GetByUUID возвращает сущность по UUID
func (request *ProductRequest) GetByUUID(uuid string) (product *entities.ProductEntity, err error) {

	response, err := request.getByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetList возвращает список сущностей
func (request *ProductRequest) GetList() (products entities.Products, err error) {

	response, _, err := request.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
