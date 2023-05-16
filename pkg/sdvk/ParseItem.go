package sdvk

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Выполнить запрос на парсинг карточки товара
func ItemRequest(link string) (item Item, ErrorParse error) {
	item.ForImage = make([]string, 0)
	item.ForTarget = make([]string, 0)
	item.ForStyle = make([]string, 0)
	item.ForType = make([]string, 0)
	item.ForTon = make([]string, 0)
	item.ForColor = make([]string, 0)
	item.Companion = make([]string, 0)
	item.Photo = make([]string, 0)

	c := colly.NewCollector()

	// Название товара
	c.OnHTML(`div[itemtype="http://schema.org/Product"]>div>div[class="arr-header"]>h1`, func(e *colly.HTMLElement) {
		item.Name = e.DOM.Text()
	})

	// Ссылка на товар
	c.OnHTML(`div[itemtype="http://schema.org/Product"]>meta[itemprop=url]`, func(e *colly.HTMLElement) {
		if Value, IsExist := e.DOM.Attr("content"); IsExist {
			item.Link = Value
		}
	})

	// Артикул
	c.OnHTML(`div[itemtype="http://schema.org/Product"]>meta[itemprop=sku]`, func(e *colly.HTMLElement) {
		if Value, IsExist := e.DOM.Attr("content"); IsExist {
			item.SKU = Value
		}
	})

	// Основное фото
	c.OnHTML(`img[class^="contained main-photo"]`, func(e *colly.HTMLElement) {
		if Value, IsExist := e.DOM.Attr("src"); IsExist {
			item.PhotoMain = Value[2:]
		}
	})
	// Фото
	c.OnHTML(`div[id="collection_photo_xs"] img`, func(e *colly.HTMLElement) {
		if Value, IsExist := e.DOM.Attr("src"); IsExist {
			if strings.Contains(Value, "jpg") {
				Value = strings.ReplaceAll(Value, "175.jpg", ".jpg")
				item.Photo = append(item.Photo, Value[2:])
			}
		}
	})

	// Компаньоны
	c.OnHTML(`div[class="companions-preview item"]>div[class=bottom]>div[class=art] a`, func(e *colly.HTMLElement) {
		FindText := e.Text
		if strings.Contains(FindText, "Артикул") {
			FindText = strings.ReplaceAll(FindText, "Артикул", "")
			FindText = strings.TrimSpace(FindText)
			item.Companion = append(item.Companion, FindText)
		}
	})

	// Цена
	c.OnHTML(`div[itemtype="https://schema.org/Offer"]>span[itemprop="price"]`, func(e *colly.HTMLElement) {
		if Value, IsExist := e.DOM.Attr("content"); IsExist {
			if ValueInt, ErrorAtoi := strconv.Atoi(Value); ErrorAtoi == nil {
				item.Price = ValueInt
			}
		}
	})

	// Описание товара
	c.OnHTML(`div[data-toggle=itemDescription]`, func(e *colly.HTMLElement) {
		item.Description = strings.TrimSpace(e.DOM.Text())
	})

	// Первые характеристики
	c.OnHTML(`table[class="attr-line-table"]`, func(e *colly.HTMLElement) {

		// Берём значение
		Value := e.DOM.Find("span").Text()
		if Value == "" {
			Value = e.DOM.Find("a").Text()
		}

		switch e.DOM.Find("label").Text() {
		case "Артикул:":
			item.SKU = Value
		case "Производитель:":
			item.Manuf = Value
		case "Страна:":
			item.Country = Value
		case "Коллекция:":
			item.Collection = Value
		case "Ширина рулона:":
			if metr, ErrorMetr := InMetr(Value); ErrorMetr == nil {
				item.Width = metr
			}
		case "Длина рулона:":
			if metr, ErrorMetr := InMetr(Value); ErrorMetr == nil {
				item.Length = metr
			}
		case "Материал:":
			item.Material = Value
		case "Основа:":
			item.Base = Value
		case "Экологичные:":
			item.IsEco = Value
		}
	})

	// Вторая характеристика
	c.OnHTML(`ul[id="itemTags"] li`, func(e *colly.HTMLElement) {
		label := e.DOM.Find("label").Text() // Название категории
		Value := e.DOM.Find("a").Text()     // Значение
		switch label {
		case "По изображению":
			item.ForImage = append(item.ForImage, Value)
		case "По назначению":
			item.ForTarget = append(item.ForImage, Value)
		case "По стилю":
			item.ForStyle = append(item.ForImage, Value)
		case "По типу":
			item.ForType = append(item.ForImage, Value)
		case "По тону":
			item.ForTon = append(item.ForImage, Value)
		case "По цвету":
			item.ForColor = append(item.ForImage, Value)
		}
	})

	c.Visit(URL + link)
	return item, nil
}

// На вход получить значение, а на выходе получить значение в метрах
//
// Пример:
//
//	'10.05 м' станет 10.05
//	'100 см' станет 1.0
func InMetr(str string) (output float64, Error error) {

	if strings.Contains(str, "см") { // Если это сантиметры
		str = strings.ReplaceAll(str, "см", "")
		str = strings.TrimSpace(str)
		if output, Error = strconv.ParseFloat(str, 64); Error != nil {
			return 0.0, Error
		}
		output /= 100
	} else {
		str = strings.ReplaceAll(str, "м", "")
		str = strings.TrimSpace(str)
		if output, Error = strconv.ParseFloat(str, 64); Error != nil {
			return 0.0, Error
		}
	}
	return output, nil
}
