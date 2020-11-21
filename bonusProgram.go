package moysklad

// BonusProgram Бонусная программа
type BonusProgram struct {
	Meta                    *Meta    `json:"meta,omitempty"`                    // Метаданные Бонусной программы
	ID                      string   `json:"id,omitempty"`                      // ID Бонусной программы (Только для чтения)
	AccountID               string   `json:"accountId,omitempty"`               // ID учетной записи (Только для чтения)
	Name                    string   `json:"name,omitempty"`                    // Наименование Бонусной программы
	Active                  bool     `json:"active,omitempty"`                  // Индикатор, является ли бонусная программа активной на данный момент
	AllProducts             bool     `json:"allProducts,omitempty"`             // Индикатор, действует ли бонусная программа на все товары (всегда true)
	AllAgents               bool     `json:"allAgents,omitempty"`               // Индикатор, действует ли скидка на всех контрагентов
	AgentTags               []string `json:"agentTags,omitempty"`               // Тэги контрагентов, к которым применяется бонусная программа. В случае пустого значения контрагентов в результате выводится пустой массив
	EarnRateRoublesToPoint  int      `json:"earnRateRoublesToPoint,omitempty"`  // Курс начисления
	SpendRatePointsToRouble int      `json:"spendRatePointsToRouble,omitempty"` // Курс списания
	MaxPaidRatePercents     int      `json:"maxPaidRatePercents,omitempty"`     // Максимальный процент оплаты баллами
}
