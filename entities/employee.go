package entities

// Employee Сотрудник
type Employee struct {
	Meta         *Meta       `json:"meta"`                   // Метаданные Сотрудника
	ID           string      `json:"id,omitempty"`           // ID Сотрудника (Только для чтения)
	AccountID    string      `json:"accountId,omitempty"`    // ID учетной записи (Только для чтения)
	Owner        *Employee   `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       bool        `json:"shared,omitempty"`       // Общий доступ
	Group        *Group      `json:"group,omitempty"`        // Отдел сотрудника
	Updated      string      `json:"updated,omitempty"`      // Момент последнего обновления Сотрудника (Только для чтения)
	Name         string      `json:"name,omitempty"`         // Наименование Сотрудника (Только для чтения)
	Description  string      `json:"description,omitempty"`  // Комментарий к Сотруднику
	ExternalCode string      `json:"externalCode,omitempty"` // Внешний код Сотрудника (Только для чтения)
	Archived     bool        `json:"archived,omitempty"`     // Добавлен ли Сотрудник в архив
	Created      string      `json:"created,omitempty"`      // Момент создания Сотрудника (Только для чтения)
	UID          string      `json:"uid,omitempty"`          // Логин Сотрудника (Только для чтения)
	Email        string      `json:"email,omitempty"`        // Электронная почта сотрудника
	Phone        string      `json:"phone,omitempty"`        // Телефон сотрудника
	FirstName    string      `json:"firstName,omitempty"`    // Имя
	MiddleName   string      `json:"middleName,omitempty"`   // Отчество
	LastName     string      `json:"lastName,omitempty"`     // Фамилия
	FullName     string      `json:"fullName,omitempty"`     // Имя Отчество Фамилия (Только для чтения)
	ShortFio     string      `json:"shortFio,omitempty"`     // Краткое ФИО (Только для чтения)
	Cashiers     []Cashier   `json:"cashiers,omitempty"`     // Массив кассиров (Только для чтения)
	Attributes   []Attribute `json:"attributes,omitempty"`   // Дополнительные поля Сотрудника // TODO
	Image        *Image      `json:"image,omitempty"`        // Фотография сотрудника
	INN          string      `json:"inn,omitempty"`          // ИНН сотрудника (в формате ИНН физического лица)
	Position     string      `json:"position,omitempty"`     // Должность сотрудника
}
