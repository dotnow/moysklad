package client

import (
	"encoding/json"

	"github.com/dotnow/moysklad/entities"
)

// CounterpartyRequest структура для запросов сущности 'counterparty'
type CounterpartyRequest struct{ *Request }

// Counterparty устанавливает нужный endpoint
func (api *APIClient) Counterparty(params Params) *CounterpartyRequest {
	return &CounterpartyRequest{api.newRequest("entity/counterparty", params)}
}

// Get возвращает сущность по UUID
func (client *CounterpartyRequest) Get(uuid string) (counterparty *entities.CounterParty, err error) {

	response, err := client.getByUUID(uuid)

	err = json.Unmarshal(response, &counterparty)
	if err != nil {
		return nil, err
	}

	return counterparty, nil
}

// GetList возвращает список сущностей
func (client *CounterpartyRequest) GetList() (counterparties []entities.CounterParty, err error) {

	response, _, err := client.getList()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &counterparties)
	if err != nil {
		return nil, err
	}

	return counterparties, nil
}
