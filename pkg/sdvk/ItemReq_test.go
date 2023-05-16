package sdvk

import (
	"fmt"
	"testing"
)

func TestItemRequest(t *testing.T) {
	link := "/oboi/Milassa/Modern/330521/"
	Item, err := ItemRequest(link)
	if err != nil {
		t.Error(err)
	}

	// fmt.Printf("Для товара %v собраны данные:\n%+v\n", URL+link, Item)

	fmt.Println("Ссылка на товар:", Item.Link) // Ссылка на товар

	fmt.Println("Название товара:", Item.Name) // Название товара
	fmt.Println("Цена:", Item.Price)           // Цена

	/* Дополнительные поля */

	fmt.Println("Артикул:", Item.SKU)             // Артикул
	fmt.Println("Производитель:", Item.Manuf)     // Производитель
	fmt.Println("Страна:", Item.Country)          // Страна
	fmt.Println("Коллекция:", Item.Collection)    // Коллекция
	fmt.Println("Ширина:", Item.Width)            // Ширина
	fmt.Println("Длина:", Item.Length)            // Длина
	fmt.Println("Материал:", Item.Material)       // Материал
	fmt.Println("Основа:", Item.Base)             // Основа
	fmt.Println("Экологичные:", Item.IsEco)       // Экологичные
	fmt.Println("По изображению:", Item.ForImage) // По изображению
	fmt.Println("По назначению:", Item.ForTarget) // По назначению
	fmt.Println("По стилю:", Item.ForStyle)       // По стилю
	fmt.Println("По типу:", Item.ForType)         // По типу
	fmt.Println("По тону:", Item.ForTon)          // По тону
	fmt.Println("По цвету:", Item.ForColor)       // По цвету
	fmt.Println("Плотность:")                     // Плотность
	fmt.Println("Раппорт:")                       // Раппорт

	fmt.Println("Описание товара:", Item.Description) // Описание товара
	fmt.Println("Компаньоны:", Item.Companion)        // Компаньоны

	fmt.Println("Главное Фото:", Item.PhotoMain) // Главное Фото
	fmt.Println("Фото:", Item.Photo)             // Фото

}

func TestInMetr(t *testing.T) {
	//	'10.05 м' станет 10.05
	//	'100 см' станет 1.0
	input1 := "10.05 м"
	answer1 := 10.05
	output1, Error1 := InMetr(input1)
	if Error1 != nil {
		t.Error(Error1)
	}
	if answer1 != output1 {
		t.Errorf("Для строки '%v' должно быть получиться %v, а получено: %v", input1, answer1, output1)
	}

	input2 := "100 см"
	answer2 := 1.0
	output2, Error2 := InMetr(input2)
	if Error2 != nil {
		t.Error(Error2)
	}
	if answer2 != output2 {
		t.Errorf("Для строки '%v' должно быть получиться %v, а получено: %v", input2, answer2, output2)
	}
}
