package main

import (
	"IndexIterm"
	"Iterm"
	"Rank"
	"TurboEngine"
	"fmt"
)

var (
	turbo          = TurboEngine.Engine{}
	itermSet       = Iterm.Iterm{}
	docItermSet    = Iterm.DocItermSet{}
	searchItermSet = Iterm.SearchItermSet{}
	organizedIndex = IndexIterm.Org1{}
	searchDocList  = Rank.SearchDocList{}
	scoreDocList   = Rank.ScoreDocList{}
)

func main() {
	var content = []string{"一般为赋值表达式，给控制变量赋初值", "关系表达式或逻辑表达式", "数学中的赋值，通过表达式来操作"}
	var searchDoc = "赋值表达式"
	//无内容
	turbo.Init()
	turbo.AddDoc("内容")
	DocIdList := turbo.Search("搜索")
	fmt.Print(DocIdList)

	//test
	itermSet.Init()
	fmt.Println("============docItermSet=============")
	for index := 0; index < len(content); index++ {
		docItermSet.AddIterm(content[index], uint16(index))
	}
	docItermSet.NumOfDocs = len(content)
	docItermSet.PrintIterm_test()
	fmt.Println("============searchItermSet=============")
	searchItermSet.AddIterm(searchDoc, 0)
	searchItermSet.PrintIterm_test()
	fmt.Println("============GoddamnIndexTest===========")
	organizedIndex.GenerateIndex(docItermSet)
	organizedIndex.PrintIndexList()
	fmt.Println("===========searchDoclistTest===========")
	searchDocList.InitSearchDoc(searchItermSet.ItermList, organizedIndex)
	searchDocList.GetOrganizeIndexListBySearch()
	fmt.Println(searchDocList.SearchDocIdList)
	fmt.Println(searchDocList.SearchIndexItermList)
	fmt.Println("============DocListScoreTest============")
	scoreDocList.InitScoreList(searchDocList)
	scoreDocList.GetScore()
	fmt.Println(scoreDocList.ScoreLists)

}
