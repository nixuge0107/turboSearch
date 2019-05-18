package turboSearch

import "fmt"

func main()  {
	init()
	turbo.AddDoc("内容")
	DocIdList := turbo.Search("搜索")
	fmt.Print(DocIdList)
}