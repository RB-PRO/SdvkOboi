package sdvk

import (
	"strconv"

	"github.com/gocolly/colly/v2"
)

// https://sdvk-oboi.ru/oboi/page-3/#filter=1&group=item&order=rating
func PageRequest(Page int) ([]string, error) {
	Pages := make([]string, 0)

	c := colly.NewCollector()

	// Название товара
	c.OnHTML(`div[class="items-list goods"]>div[itemprop="itemListElement"]>div[class="content"]`, func(e *colly.HTMLElement) {
		if Value, IsExist := e.DOM.Attr("data-href"); IsExist {
			Pages = append(Pages, Value)
		}
	})
	c.Visit(URL + "/oboi/page-" + strconv.Itoa(Page) + "/#filter=1&group=item&order=rating")

	return Pages, nil
}
