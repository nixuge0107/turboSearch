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
	organizedIndex = IndexIterm.OrganizedIndex{}
	searchDocList  = Rank.SearchDocList{}
)

func main() {
	//无内容
	turbo.Init()
	turbo.AddDoc("内容")
	DocIdList := turbo.Search("搜索")
	fmt.Print(DocIdList)

	//test
	itermSet.Init()
	fmt.Println("============docItermSet=============")
	docItermSet.AddIterm_test()
	docItermSet.PrintIterm_test()
	fmt.Println("============searchItermSet=============")
	searchItermSet.AddIterm_test()
	searchItermSet.PrintIterm_test()
	fmt.Println("============GoddamnIndexTest===========")
	organizedIndex.GenerateIndex(docItermSet)
	organizedIndex.PrintIndexList()
	fmt.Println("===========searchDoclistTest===========")
	searchDocList.InitSearchDoc(searchItermSet.ItermList, organizedIndex)
	searchDocList.GetOrganizeIndexListBySearch()
	fmt.Println("============DocListScoreTest============")

}
