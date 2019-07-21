package Item

import "fmt"

type DocItemSet struct {
	ItemList  []Item
	Sum         uint16
	NumOfDocs   int
}

func (docitemSet *DocItemSet) AddItem(content string, docid uint16) bool {
	words := g_Item.getItemWord(content)
	var distance = 0
	for _, word := range words {
		item := Item{}

		item.Keyword  = string(word)
		item.Location = distance
		item.DocId    = docid

		for _, singleword := range words {
			if singleword == word {
				item.Freq++
			}
		}

		distance = len(item.Keyword) + item.Location

		docitemSet.ItemList = append(docitemSet.ItemList, item)
		docitemSet.Sum++
	}
	return true
}

//test
// func (docitemSet *DocItemSet) AddItem_test() {
// 	docitemSet.AddItem("一般为赋值表达式，给控制变量赋初值", 1)
// 	docitemSet.AddItem("关系表达式或逻辑表达式", 2)
// 	docitemSet.AddItem("数学中的赋值，通过表达式来操作", 3)
// }

func (docitemSet *DocItemSet) PrintItem_test() {
	fmt.Print(docitemSet.Sum)
	fmt.Println("--------------------")
	for _, item := range docitemSet.ItemList {
		fmt.Println(item)
	}
}
