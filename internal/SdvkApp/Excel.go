package sdvkapp

import (
	"strconv"
	"strings"

	"ginthub.com/RB-PRO/SdvkOboi/pkg/sdvk"
	"github.com/xuri/excelize/v2"
)

func SaveXlsx(FileName string, Data []sdvk.Item) error {
	ssheet := "main"
	f := excelize.NewFile()
	defer f.Close()
	f.NewSheet(ssheet)
	f.DeleteSheet("Sheet1")
	if err := f.SaveAs(FileName); err != nil {
		return err
	}

	writeHeadOne(f, ssheet, 1, 1, "Артикул")
	writeHeadOne(f, ssheet, 2, 1, "Бренд")
	writeHeadOne(f, ssheet, 3, 1, "Коллекция")
	writeHeadOne(f, ssheet, 4, 1, "Наименование")
	writeHeadOne(f, ssheet, 5, 1, "Длина")
	writeHeadOne(f, ssheet, 6, 1, "Ширина")
	writeHeadOne(f, ssheet, 7, 1, "Страна")
	writeHeadOne(f, ssheet, 8, 1, "Рисунок")
	writeHeadOne(f, ssheet, 9, 1, "Тип покрытия")
	writeHeadOne(f, ssheet, 10, 1, "Основа")
	writeHeadOne(f, ssheet, 11, 1, "Покрытие")
	writeHeadOne(f, ssheet, 12, 1, "Тип помещения")
	writeHeadOne(f, ssheet, 13, 1, "Раппорт")
	writeHeadOne(f, ssheet, 14, 1, "Цена")
	writeHeadOne(f, ssheet, 15, 1, "Стиль")
	writeHeadOne(f, ssheet, 16, 1, "По типу")
	writeHeadOne(f, ssheet, 17, 1, "Стойкость")
	writeHeadOne(f, ssheet, 18, 1, "По тону")
	writeHeadOne(f, ssheet, 19, 1, "По изображению")
	writeHeadOne(f, ssheet, 20, 1, "По цвету")
	writeHeadOne(f, ssheet, 21, 1, "Плотность")
	writeHeadOne(f, ssheet, 22, 1, "Ссылка на фото")

	for index, item := range Data {
		writeHeadOne(f, ssheet, 1, index+2, item.SKU)
		writeHeadOne(f, ssheet, 2, index+2, item.Manuf)
		writeHeadOne(f, ssheet, 3, index+2, item.Collection)
		writeHeadOne(f, ssheet, 4, index+2, item.Name)
		writeHeadOne(f, ssheet, 5, index+2, item.Length)
		writeHeadOne(f, ssheet, 6, index+2, item.Width)
		writeHeadOne(f, ssheet, 7, index+2, item.Country)
		writeHeadOne(f, ssheet, 8, index+2, item.PhotoMain)
		writeHeadOne(f, ssheet, 9, index+2, "Тип покрытия")
		writeHeadOne(f, ssheet, 10, index+2, item.Base)
		writeHeadOne(f, ssheet, 11, index+2, item.Material)
		writeHeadOne(f, ssheet, 12, index+2, strings.Join(item.ForTarget, ";"))
		writeHeadOne(f, ssheet, 13, index+2, "Раппорт")
		writeHeadOne(f, ssheet, 14, index+2, item.Price)
		writeHeadOne(f, ssheet, 15, index+2, strings.Join(item.ForStyle, ";"))
		writeHeadOne(f, ssheet, 16, index+2, strings.Join(item.ForType, ";"))
		writeHeadOne(f, ssheet, 17, index+2, "Стойкость")
		writeHeadOne(f, ssheet, 18, index+2, strings.Join(item.ForTon, ";"))
		writeHeadOne(f, ssheet, 19, index+2, strings.Join(item.ForImage, ";"))
		writeHeadOne(f, ssheet, 20, index+2, strings.Join(item.ForColor, ";"))
		writeHeadOne(f, ssheet, 21, index+2, "Плотность")
		writeHeadOne(f, ssheet, 22, index+2, strings.Join(item.Photo, ";"))
	}
	f.Save()
	// strings.Join(asd, ";")
	return nil
}

// Вписать шапку
func writeHeadOne(f *excelize.File, ssheet string, col int, row int, val interface{}) error {
	collumn, ErrCol := excelize.ColumnNumberToName(col)
	if ErrCol != nil {
		return ErrCol
	}
	ErrSetCellValue := f.SetCellValue(ssheet, collumn+strconv.Itoa(row), val)
	if ErrSetCellValue != nil {
		return ErrSetCellValue
	}

	return nil
}
