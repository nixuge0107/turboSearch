package Item

import "fmt"

type SearchItemSet struct {
	ItemList []Item
	Sum        uint16
}

func (searchitemSet *SearchItemSet) AddItem(content string, docid uint16) bool {
	words := g_Item.getItemWord(content)
	var distance = 0
	for _, word := range words {
		item := Item{}
		item.Keyword = string(word)
		item.Location = distance
		item.DocId = docid

		for _, singleword := range words {
			if singleword == word {
				item.Freq++
			}
		}

		distance = len(item.Keyword) + item.Location

		searchitemSet.ItemList = append(searchitemSet.ItemList, item)
		searchitemSet.Sum++
	}
	return true
}

//test
// func (searchitemSet *SearchItemSet) AddItem_test() {
// 	searchitemSet.AddItem("赋值表达式", 1)
// }

func (searchitemSet *SearchItemSet) PrintItem_test() {
	fmt.Print(searchitemSet.Sum)
	fmt.Println("--------------------")
	for _, item := range searchitemSet.ItemList {
		fmt.Println(item)
	}
}
