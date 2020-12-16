package entities

// CounterParty Контрагент
type CounterParty struct {
	Meta               *Meta        `json:"meta"`                         // Метаданные Контрагента
	ID                 string       `json:"id,omitempty"`                 // ID Контрагента Только для чтения
	AccountID          string       `json:"accountId,omitempty"`          // ID учетной записи Только для чтения
	Owner              *Employee    `json:"owner,omitempty"`              // Владелец (Сотрудник)
	Shared             bool         `json:"shared,omitempty"`             // Общий доступ
	Group              Group        `json:"group,omitempty"`              // Отдел сотрудника
	SyncID             string       `json:"syncId,omitempty"`             // ID синхронизации После заполнения недоступен для изменения
	Updated            string       `json:"updated,omitempty"`            // Момент последнего обновления Контрагента Только для чтения
	Name               string       `json:"name,omitempty"`               // Наименование Контрагента Необходимое при создании
	Description        string       `json:"description,omitempty"`        // Комментарий к Контрагенту
	Code               string       `json:"code,omitempty"`               // Код Контрагента
	ExternalCode       string       `json:"externalCode,omitempty"`       // Внешний код Контрагента Только для чтения
	Archived           bool         `json:"archived,omitempty"`           // Добавлен ли Контрагент в архив
	Created            string       `json:"created,omitempty"`            // Момент создания
	Email              string       `json:"email,omitempty"`              // Адрес электронной почты
	Phone              string       `json:"phone,omitempty"`              // Номер городского телефона
	FAX                string       `json:"fax,omitempty"`                // Номер факса
	ActualAddress      string       `json:"actualAddress,omitempty"`      // Фактический адрес Контрагента
	ActualAddressFull  AddressFull  `json:"actualAddressFull,omitempty"`  // Фактический адрес Контрагента с детализацией по отдельным полям
	Accounts           Accounts     `json:"accounts,omitempty"`           // Массив счетов Контрагентов
	CompanyType        string       `json:"companyType,omitempty"`        // Тип Контрагента. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	DiscountCardNumber string       `json:"discountCardNumber,omitempty"` // Номер дисконтной карты Контрагента
	State              Status       `json:"state,omitempty"`              // Метаданные Статуса Контрагента
	SalesAmount        float64      `json:"salesAmount,omitempty"`        // Сумма продаж Только для чтения
	BonusProgram       BonusProgram `json:"bonusProgram,omitempty"`       // Метаданные активной Бонусной программы
	BonusPoints        int          `json:"bonusPoints,omitempty"`        // Бонусные баллы по активной бонусной программе Только для чтения
	Files              struct {
		Meta `json:"meta"`
	} `json:"files,omitempty"` // Массив метаданных Файлов (Максимальное количество файлов - 100)
}
