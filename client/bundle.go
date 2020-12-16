package client

import (
	"encoding/json"
	"strings"

	"github.com/dotnow/moysklad/entities"
)

// BundleRequest структура для запросов сущности 'bundle'
type BundleRequest struct{ *Request }

// Bundle устанавливает нужный endpoint
func (client *APIClient) Bundle(params Params) *BundleRequest {
	return &BundleRequest{client.newRequest("entity/bundle", params)}
}

// WithRelations устанавливает связи
// expand=components.assortment
func (request *BundleRequest) WithRelations(relations ...string) *BundleRequest {

	if len(relations) == 0 {
		return request
	}

	expand := strings.Join(relations, ",")

	request.params.AddQuery("expand", expand)

	return request
}

// GetByUUID возвращает сущность по UUID
func (request *BundleRequest) GetByUUID(uuid string) (bundle *entities.Bundle, err error) {

	response, err := request.getByUUID(uuid)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &bundle)
	if err != nil {
		return nil, err
	}

	return
}
