package moysklad

import (
	"encoding/json"
)

// CounterParty Контрагент
type CounterParty struct {
	Meta               *Meta        `json:"meta"`                         // Метаданные Контрагента
	ID                 string       `json:"id"`                           // ID Контрагента Только для чтения
	AccountID          string       `json:"accountId"`                    // ID учетной записи Только для чтения
	Owner              *Employee    `json:"owner"`                        // Владелец (Сотрудник)
	Shared             bool         `json:"shared"`                       // Общий доступ
	Group              Group        `json:"group"`                        // Отдел сотрудника
	SyncID             string       `json:"syncId,omitempty"`             // ID синхронизации После заполнения недоступен для изменения
	Updated            string       `json:"updated"`                      // Момент последнего обновления Контрагента Только для чтения
	Name               string       `json:"name"`                         // Наименование Контрагента Необходимое при создании
	Description        string       `json:"description,omitempty"`        // Комментарий к Контрагенту
	Code               string       `json:"code,omitempty"`               // Код Контрагента
	ExternalCode       string       `json:"externalCode,omitempty"`       // Внешний код Контрагента Только для чтения
	Archived           bool         `json:"archived"`                     // Добавлен ли Контрагент в архив
	Created            string       `json:"created"`                      // Момент создания
	Email              string       `json:"email,omitempty"`              // Адрес электронной почты
	Phone              string       `json:"phone,omitempty"`              // Номер городского телефона
	FAX                string       `json:"fax,omitempty"`                // Номер факса
	ActualAddress      string       `json:"actualAddress,omitempty"`      // Фактический адрес Контрагента
	ActualAddressFull  AddressFull  `json:"actualAddressFull,omitempty"`  // Фактический адрес Контрагента с детализацией по отдельным полям
	Accounts           []Account    `json:"accounts"`                     // Массив счетов Контрагентов
	CompanyType        string       `json:"companyType"`                  // Тип Контрагента. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	DiscountCardNumber string       `json:"discountCardNumber,omitempty"` // Номер дисконтной карты Контрагента
	State              Status       `json:"state"`                        // Метаданные Статуса Контрагента
	SalesAmount        int          `json:"salesAmount"`                  // Сумма продаж Только для чтения
	BonusProgram       BonusProgram `json:"bonusProgram,omitempty"`       // Метаданные активной Бонусной программы
	BonusPoints        int          `json:"bonusPoints,omitempty"`        // Бонусные баллы по активной бонусной программе Только для чтения
	Files              struct {
		Meta `json:"files,omitempty"`
	} `json:"files,omitempty"` // Массив метаданных Файлов (Максимальное количество файлов - 100)
}

// CounterpartyClient структура для запросов сущности 'counterparty'
type CounterpartyClient struct{ *Client }

// Counterparty устанавливает нужный endpoint
func (api *APIClient) Counterparty(params Params) *CounterpartyClient {
	return &CounterpartyClient{
		&Client{
			endPoint: "entity/counterparty",
			params:   params,
			api:      api,
		},
	}
}

// Get возвращает сущность по UUID
func (client *CounterpartyClient) Get(uuid string) (counterparty *CounterParty, err error) {

	response, err := client.getByUUID(uuid)

	err = json.Unmarshal(response, &counterparty)
	if err != nil {
		return nil, err
	}

	return counterparty, nil
}

// All возвращает список сущностей
func (client *CounterpartyClient) All() (counterparties []CounterParty, err error) {

	response, err := client.all()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &counterparties)
	if err != nil {
		return nil, err
	}

	return counterparties, nil
}
