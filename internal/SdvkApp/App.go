package sdvkapp

import (
	"log"

	"ginthub.com/RB-PRO/SdvkOboi/pkg/sdvk"
	"github.com/cheggaaa/pb"
)

// Функция парсинга товаров и сохранения в xlsx
func Start() {
	lens := 1482
	lens = 10
	bar := pb.StartNew(lens)
	defer bar.Finish()
	Items := make([]sdvk.Item, 0)
	for i := 1; i <= lens; i++ {
		links, ErrorPage := sdvk.PageRequest(i)
		if ErrorPage != nil {
			log.Fatalln(ErrorPage)
		}
		TecalItems := make([]sdvk.Item, 0)
		for _, link := range links {
			item, _ := sdvk.ItemRequest(link)
			TecalItems = append(TecalItems, item)
		}
		// items
		Items = append(Items, TecalItems...)
		bar.Increment()
	}
	SaveXlsx("sdvk.xlsx", Items)
}
