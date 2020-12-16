package client

import (
	"encoding/json"

	"github.com/dotnow/moysklad/entities"
)

// OrganizationRequest структура для запросов сущности 'employee'
type OrganizationRequest struct{ *Request }

// Organization устанавливает нужный endpoint
func (api *APIClient) Organization(params Params) *OrganizationRequest {
	return &OrganizationRequest{api.newRequest("entity/organization", params)}
}

// GetByUUID возвращает сущность по UUID
func (client *OrganizationRequest) GetByUUID(uuid string) (organization *entities.Organization, err error) {

	response, err := client.getByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

// GetList возвращает список сущностей
func (client *OrganizationRequest) GetList() (organizations []entities.Organization, err error) {

	response, _, err := client.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &organizations)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}
