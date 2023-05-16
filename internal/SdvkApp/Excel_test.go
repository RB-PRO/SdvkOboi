package sdvkapp

import (
	"testing"

	"ginthub.com/RB-PRO/SdvkOboi/pkg/sdvk"
)

func TestSaveXlsx(t *testing.T) {

	link := "/oboi/Milassa/Modern/330521/"
	Item, err := sdvk.ItemRequest(link)
	if err != nil {
		t.Error(err)
	}
	Items := []sdvk.Item{Item}
	SaveXlsx("TestSave.xlsx", Items)

}
