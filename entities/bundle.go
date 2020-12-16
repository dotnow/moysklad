package entities

// Component ...
type Component struct {
	ID         string                 `json:"id,omitempty"`         // ID компонента (Только для чтения)
	AccountID  string                 `json:"accountId,omitempty"`  // ID учетной записи (Только для чтения)
	Quantity   float64                `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в компоненте (Только для чтения)
	Assortment map[string]interface{} `json:"assortment,omitempty"` // Метаданные товара/услуги/серии, которую представляет собой компонент
}

// Components ...
type Components struct {
	Meta *Meta       `json:"meta,omitempty"`
	Rows []Component `json:"rows,omitempty"`
}

// Bundle Комплект
type Bundle struct {
	Meta               *Meta         `json:"meta"`                    // Метаданные Комплекта
	ID                 string        `json:"id"`                      // ID Комплекта (Только для чтения)
	AccountID          string        `json:"accountId"`               // ID учетной записи (Только для чтения)
	Owner              Employee      `json:"owner"`                   // Метаданные владельца (Сотрудника)
	Shared             bool          `json:"shared"`                  // Общий доступ
	Group              Group         `json:"group"`                   // Метаданные отдела сотрудника
	SyncID             string        `json:"syncId,omitempty"`        // ID синхронизации
	Updated            string        `json:"updated"`                 // Момент последнего обновления сущности (Только для чтения)
	Name               string        `json:"name"`                    // Наименование Комплекта
	Description        string        `json:"description,omitempty"`   // Описание Комплекта
	Code               string        `json:"code,omitempty"`          // Код Комплекта
	ExternalCode       string        `json:"externalCode"`            // Внешний код Комплекта
	Archived           bool          `json:"archived"`                // Добавлен ли Комплект в архив
	PathName           string        `json:"pathName"`                // Наименование группы, в которую входит Комплект (Только для чтения)
	Vat                int           `json:"vat,omitempty"`           // НДС %
	EffectiveVat       int           `json:"effectiveVat,omitempty"`  // Реальный НДС % (Только для чтения)
	ProductFolder      ProductFolder `json:"productFolder,omitempty"` // Метаданные группы Комплекта
	Uom                Uom           `json:"uom,omitempty"`           // Единица измерения
	Images             Images        `json:"images,omitempty"`        // Изображения Комплекта
	MinPrice           Price         `json:"minPrice,omitempty"`      // Минимальная цена
	SalePrices         []SalePrices  `json:"salePrices,omitempty"`    // Цены продажи
	Attributes         []Attribute   `json:"attributes,omitempty"`    // Коллекция доп. полей
	Country            Country       `json:"country,omitempty"`       // Метаданные Страны
	Article            string        `json:"article,omitempty"`       // Артикул
	Weight             float64       `json:"weight,omitempty"`        // Вес
	Volume             float64       `json:"volume,omitempty"`        // Объем
	Barcodes           []Barcode     `json:"barcodes,omitempty"`      // Штрихкоды Комплекта
	DiscountProhibited bool          `json:"discountProhibited"`      // Признак запрета скидок
	Overhead           Price         `json:"overhead,omitempty"`      // Дополнительные расходы
	Components         Components    `json:"components,omitempty"`    // Компоненты Комплекта
	TrackingType       string        `json:"trackingType,omitempty"`  // Тип маркируемой продукции
	// --
	// Значения поля trackingType.
	// NOT_TRACKED	Без маркировки
	// TOBACCO	Тип маркировки "Табак"
	// SHOES	Тип маркировки "Обувь"
	// LP_CLOTHES	Тип маркировки "Одежда"
	// LP_LINENS	Тип маркировки "Постельное белье"
	// PERFUMERY	Духи и туалетная вода
	// ELECTRONICS	Фотокамеры и лампы-вспышки
	// --
	Tnved string `json:"tnved,omitempty"` // Код ТН ВЭД	—	нет
	// paymentItemType	Enum	Признак предмета расчета. Подробнее тут	—	нет

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
