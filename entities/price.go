package entities

// Price Цена
type Price struct {
	Value    float64   `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

// PriceType Тип цен
type PriceType struct {
	Meta         *Meta  `json:"meta,omitempty"`         // Метаданные Типа цены (Только для чтения)
	ID           string `json:"id,omitempty"`           // ID типа цены (Только для чтения)
	Name         string `json:"name,omitempty"`         // Наименование Типа цены
	ExternalCode string `json:"externalCode,omitempty"` // Внешний код Типа цены
}
