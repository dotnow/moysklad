package client

import (
	"encoding/json"

	"github.com/dotnow/moysklad/entities"
)

// ServiceRequest структура для запросов сущности 'service'
type ServiceRequest struct{ *Request }

// Service устанавливает нужный endpoint
func (api *APIClient) Service(params Params) *ServiceRequest {
	return &ServiceRequest{api.newRequest("entity/service", params)}
}

// GetByUUID возвращает сущность по UUID
func (client *ServiceRequest) GetByUUID(uuid string) (service *entities.Service, err error) {

	response, err := client.getByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &service)
	if err != nil {
		return nil, err
	}
	return service, nil
}

// GetList возвращает список сущностей
func (client *ServiceRequest) GetList() (services []entities.Service, err error) {

	response, _, err := client.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &services)
	if err != nil {
		return nil, err
	}

	return services, nil
}
