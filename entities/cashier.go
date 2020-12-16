package entities

// Cashier Кассир
type Cashier struct {
	Meta        *Meta        `json:"meta"`        // Метаданные Кассира (Только для чтения)
	ID          string       `json:"id"`          // ID Сотрудника (Только для чтения)
	AccountID   string       `json:"accountId"`   // ID учетной записи (Только для чтения)
	Employee    *Employee    `json:"employee"`    // Метаданные сотрудника, которого представляет собой кассир
	RetailStore *RetailStore `json:"retailStore"` // Метаданные точки продаж, к которой прикреплен кассир
}
