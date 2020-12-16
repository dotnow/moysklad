package entities

// Service Услуга
type Service struct {
	Meta               *Meta          `json:"meta"`                         // Метаданные Услуги
	ID                 string         `json:"id"`                           // ID Услуги (Только для чтения)
	AccountID          string         `json:"accountId,omitempty"`          // ID учетной записи (Только для чтения)
	Owner              *Employee      `json:"owner,omitempty"`              // Метаданные владельца (Сотрудника)
	Shared             bool           `json:"shared,omitempty"`             // Общий доступ
	Group              *Group         `json:"group,omitempty"`              // Метаданные отдела сотрудника
	SyncID             string         `json:"syncId,omitempty"`             // ID синхронизации
	Updated            string         `json:"updated,omitempty"`            // Момент последнего обновления сущности (Только для чтения)
	Name               string         `json:"name,omitempty"`               // Наименование Услуги
	Description        string         `json:"description,omitempty"`        // Описание Услуги
	Code               string         `json:"code,omitempty"`               // Код Услуги
	ExternalCode       string         `json:"externalCode,omitempty"`       // Внешний код Услуги
	Archived           bool           `json:"archived,omitempty"`           // Добавлена ли Услуга в архив
	PathName           string         `json:"pathName,omitempty"`           // Наименование группы, в которую входит Услуга (Только для чтения)
	Vat                int            `json:"vat,omitempty"`                // НДС %
	EffectiveVat       int            `json:"effectiveVat,omitempty"`       // Реальный НДС % (Только для чтения)
	ProductFolder      *ProductFolder `json:"productFolder,omitempty"`      //	Метаданные группы
	Uom                *Uom           `json:"uom,omitempty"`                // Единицы измерения
	MinPrice           *Price         `json:"minPrice,omitempty"`           // Минимальная цена
	SalePrices         []SalePrices   `json:"salePrices,omitempty"`         // Цены продажи
	BuyPrice           *BuyPrice      `json:"buyPrice,omitempty"`           // Закупочная продажи
	Attributes         []Attribute    `json:"attributes,omitempty"`         // Коллекция доп. полей
	Barcodes           []Barcode      `json:"barcodes,omitempty"`           // Штрихкоды
	DiscountProhibited bool           `json:"discountProhibited,omitempty"` // Признак запрета скидок
	PaymentItemType    string         `json:"paymentItemType,omitempty"`    // Признак предмета расчета

	// --
	// Значения поля paymentItemType
	// SERVICE	Услуга
	// WORK	Работа
	// PROVIDING_RID	Предоставление РИД
	// COMPOUND_PAYMENT_ITEM	Составной предмет расчета
	// ANOTHER_PAYMENT_ITEM	Иной предмет расчета
	// --

	TaxSystem string `json:"taxSystem,omitempty"` // Код системы налогообложения

	// --
	// Значения поля taxSystem.
	// TAX_SYSTEM_SAME_AS_GROUP	Совпадает с группой
	// GENERAL_TAX_SYSTEM	ОСН
	// SIMPLIFIED_TAX_SYSTEM_INCOME	УСН. Доход
	// SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME	УСН. Доход-Расход
	// UNIFIED_AGRICULTURAL_TAX	ЕСХН
	// PRESUMPTIVE_TAX_SYSTEM	ЕНВД
	// PATENT_BASED	Патент
	// --

	Files struct {
		Meta `json:"files,omitempty"`
	} `json:"files,omitempty"` // Массив метаданных Файлов (Максимальное количество файлов - 100)
}
