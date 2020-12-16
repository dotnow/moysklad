package client

import (
	"encoding/json"

	"github.com/dotnow/moysklad/entities"
)

// StoreRequest структура для запросов сущности 'retailstore'
type StoreRequest struct{ *Request }

// Store устанавливает нужный endpoint
func (api *APIClient) Store(params Params) *StoreRequest {
	return &StoreRequest{api.newRequest("entity/store", params)}
}

// GetByUUID возвращает сущность по UUID
func (client *StoreRequest) GetByUUID(uuid string) (store *entities.Store, err error) {

	response, err := client.getByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

// GetList возвращает список сущностей
func (client *StoreRequest) GetList() (stores []entities.Store, err error) {

	response, _, err := client.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &stores)
	if err != nil {
		return nil, err
	}

	return stores, nil
}
