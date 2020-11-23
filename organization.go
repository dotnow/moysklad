package moysklad

import "encoding/json"

// Organization структура сущности Юрлицо
type Organization struct {
	Meta              *Meta        `json:"meta,omitempty"`              // Метаданные Юрлица
	ID                string       `json:"id,omitempty"`                // ID Юрлица (Только для чтения)
	AccountID         string       `json:"accountId,omitempty"`         // ID учетной записи (Только для чтения)
	Owner             *Employee    `json:"owner,omitempty"`             // Владелец (Сотрудник)
	Shared            bool         `json:"shared,omitempty"`            // Общий доступ
	Group             *Group       `json:"group,omitempty"`             // Отдел сотрудника
	SyncID            string       `json:"syncId,omitempty"`            // ID синхронизации
	Updated           string       `json:"updated,omitempty"`           // Момент последнего обновления Юрлица (Только для чтения)
	Name              string       `json:"name,omitempty"`              // Наименование Юрлица
	Description       string       `json:"description,omitempty"`       // Комментарий к Юрлицу
	Code              string       `json:"code,omitempty"`              // Код Юрлица
	ExternalCode      string       `json:"externalCode,omitempty"`      // Внешний код Юрлица (Только для чтения)
	Achived           bool         `json:"archived,omitempty"`          // Добавлено ли Юрлицо в архив
	Created           string       `json:"created,omitempty"`           // Дата создания
	ActualAddress     string       `json:"actualAddress,omitempty"`     // Фактический адрес Юрлица
	ActualAddressFull *AddressFull `json:"actualAddressFull,omitempty"` // Фактический адрес Юрлица с детализацией по отдельным полям
	CompanyType       string       `json:"companyType,omitempty"`       // Тип Юрлица. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	// --
	// companyType
	// legal	Юридическое лицо
	// entrepreneur	Индивидуальный предприниматель
	// individual	Физическое лицо
	// --
	TrackingContractNumber string        `json:"trackingContractNumber,omitempty"` // Номер договора с ЦРПТ
	TrackingContractDate   string        `json:"trackingContractDate,omitempty"`   // Дата договора с ЦРПТ
	BonusProgram           *BonusProgram `json:"bonusProgram,omitempty"`           // Метаданные активной бонусной программы
	BonusPoints            int           `json:"bonusPoints,omitempty"`            // Бонусные баллы по активной бонусной программе
	LegalTitle             string        `json:"legalTitle,omitempty"`             // Полное наименование. Игнорируется, если передано одно из значений для ФИО. Формируется автоматически на основе получаемых ФИО Юрлица
	LegalLastName          string        `json:"legalLastName,omitempty"`          // Фамилия для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalFirstName         string        `json:"legalFirstName,omitempty"`         // Имя для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalMiddleName        string        `json:"legalMiddleName,omitempty"`        // Отчество для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalAddress           string        `json:"legalAddress,omitempty"`           // Юридический адреса Юрлица
	LegalAddressFull       *AddressFull  `json:"legalAddressFull,omitempty"`       // Юридический адрес Юрлица с детализацией по отдельным полям
	INN                    string        `json:"inn,omitempty"`                    // ИНН
	KPP                    string        `json:"kpp,omitempty"`                    // КПП
	OGRN                   string        `json:"ogrn,omitempty"`                   // ОГРН
	OGRNIP                 string        `json:"ogrnip,omitempty"`                 // ОГРНИП
	OKPO                   string        `json:"okpo,omitempty"`                   // ОКПО
	// certificateNumber	Meta	// Номер свидетельства
	CertificateDate string      `json:"certificateDate,omitempty"` // Дата свидетельства
	Email           string      `json:"email,omitempty"`           // Адрес электронной почты
	Phone           string      `json:"phone,omitempty"`           // Номер городского телефона
	Fax             string      `json:"fax,omitempty"`             // Номер факса
	Accounts        []Account   `json:"accounts,omitempty"`        // Метаданные счетов юрлица
	Attributes      []Attribute `json:"attributes,omitempty"`      // Массив метаданных дополнительных полей юрлица
	IsEgaisEnable   bool        `json:"isEgaisEnable,omitempty"`   // Включен ли ЕГАИС для данного юрлица
	FSRARID         string      `json:"fsrarId,omitempty"`         // Идентификатор в ФСРАР
	PayerVat        bool        `json:"payerVat,omitempty"`        // Является ли данное юрлицо плательщиком НДС
	UTMURL          string      `json:"utmUrl,omitempty"`          // IP-адрес УТМ
	Director        string      `json:"director,omitempty"`        // Руководитель
	ChiefAccountant string      `json:"chiefAccountant,omitempty"` // Главный бухгалтер
}

// Account Счета юрлица
type Account struct {
	ID                   string        `json:"id,omitempty"`                   // ID Счета (Только для чтения)
	AccountID            string        `json:"accountId,omitempty"`            // ID учетной записи (Только для чтения)
	Updated              string        `json:"updated,omitempty"`              // Момент последнего обновления юрлица (Только для чтения)
	IsDefault            bool          `json:"isDefault,omitempty"`            // Является ли счет основным счетом юрлица
	AccountNumber        string        `json:"accountNumber,omitempty"`        // Номер счета	Необходимое при создании
	BankName             string        `json:"bankName,omitempty"`             // Наименование банка
	BankLocation         string        `json:"bankLocation,omitempty"`         // Адрес банка
	CorrespondentAccount string        `json:"correspondentAccount,omitempty"` // Корр счет
	BIC                  string        `json:"bic,omitempty"`                  // БИК
	Agent                *Organization `json:"agent,omitempty"`                // Метаданные юрлица
}

// OrganizationClient структура для запросов сущности 'employee'
type OrganizationClient struct{ *Client }

// Organization устанавливает нужный endpoint
func (api *APIClient) Organization(params Params) *OrganizationClient {
	return &OrganizationClient{
		&Client{
			endPoint: "entity/organization",
			params:   params,
			api:      api,
		},
	}
}

// GetByUUID возвращает сущность по UUID
func (client *OrganizationClient) GetByUUID(uuid string) (organization *Organization, err error) {

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

// Get возвращает список сущностей
func (client *OrganizationClient) Get() (organizations []Organization, err error) {

	response, err := client.all()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &organizations)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}
