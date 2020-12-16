package entities

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
