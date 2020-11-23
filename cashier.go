package moysklad

import "encoding/json"

// Cashier Кассир
type Cashier struct {
	Meta        *Meta        `json:"meta,omitempty"`        // Метаданные Кассира (Только для чтения)
	ID          string       `json:"id,omitempty"`          // ID Сотрудника (Только для чтения)
	AccountID   string       `json:"accountId,omitempty"`   // ID учетной записи (Только для чтения)
	Employee    *Employee    `json:"employee,omitempty"`    // Метаданные сотрудника, которого представляет собой кассир
	RetailStore *RetailStore `json:"retailStore,omitempty"` // Метаданные точки продаж, к которой прикреплен кассир
}

// CashierClient структура для запросов сущности 'employee'
type CashierClient struct{ *Client }

// Cashier устанавливает нужный endpoint
func (api *APIClient) Cashier(params Params) *CashierClient {
	return &CashierClient{
		&Client{
			endPoint: "entity/cashier",
			params:   params,
			api:      api,
		},
	}
}

// GetByUUID возвращает сущность по UUID
func (client *CashierClient) GetByUUID(uuid string) (cashier *Cashier, err error) {

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

// Get возвращает список сущностей
func (client *CashierClient) Get() (cashiers []Cashier, err error) {

	response, err := client.all()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &cashiers)
	if err != nil {
		return nil, err
	}

	return cashiers, nil
}
