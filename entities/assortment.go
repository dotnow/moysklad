package entities

// Assortment струтура с разделением на типы объектов
type Assortment struct {
	Products []ProductEntity `json:"products"` // Товары
	Variants []Variant       `json:"variants"` // Модификации
	Services []Service       `json:"services"` // Услуги
	Bundles  []Bundle        `json:"bundles"`  // Комплекты
}

// Size возвращает общее количество объектов
func (a *Assortment) Size() int {
	return len(a.Products) + len(a.Variants) + len(a.Services) + len(a.Bundles)
}
