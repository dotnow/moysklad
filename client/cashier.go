package client

import (
	"encoding/json"

	"github.com/dotnow/moysklad/entities"
)

// CashierRequest структура для запросов сущности 'employee'
type CashierRequest struct{ *Request }

// Cashier устанавливает нужный endpoint
func (api *APIClient) Cashier(params Params) *CashierRequest {
	return &CashierRequest{api.newRequest("entity/cashier", params)}
}

// GetByUUID возвращает сущность по UUID
func (client *CashierRequest) GetByUUID(uuid string) (cashier *entities.Cashier, err error) {

	response, err := client.getByUUID(uuid)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &cashier)
	if err != nil {
		return nil, err
	}

	return cashier, nil
}

// GetList возвращает список сущностей
func (client *CashierRequest) GetList() (cashiers []entities.Cashier, err error) {

	response, _, err := client.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &cashiers)
	if err != nil {
		return nil, err
	}

	return cashiers, nil
}
