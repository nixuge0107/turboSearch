package main

import (
	"IndexItem"
	"Item"
	"Rank"
	"fmt"
)

var (
	itemSet        = Item.Item{}
	docItemSet     = Item.DocItemSet{}
	searchItemSet  = Item.SearchItemSet{}

	organizedIndex = IndexItem.Org1{}
	
	searchDocList  = Rank.SearchDocList{}
	scoreDocList   = Rank.ScoreDocList{}
)

func main() {
	var content   = []string{"一般为赋值表达式，给控制变量赋初值", "关系表达式或逻辑表达式", "数学中的赋值，通过表达式来操作"}
	var searchDoc = "赋值表达式赋值数学"

	//test
	itemSet.Init()
	fmt.Println("============docItemSet=============")
	for index := 0; index < len(content); index++ {
		docItemSet.AddItem(content[index], uint16(index))
	}
	docItemSet.NumOfDocs = len(content)
	docItemSet.PrintItem_test()
	fmt.Println("============searchItemSet=============")
	searchItemSet.AddItem(searchDoc, 0)
	searchItemSet.PrintItem_test()
	fmt.Println("============GoddamnIndexTest===========")
	organizedIndex.GenerateIndex(docItemSet)
	organizedIndex.PrintIndexList()
	fmt.Println("===========searchDoclistTest===========")
	searchDocList.InitSearchDoc(searchItemSet.ItemList, organizedIndex, docItemSet.NumOfDocs)
	searchDocList.GetOrganizeIndexListBySearch()
	fmt.Println(searchDocList.SearchDocIdList)
	fmt.Println(searchDocList.SearchIndexItemList)
	fmt.Println("============DocListScoreTest============")
	scoreDocList.InitScoreList(searchDocList)
	scoreDocList.Bm25score()
}
