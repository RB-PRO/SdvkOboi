package sdvk

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

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
			item.PhotoMain = Value
		}
	})
	// Фото
	c.OnHTML(`div[id="collection_photo_xs"] img`, func(e *colly.HTMLElement) {
		if Value, IsExist := e.DOM.Attr("src"); IsExist {
			if strings.Contains(Value, "jpg") {
				Value = strings.ReplaceAll(Value, "175.jpg", ".jpg")
				item.Photo = append(item.Photo, Value)
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

	// Сохранить
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		html, _ := e.DOM.Html()
		if err := os.WriteFile("file.html", []byte(html), 0666); err != nil {
			log.Fatal(err)
		}
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
			break
		case "Производитель:":
			item.Manuf = Value
			break
		case "Страна:":
			item.Country = Value
			break
		case "Коллекция:":
			item.Collection = Value
			break
		case "Ширина рулона:":
			// item.Width = Value
			break
		case "Длина рулона:":
			// item.Length = Value
			break
		case "Материал:":
			item.Material = Value
			break
		case "Основа:":
			item.Base = Value
			break
		case "Экологичные:":
			item.IsEco = Value
			break
		default:
			break
		}
	})

	// Вторая характеристика
	c.OnHTML(`ul[id="itemTags"] li`, func(e *colly.HTMLElement) {
		label := e.DOM.Find("label").Text() // Название категории
		Value := e.DOM.Find("a").Text()     // Значение
		switch label {
		case "По изображению":
			item.ForImage = append(item.ForImage, Value)
			break
		case "По назначению":
			item.ForTarget = append(item.ForImage, Value)
			break
		case "По стилю":
			item.ForStyle = append(item.ForImage, Value)
			break
		case "По типу":
			item.ForType = append(item.ForImage, Value)
			break
		case "По тону":
			item.ForTon = append(item.ForImage, Value)
			break
		case "По цвету":
			item.ForColor = append(item.ForImage, Value)
			break
		default:
			break
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
