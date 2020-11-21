package moysklad

// Meta - Метаданные сущности
type Meta struct {
	Href         string `json:"href"`                   // Ссылка на объект
	MetadataHref string `json:"metadataHref,omitempty"` // Ссылка на метаданные сущности (Другой вид метаданных. Присутствует не во всех сущностях)
	Type         string `json:"type"`                   // Тип объекта
	MediaType    string `json:"mediaType"`              // application/json
	UUUIDHref    string `json:"uuidHref,omitempty"`     // Ссылка на объект на UI. Присутствует не во всех сущностях. Может быть использована для получения uuid
	DownloadHref string `json:"downloadHref,omitempty"` // Ссылка на скачивание Изображения. Данный параметр указывается только в meta для Изображения у Товара или Комплекта.
	Size         int    `json:"size,omitempty"`         // Размер выданного списка
	Limit        int    `json:"limit,omitempty"`        // Максимальное количество элементов в выданном списке. Максимальное количество элементов в списке равно 1000.
	Offset       int    `json:"offset,omitempty"`       // Отступ в выданном списке
	NextHref     string `json:"nextHref,omitempty"`     // Ссылка на следующую страницу сущностей.
	PreviousHref string `json:"previousHref,omitempty"` // Ссылка на предыдущую страницу сущностей.
}

// TODO: MSErrors – Структура ошибки
// Возвращаемые HTTP статусы ошибок и их значения:
// 301	Запрашиваемый ресурс находится по другому URL.
// 303	Запрашиваемый ресурс может быть найден по другому URI и должен быть найден с использоваием GET запроса
// 400	Ошибка в структуре JSON передаваемого запроса
// 401	Имя и/или пароль пользователя указаны неверно или заблокированы пользователь или аккаунт
// 403	У вас нет прав на просмотр данного объекта
// 404	Запрошенный ресурс не существует
// 405	http-метод указан неверно для запрошенного ресурса
// 409	Указанный объект используется и не может быть удален
// 410	Версия API больше не поддерживается
// 412	Не указан обязательный параметр строки запроса или поле структуры JSON
// 413	Размер запроса или количество элементов запроса превышает лимит (например, количество позиций, передаваемых в массиве positions, превышает 1000)
// 429	Превышен лимит количества запросов
// 500	При обработке запроса возникла непредвиденная ошибка
// 502	Сервис временно недоступен
// 503	Сервис временно отключен
// 504	Превышен таймаут обращения к сервису, повторите попытку позднее
// type MSErrors struct {
// 	Errors []struct {
// 		Error     string `json:"error,omitempty"`         // Заголовок ошибки
// 		Parameter string `json:"parameter,omitempty"`     // Параметр, на котором произошла ошибка
// 		Code      int    `json:"code,omitempty"`          // Код ошибки (Если поле ничего не содержит, смотрите HTTP status code)
// 		Message   string `json:"error_message,omitempty"` // Сообщение, прилагаемое к ошибке.
// 	} `json:"errors"`
// }

// Group Отдел
type Group struct {
	Meta *Meta  `json:"meta,omitempty"` // Метаданные Отдела
	Name string `json:"name,omitempty"` // Наименование Отдела
}

// Image структура изображения (image)
type Image struct {
	Meta      *Meta  `json:"meta,omitempty"`      // Метаданные объекта
	Title     string `json:"title,omitempty"`     // Название Изображения
	Filename  string `json:"filename,omitempty"`  // Имя файла
	Size      int    `json:"size,omitempty"`      // Размер файла в байтах
	Updated   string `json:"updated,omitempty"`   // Время последнего изменения
	Miniature Meta   `json:"miniature,omitempty"` // Метаданные миниатюры изображения
	Tiny      Meta   `json:"tiny,omitempty"`      // Метаданные уменьшенного изображения
}

// SalePrices Цена продажи
type SalePrices struct {
	Value     float64   `json:"value,omitempty"`     // Значение цены
	Currency  *Currency `json:"currency,omitempty"`  // Ссылка на валюту в формате
	PriceType PriceType `json:"priceType,omitempty"` // Тип цены
}

// ProductFolder Группа Товаров
type ProductFolder struct {
	Meta          *Meta          `json:"meta,omitempty"`          // Метаданные Группы Товара (Только для чтения)
	ID            string         `json:"id,omitempty"`            // ID Группы товаров (Только для чтения)
	AccountID     string         `json:"accountId,omitempty"`     // ID учетной записи (Только для чтения)
	Owner         *Employee      `json:"owner,omitempty"`         // Метаданные владельца (Сотрудника)
	Shared        bool           `json:"shared,omitempty"`        // Общий доступ
	Group         *Group         `json:"group,omitempty"`         // Метаданные отдела сотрудника
	Updated       string         `json:"updated,omitempty"`       // Момент последнего обновления сущности (Только для чтения)
	Name          string         `json:"name,omitempty"`          // Наименование Группы товаров
	Description   string         `json:"description,omitempty"`   // Описание Группы товаров
	Code          string         `json:"Description,omitempty"`   // Код Группы товаров
	ExternalCode  string         `json:"externalCode,omitempty"`  // Внешний код Группы товаров
	Archived      bool           `json:"archived,omitempty"`      // Добавлена ли Группа товаров в архив (Только для чтения)
	PathName      string         `json:"pathName,omitempty"`      // Наименование Группы товаров, в которую входит данная Группа товаров (Только для чтения)
	Vat           int            `json:"vat,omitempty"`           // НДС %
	EffectiveVat  int            `json:"effectiveVat,omitempty"`  // Реальный НДС % (Только для чтения)
	ProductFolder *ProductFolder `json:"productFolder,omitempty"` // Ссылка на Группу товаров, в которую входит данная Группа товаров, в формате Метаданных
	TaxSystem     string         `json:"taxSystem,omitempty"`     // Код системы налогообложения
	// --
	// taxSystem
	// TAX_SYSTEM_SAME_AS_GROUP	Совпадает с группой
	// GENERAL_TAX_SYSTEM	ОСН
	// SIMPLIFIED_TAX_SYSTEM_INCOME	УСН. Доход
	// SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME	УСН. Доход-Расход
	// UNIFIED_AGRICULTURAL_TAX	ЕСХН
	// PRESUMPTIVE_TAX_SYSTEM	ЕНВД
	// PATENT_BASED	Патент
	// --
}

// Currency Валюта
type Currency struct {
	Meta           *Meta         `json:"meta,omitempty"`           // Метаданные Валюты
	ID             string        `json:"id,omitempty"`             // ID Валюты (Только для чтения)
	Name           string        `json:"name,omitempty"`           // Краткое аименование Валюты
	FullName       string        `json:"fullName,omitempty"`       // Полное наименование Валюты
	Code           string        `json:"code,omitempty"`           // Цифровой код Валюты
	ISOCode        string        `json:"isoCode,omitempty"`        // Буквенный код Валюты
	Rate           float64       `json:"rate,omitempty"`           // Курс Валюты
	Multiplicity   int           `json:"multiplicity,omitempty"`   // Кратность курса Валюты
	Indirect       bool          `json:"indirect,omitempty"`       // Признак обратного курса Валюты
	RateUpdateType bool          `json:"rateUpdateType,omitempty"` // Способ обновления курса Валюты (Только для чтения)
	MajorUnit      *CurrencyUnit `json:"majorUnit,omitempty"`      // Формы единиц целой части Валюты
	MinorUnit      *CurrencyUnit `json:"minorUnit,omitempty"`      // Формы единиц дробной части Валюты
	Archived       bool          `json:"archived,omitempty"`       // Добавлена ли Валюта в архив
	System         bool          `json:"system,omitempty"`         // Основана ли валюта на валюте из системного справочника (Только для чтения)
	Default        bool          `json:"default,omitempty"`        // Является ли валюта валютой учета (Только для чтения)
}

// CurrencyUnit Форма единиц
type CurrencyUnit struct {
	Meta   *Meta  `json:"meta,omitempty"`   // Метаданные Формы единиц
	Gender string `json:"gender,omitempty"` // Грамматический род единицы валюты (допустимые значения masculine - мужской, feminine - женский)
	S1     string `json:"s1,omitempty"`     // Форма единицы, используемая при числительном 1
	S2     string `json:"s2,omitempty"`     // Форма единицы, используемая при числительном 2
	S5     string `json:"s5,omitempty"`     // Форма единицы, используемая при числительном 5
}

// Rate ...
type Rate struct {
	Currency *Currency `json:"currency,omitempty"`
}

// BuyPrice Закупочная цена
type BuyPrice struct {
	Value    float64   `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате
}

// Attribute Дополнительное поле
type Attribute struct {
	Meta  Meta        `json:"meta"`  // Ссылка на метаданные доп. поля
	ID    string      `json:"id"`    // Id соответствующего доп. поля
	Name  string      `json:"name"`  // Наименование доп. поля
	Value interface{} `json:"value"` // Значение, указанное в доп. поле.
}

// Barcode Штрихкоды
type Barcode struct {
	EAN13   string `json:"ean13,omitempty"`   // штрихкод в формате EAN13, если требуется создать штрихкод в формате EAN13
	EAN8    string `json:"ean8,omitempty"`    // штрихкод в формате EAN8, если требуется создать штрихкод в формате EAN8
	Code128 string `json:"code128,omitempty"` // штрихкод в формате Code128, если требуется создать штрихкод в формате Code128
	GTIN    string `json:"gtin,omitempty"`    // штрихкод в формате GTIN, если требуется создать штрихкод в формате GTIN. Валидируется на соответствие формату GS1
}

// File Файл
type File struct {
	Meta Meta `json:"meta"` // Метаданные объекта
	// Title     string   `json:"title,omitempty"`     // Название Файла
	// Filename  string   `json:"filename,omitempty"`  // Имя Файла
	// Size      int      `json:"size,omitempty"`      // Размер Файла в байтах
	// Created   string   `json:"created,omitempty"`   // Время загрузки Файла на сервер
	// CreatedBy Employee `json:"createdBy,omitempty"` // Метаданные сотрудника, загрузившего Файл
	// Miniature Meta     `json:"miniature,omitempty"` // Метаданные миниатюры изображения (поле передается только для Файлов изображений)
	// Tiny      Meta     `json:"tiny,omitempty"`      //	Метаданные уменьшенного изображения (поле передается только для Файлов изображений)
}

// Characteristics ...
type Characteristics []struct {
	Meta  *Meta  `json:"meta,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Uom Единица измерения
type Uom struct {
	Meta *Meta `json:"meta,omitempty"` // Метаданные Единицы измерения
}

// Stock ...
type Stock struct {
	Stock map[string]string `json:"stock,omitempty"`
}

// Country Метаданные страны
type Country struct {
	Meta *Meta `json:"meta,omitempty"` // Метаданные страны
}

// Region Метаданные региона
type Region struct {
	Meta *Meta `json:"meta,omitempty"` // Метаданные региона
}

// AddressFull Адрес
type AddressFull struct {
	PostalCode string   `json:"postalCode,omitempty"` // Почтовый индекс
	Country    *Country `json:"country,omitempty"`    // Метаданные страны
	Region     *Region  `json:"region,omitempty"`     // Метаданные региона
	City       string   `json:"city,omitempty"`       // Город
	Street     string   `json:"street,omitempty"`     // Улица
	House      string   `json:"house,omitempty"`      // Дом
	Apartment  string   `json:"apartment,omitempty"`  // Квартира
	AddInfo    string   `json:"addInfo,omitempty"`    // Другое
	Comment    string   `json:"comment,omitempty"`    // Комментарий
}

// LastOperation Последние операции
type LastOperation struct {
	Entity string `json:"entity,omitempty"` // Ключевое слово, обозначающее тип последней операции (Только для чтения)
	Name   string `json:"name,omitempty"`   // Наименование (номер) последней операции (Только для чтения)
}

// Status Статус документа
type Status struct {
	Meta      *Meta  `json:"meta,omitempty"`      // Метаданные Статуса
	ID        string `json:"id,omitempty"`        // ID Статуса (Только для чтения)
	AccountID string `json:"accountId,omitempty"` // ID учетной записи (Только для чтения)
	Name      string `json:"name,omitempty"`      // Наименование Статуса
	Color     string `json:"color,omitempty"`     // Цвет Статуса
	StateType string `json:"stateType,omitempty"` // Тип Статуса
	// --
	// Regular	Обычный (значение по умолчанию)
	// Successful	Финальный положительный
	// Unsuccessful	Финальный отрицательный
	// --
	EntityType string `json:"entityType,omitempty"` // Тип сущности, к которой относится Статус (ключевое слово в рамках JSON API)
}

// PrintTemplate Шаблон печатной формы
type PrintTemplate struct {
	Meta    *Meta  `json:"meta,omitempty"`    // 	Метаданные шаблона
	ID      string `json:"id,omitempty"`      // ID шаблона
	Name    string `json:"name,omitempty"`    // Наименование шаблона
	Type    string `json:"type,omitempty"`    // Тип шаблона (entity - документ)
	Content string `json:"content,omitempty"` //	Ссылка на скачивание
}
