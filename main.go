package main

import (
	"IndexIterm"
	"Iterm"
	"TurboEngine"
	"fmt"
)

var (
	turbo          = TurboEngine.Engine{}
	itermSet       = Iterm.Iterm{}
	docItermSet    = Iterm.DocItermSet{}
	searchItermSet = Iterm.SearchItermSet{}
	organizedIndex = IndexIterm.OrganizedIndex{}
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
	organizedIndex.GenerateIndex(searchItermSet)
	organizedIndex.PrintIndexList()

}
