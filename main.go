package main

import (
	"Iterm"
	"TurboEngine"
	"fmt"
)

var (
	turbo    = TurboEngine.Engine{}
	itermSet = Iterm.ItermSet{}
)

func main() {
	turbo.Init()
	turbo.AddDoc("内容")
	DocIdList := turbo.Search("搜索")
	fmt.Print(DocIdList)

	//test
	itermSet.AddIterm_test()
	itermSet.PrintIterm_test()

}
