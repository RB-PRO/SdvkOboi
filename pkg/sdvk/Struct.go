package sdvk

// Ссылка на сайт
const URL string = "https://sdvk-oboi.ru"

type Item struct {
	Link string // Ссылка на товар

	Name  string // Название товара
	Price int    // Цена

	/* Дополнительные поля */

	SKU        string  // Артикул
	Manuf      string  // Производитель
	Country    string  // Страна
	Collection string  // Коллекция
	Width      float64 // Ширина
	Length     float64 // Длина
	Material   string  // Материал
	Base       string  // Основа
	IsEco      string  // Экологичные

	ForImage  []string // По изображению
	ForTarget []string // По назначению
	ForStyle  []string // По стилю
	ForType   []string // По типу
	ForTon    []string // По тону
	ForColor  []string // По цвету
	// Плотность
	// Раппорт

	Description string   // Описание товара
	Companion   []string // Компаньоны

	PhotoMain string   // Главное Фото
	Photo     []string // Фото
}
