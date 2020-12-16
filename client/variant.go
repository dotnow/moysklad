package client

import (
	"encoding/json"

	"github.com/dotnow/moysklad/entities"
)

// VariantRequest структура для запросов сущности 'variant'
type VariantRequest struct{ *Request }

// Variant устанавливает нужный endpoint
func (api *APIClient) Variant(params Params) *VariantRequest {
	return &VariantRequest{api.newRequest("entity/variant", params)}
}

// GetByUUID возвращает сущность по UUID
func (client *VariantRequest) GetByUUID(uuid string) (variant *entities.Variant, err error) {

	response, err := client.getByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &variant)
	if err != nil {
		return nil, err
	}

	return variant, nil
}

// GetList возвращает список сущностей
func (client *VariantRequest) GetList() (variants []entities.Variant, err error) {

	response, _, err := client.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &variants)
	if err != nil {
		return nil, err
	}

	return variants, nil
}
