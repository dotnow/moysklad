package moysklad

// Bundle Комплект
type Bundle struct {
	Meta               *Meta          `json:"meta,omitempty"`               // Метаданные Комплекта
	ID                 string         `json:"id,omitempty"`                 // ID Комплекта (Только для чтения)
	AccountID          string         `json:"accountId,omitempty"`          // ID учетной записи (Только для чтения)
	Owner              *Employee      `json:"owner,omitempty"`              // Метаданные владельца (Сотрудника)
	Shared             bool           `json:"shared,omitempty"`             // Общий доступ
	Group              *Group         `json:"group,omitempty"`              // Метаданные отдела сотрудника
	SyncID             string         `json:"syncId,omitempty"`             // ID синхронизации
	Updated            string         `json:"updated,omitempty"`            // Момент последнего обновления сущности (Только для чтения)
	Name               string         `json:"name,omitempty"`               // Наименование Комплекта
	Description        string         `json:"description,omitempty"`        // Описание Комплекта
	Code               string         `json:"code,omitempty"`               // Код Комплекта
	ExternalCode       string         `json:"externalCode,omitempty"`       // Внешний код Комплекта
	Archived           bool           `json:"archived,omitempty"`           // Добавлен ли Комплект в архив
	PathName           string         `json:"pathName,omitempty"`           // Наименование группы, в которую входит Комплект (Только для чтения)
	Vat                int            `json:"vat,omitempty"`                // НДС %
	EffectiveVat       int            `json:"effectiveVat,omitempty"`       // Реальный НДС % (Только для чтения)
	ProductFolder      *ProductFolder `json:"productFolder,omitempty"`      // Метаданные группы Комплекта
	Uom                *Uom           `json:"uom,omitempty"`                // Единица измерения
	Images             []Image        `json:"images,omitempty"`             // Изображения Комплекта
	MinPrice           float64        `json:"minPrice,omitempty"`           // Минимальная цена
	SalePrices         []SalePrices   `json:"salePrices,omitempty"`         // Цены продажи
	Attributes         []Attribute    `json:"attributes,omitempty"`         // Коллекция доп. полей
	Country            *Country       `json:"country,omitempty"`            // Метаданные Страны
	Article            string         `json:"article,omitempty"`            // Артикул
	Weight             int            `json:"weight,omitempty"`             // Вес
	Volume             int            `json:"volume,omitempty"`             // Объем
	Barcodes           []Barcode      `json:"barcodes,omitempty"`           // Штрихкоды Комплекта
	DiscountProhibited bool           `json:"discountProhibited,omitempty"` // Признак запрета скидок
	Overhead           *Price         `json:"overhead,omitempty"`           // Дополнительные расходы
	Components         []struct {
		ID         string                   `json:"id,omitempty"`         // ID компонента (Только для чтения)
		AccountID  string                   `json:"accountId,omitempty"`  // ID учетной записи (Только для чтения)
		Quantity   int                      `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в компоненте (Только для чтения)
		Assortment []map[string]interface{} `json:"assortment,omitempty"` // Метаданные товара/услуги/серии, которую представляет собой компонент
	} `json:"components,omitempty"` // Компоненты Комплекта
	TrackingType string `json:"trackingType,omitempty"` // Тип маркируемой продукции
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
