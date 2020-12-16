package client

import (
	"encoding/json"

	"github.com/dotnow/moysklad/entities"
)

// RetailstoreRequest структура для запросов сущности 'retailstore'
type RetailstoreRequest struct{ *Request }

// Retailstore устанавливает нужный endpoint
func (api *APIClient) Retailstore(params Params) *RetailstoreRequest {
	return &RetailstoreRequest{api.newRequest("entity/retailstore", params)}
}

// GetByUUID возвращает сущность по UUID
func (client *RetailstoreRequest) GetByUUID(uuid string) (retailstore *entities.RetailStore, err error) {

	response, err := client.getByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &retailstore)
	if err != nil {
		return nil, err
	}

	return retailstore, nil
}

// GetList возвращает список сущностей
func (client *RetailstoreRequest) GetList() (retailstores []entities.RetailStore, err error) {

	response, _, err := client.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &retailstores)
	if err != nil {
		return nil, err
	}

	return retailstores, nil
}
