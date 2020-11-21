package moysklad

// Variant Модификация
type Variant struct {
	Meta            *Meta            `json:"meta,omitempty"`
	ID              string           `json:"id,omitempty"`
	AccountID       string           `json:"accountId,omitempty"`
	Version         int              `json:"version,omitempty"`
	Updated         string           `json:"updated,omitempty"`
	Name            string           `json:"name,omitempty"`
	Code            string           `json:"code,omitempty"`
	ExternalCode    string           `json:"externalCode,omitempty"`
	Archived        bool             `json:"archived,omitempty"`
	Characteristics *Characteristics `json:"characteristics,omitempty"`
	SalePrices      *SalePrices      `json:"salePrices,omitempty"`
	Barcodes        []string         `json:"barcodes,omitempty"`
	Product         *ProductEntity   `json:"product,omitempty"`
	Stock           int              `json:"stock,omitempty"`
	Reserve         int              `json:"reserve,omitempty"`
	InTransit       int              `json:"inTransit,omitempty"`
	Quantity        int              `json:"quantity,omitempty"`
}
