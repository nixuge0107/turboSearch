package turboSearch

import (
	"TurboEngine"
	"fmt"
)

var (
	turbo = TurboEngine.Engine{}
)

func main() {
	turbo.Init()
	turbo.AddDoc("内容")
	DocIdList := turbo.Search("搜索")
	fmt.Print(DocIdList)
}
