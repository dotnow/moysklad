package moysklad

import "encoding/json"

// Store Склад
type Store struct {
	Meta         *Meta        `json:"meta,omitempty"`         // Метаданные Склада
	ID           string       `json:"id,omitempty"`           // ID Склада (Только для чтения)
	AccountID    string       `json:"accountId,omitempty"`    // ID учетной записи (Только для чтения)
	Owner        *Employee    `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       bool         `json:"shared,omitempty"`       // Общий доступ
	Group        *Group       `json:"group,omitempty"`        // Отдел сотрудника
	Updated      string       `json:"updated,omitempty"`      // Момент последнего обновления Склада (Только для чтения)
	Name         string       `json:"name,omitempty"`         // Наименование Склада
	Description  string       `json:"description,omitempty"`  // Комментарий к Складу
	Code         string       `json:"code,omitempty"`         // Код Склада
	ExternalCode string       `json:"externalCode,omitempty"` // Внешний код Склада (Только для чтения)
	Archived     bool         `json:"archived,omitempty"`     // Добавлен ли Склад в архив
	Address      string       `json:"address,omitempty"`      // Адрес склада
	AddressFull  *AddressFull `json:"addressFull,omitempty"`  // Адрес с детализацией по отдельным полям
	Parent       *Store       `json:"parent,omitempty"`       // Метаданные родительского склада (Группы)
	PathName     string       `json:"pathName,omitempty"`     // Группа Склада
	Attributes   []Attribute  `json:"attributes,omitempty"`   // Массив метаданных дополнительных полей склада
}

// StoreClient структура для запросов сущности 'retailstore'
type StoreClient struct{ *Client }

// Store устанавливает нужный endpoint
func (api *APIClient) Store(params Params) *StoreClient {
	return &StoreClient{
		&Client{

			endPoint: "entity/store",
			params:   params,
			api:      api,
		},
	}
}

// GetByUUID возвращает сущность по UUID
func (client *StoreClient) GetByUUID(uuid string) (store *Store, err error) {

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

// Get возвращает список сущностей
func (client *StoreClient) Get() (stores []Store, err error) {

	response, err := client.all()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &stores)
	if err != nil {
		return nil, err
	}

	return stores, nil
}
